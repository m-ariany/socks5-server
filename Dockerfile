ARG GOLANG_VERSION="1.19.1"

FROM golang:$GOLANG_VERSION-alpine as builder
RUN apk --no-cache add tzdata
WORKDIR /proxy
COPY *.go /proxy
COPY go.* /proxy/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./socks5

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /proxy/socks5 /
ENTRYPOINT ["/socks5"]
