package hello

import (
	"context"
	"fmt"
	"log"
	"time"

	hellopb "github.com/ghazninattarshah/tracing-with-jaeger/gunk/hello"
	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

type HelloService struct {
	Service     string
	tracer      opentracing.Tracer
	nextService hellopb.HelloServiceClient
	hellopb.UnimplementedHelloServiceServer
}

func NewHelloService(serviceID string, tracer opentracing.Tracer, nextService hellopb.HelloServiceClient) *HelloService {
	return &HelloService{Service: serviceID, tracer: tracer, nextService: nextService}
}

// SayHello returns the passed message from respective service
func (s *HelloService) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span := s.tracer.StartSpan(s.Service+":SayHello", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if req.ServiceID.String() != s.Service && span != nil && s.nextService != nil {
		log.Println("proceeding to next service")
		nextServiceTag := hellopb.ServiceID_TWO.String()
		time.Sleep(time.Millisecond * 100)
		if s.Service == hellopb.ServiceID_TWO.String() {
			time.Sleep(time.Millisecond * 200)
			nextServiceTag = hellopb.ServiceID_THREE.String()
		}
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, nextServiceTag)
		span.SetTag("next-service", nextServiceTag)
		return s.nextService.SayHello(ctx, req)
	}

	return &hellopb.HelloResponse{
		Msg: fmt.Sprintf("Hello from service %s", s.Service),
	}, nil
}
