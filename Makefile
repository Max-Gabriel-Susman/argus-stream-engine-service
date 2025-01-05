
deps: 
	sudo apt-get update
	sudo apt-get install -y gstreamer1.0-tools gstreamer1.0-plugins-base \
		gstreamer1.0-plugins-good gstreamer1.0-plugins-bad \
		gstreamer1.0-plugins-ugly
	sudo apt-get install -y ffmpeg
	sudo apt-get install -y libnginx-mod-rtmp

build: 
	gcc -c -o gst_pipeline.o gst_pipeline.c `pkg-config --cflags gstreamer-1.0`
	ar rcs libgst_pipeline.a gst_pipeline.o

clean: 
	rm gst_pipeline.o libgst_pipeline.a

run: 
	go run main.go

test: 
	go test ./...

stream:
	ffmpeg -re -i imagery/input.mp4 \
    -c:v libx264 -preset veryfast -c:a aac -ar 44100 \
    -f flv rtmp://localhost/outgoing/myRestream
