FROM golang:alpine AS builder

ADD . /root

WORKDIR /root
RUN go build -o argus-stream-engine-service

FROM alpine AS runner

RUN apk add gcompat

COPY --from=builder /root/argus-stream-engine-service /usr/bin/argus-stream-engine-service

EXPOSE 1935
EXPOSE 443

ENTRYPOINT ["/usr/bin/argus-stream-engine-service"]
