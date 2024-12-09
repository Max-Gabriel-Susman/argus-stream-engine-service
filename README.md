# Argus Stream Engine Service

The Argus Stream Engine Service manages RTMP streams between ingress and egress Argus clients. 

## Local Setup

install VLC @ https://www.videolan.org/vlc/download-macosx.html

install ffmpeg
```
brew install ffmpeg
```

It should be noted that an input.mp4 file needs to be included in the videos directory for the included make targets to work properly(see videos/overview.md).

Start the service: 
```
go run cmd/argus-stream-engine-service/main.go
```

Change directory into videos and make stream
```
cd videos && make stream
```

Then in VLC open network for URL: rtmp://localhost/live/stream_key

## Docker 

You can find the docker image for this project available in Docker Hub: [https://hub.docker.com/r/brometheus/argus-stream-engine-service](https://hub.docker.com/r/brometheus/argus-stream-engine-service)

pull the service image:
```
make pull
```

run the image as a container:
```
make run
```
