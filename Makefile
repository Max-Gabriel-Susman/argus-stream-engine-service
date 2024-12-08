

stream: 
	ffmpeg -re -i input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key

build: 
	docker build -t argus-stream-engine-service .

run: 
	docker run argus-stream-engine-service

tag: 
	docker tag argus-stream-engine-service brometheus/argus-stream-engine-service:latest

push: 
	docker push brometheus/argus-stream-engine-service:latest

test: 
	go test ./...