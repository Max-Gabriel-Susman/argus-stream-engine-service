FROM golang:alpine AS builder

# Copy the entire project into the container
ADD . /app

# Set the working directory to where your main Go file is located
WORKDIR /app/cmd/argus-stream-engine-service

# Build the Go application
RUN go build -o /argus-stream-engine-service

FROM alpine AS runner

# Install necessary runtime dependencies
RUN apk add --no-cache gcompat

# Copy the built binary from the builder stage
COPY --from=builder /argus-stream-engine-service /usr/bin/argus-stream-engine-service

# Expose the ports used by your service
EXPOSE 1935
EXPOSE 443

# Set the entry point to your application
ENTRYPOINT ["/usr/bin/argus-stream-engine-service"]
