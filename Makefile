

stream: 
	ffmpeg -re -i input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key

build: 
	docker build -t argus-rtmp-service .

run: 
	docker run argus-rtmp-service

tag: 
	docker tag argus-rtmp-service brometheus/argus-rtmp-service:latest

push: 
	docker push brometheus/argus-rtmp-service:latest