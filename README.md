# Argus Stream Engine Service

The Argus Stream Engine Service manages RTMP streams between ingress and egress clients. 

## Local Setup and Usage

install VLC @ https://www.videolan.org/vlc/download-macosx.html

install ffmpeg
```
brew install ffmpeg
```

It should be noted that an input.mp4 file needs to be included in the videos directory for the included make targets to work properly(see videos/overview.md).

tidy the dependencies:
```
go mod tidy
```

vendor the dependencies: 
```
go mod vendor
```

Start the service: 
```
go run cmd/argus-stream-engine-service/main.go
```

At this point open a second terminal, from the root level of this repository change directory into videos and make stream
```
cd videos && make stream
```

NOTE: It's very important that you do this next step while the stream operation from the previous step is still going, other wise there will be nothing for the media player to receive. 

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
