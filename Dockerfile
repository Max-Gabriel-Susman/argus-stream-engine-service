FROM golang:alpine AS builder

ADD . /root

WORKDIR /root
RUN go build -o argus-rtmp-service

FROM alpine AS runner

RUN apk add gcompat

COPY --from=builder /root/argus-rtmp-service /usr/bin/argus-rtmp-service

EXPOSE 1935
EXPOSE 443

ENTRYPOINT ["/usr/bin/argus-rtmp-service"]
