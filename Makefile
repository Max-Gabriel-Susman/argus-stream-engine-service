
deps: 
	go get "github.com/Max-Gabriel-Susman/rtmp@v0.1.0"
	go mod tidy
	go mod vendor
	git submodule update --init --recursive


build:
	gcc -c -o internal/detection/detector.o internal/detection/libdetector/detector.c
	ar rcs internal/detection/libdetector.a internal/detection/detector.o
	gcc -c -o internal/stream/pipeline.o internal/stream/libpipeline/pipeline.c
	ar rcs internal/stream/libpipeline.a internal/stream/pipeline.o

clean: 
	rm internal/stream/pipeline.o \
	internal/stream/libpipeline.a \
	internal/detection/detector.o \
	internal/detection/libdetector.a
	rm -r vendor
	rm go.sum

run: 
	go run cmd/argus-stream-engine-service/main.go

yolo:
	make clean 
	make deps
	make build 
	make run

test: 
	go test ./...

stream: 
	ffmpeg -re -i imagery/input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key

# build: 
# 	docker build -t argus-stream-engine-service .

# run: 
# 	docker run -p 1935:1935 argus-stream-engine-service

tag: 
	docker tag argus-stream-engine-service brometheus/argus-stream-engine-service:latest

push: 
	docker push brometheus/argus-stream-engine-service:latest
