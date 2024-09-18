#FROM golang:1.19.2 AS build
#FROM golang:1.20.5 AS build
FROM anolis-registry.cn-zhangjiakou.cr.aliyuncs.com/openanolis/golang:1.20.5 AS build

WORKDIR /go/src/tros_example_server
COPY . /go/src/tros_example_server

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -v \
    -o /bin/tros_example_server

FROM ubuntu:20.04

WORKDIR /tros_example_server

COPY --from=build /bin/tros_example_server /tros_example_server/
COPY conf/app.online.yaml /tros_example_server/conf/app.yaml
RUN mkdir -p runtime/logs

RUN apt-get update && \
    apt-get install -y ca-certificates

EXPOSE 9091
EXPOSE 9092

CMD ["./tros_example_server"]
