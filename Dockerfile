#syntax=docker/dockerfile:1.2

FROM golang:1.16-alpine AS go-build-env
RUN mkdir -p gunk services
COPY ./gunk/ /src/gunk
COPY ./services/ /src/services
COPY ./go.mod /src
COPY ./go.sum /src
COPY main.go /src

RUN mkdir vendor
COPY ./vendor /src/vendor

WORKDIR /src

RUN --mount=type=cache,target=/root/go/ go build -mod=vendor -o /hello

FROM alpine:3.13
COPY --from=go-build-env /hello /src/hello

WORKDIR /src

ENTRYPOINT [ "sh", "-c", "./hello" ]