package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	hellopb "github.com/ghazninattarshah/tracing-with-jaeger/gunk/hello"
	"github.com/ghazninattarshah/tracing-with-jaeger/services/hello"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	metrics_prom "github.com/uber/jaeger-lib/metrics/prometheus"

	"google.golang.org/grpc"
)

var (
	serviceName     = "hello"
	defaultHttpPort = ":9000"
	defaultGrpcPort = "localhost:9001"
)

func main() {
	flag.Parse()

	addr := defaultHttpPort
	endpoint := defaultGrpcPort

	if envServiceName := os.Getenv("SERVICE_NAME"); envServiceName != "" {
		serviceName = envServiceName
	}
	if httpPort := os.Getenv("HTTP_PORT"); httpPort != "" {
		addr = ":" + httpPort
	}
	if grpcPort := os.Getenv("GRPC_PORT"); grpcPort != "" {
		endpoint = strings.ToLower(serviceName) + ":" + grpcPort
	}

	log.Println("Starting service :", serviceName, "on", addr)
	ctx := context.Background()

	log.Println("Loading the tracer from env")
	// initialize tracer
	cfg, _ := config.FromEnv()

	// create tracer from config
	tracer, closer, err := cfg.NewTracer(
		config.Metrics(metrics_prom.New()),
	)
	defer closer.Close()
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to create a new tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)

	// graceful shutdown
	errch := make(chan error, 2)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
	)

	helloService := hello.NewHelloService(serviceName, tracer, nil)

	// next service
	if nextServiceHost := os.Getenv("NEXT_SERVICE_HOST"); nextServiceHost != "" {
		nextServicePort := os.Getenv("NEXT_SERVICE_PORT")
		if nextServicePort == "" {
			log.Fatalln("no next service port specified")
		}
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                time.Second * 10,
				Timeout:             time.Second * 5,
				PermitWithoutStream: true,
			}),
			grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
				grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(tracer)),
			)),
			grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
				grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
			)),
		}
		grpcAddr := net.JoinHostPort(nextServiceHost, nextServicePort)
		conn, err := grpc.Dial(grpcAddr, opts...)
		if err != nil {
			log.Fatalf("failed dialing next service grpc.Dial %q err: %v", grpcAddr, err)
		}
		defer conn.Close()

		nextService := hellopb.NewHelloServiceClient(conn)
		helloService = hello.NewHelloService(serviceName, tracer, nextService)
	}

	// grpc server
	go func() {
		errch <- func() error {
			l, err := net.Listen("tcp", endpoint)
			if err != nil {
				log.Fatalf("error listing on : %v", err)
			}
			reflection.Register(s)
			hellopb.RegisterHelloServiceServer(s, helloService)
			return s.Serve(l)
		}()
	}()

	// http server
	go func() {
		errch <- func() error {
			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithInsecure()}
			if err := hellopb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
				return err
			}
			return http.ListenAndServe(addr, mux)
		}()
	}()

	// catch stop signal
	go func() {
		errch <- func() error {
			<-sigs
			log.Println("Stopping GRPC Server..")
			s.GracefulStop()
			return errors.New("server gracefully stopped ")
		}()
	}()

	log.Println(<-errch)
}
