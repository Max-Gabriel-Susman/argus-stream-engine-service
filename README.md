# Argus Stream Engine Service

The Argus Stream Engine Service Transforms incoming RTMP streams from ingress clients into HLS streams for egress clients. 

## Local Setup

*.mp4 files may be included in the videos directory

## Docker 

You can find the docker image for this project available in Docker Hub: [https://hub.docker.com/r/brometheus/argus-stream-engine-service](https://hub.docker.com/r/brometheus/argus-stream-engine-service)

pull the service image:
```
docker pull brometheus/argus-stream-engine-service
```

run the image as a container:
```
docker run -p 1935:1935 brometheus/argus-stream-engine-service
```

## Kubernetes Usage

deploy 
```
kubectl apply -f argus-stream-engine-service-deployment.yaml
```

## Connecting To the Service

Argus Stream Engine Service accepts connections formatted like: rtmp://{HOST}/{CHANNEL}/{KEY}

to use an input mp4 from your fileystem run: 
```
ffmpeg -re -i input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key
```
