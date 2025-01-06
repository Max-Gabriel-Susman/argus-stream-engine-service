deps:
	sudo apt-get update
	sudo apt-get install -y software-properties-common
	sudo add-apt-repository universe
	sudo apt-get update
	sudo apt-get install -y build-essential
	sudo apt-get install -y \
		gstreamer1.0-tools \
		gstreamer1.0-plugins-base \
		gstreamer1.0-plugins-good \
		gstreamer1.0-plugins-bad \
		gstreamer1.0-plugins-ugly \
		ffmpeg \
		libnginx-mod-rtmp
	sudo apt-get install -y \
		libunwind-dev \
		libgstreamer1.0-dev \
		libgstreamer-plugins-base1.0-dev

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
