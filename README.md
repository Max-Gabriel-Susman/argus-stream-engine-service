# Argus RTMP Service

The Argus RTMP Service receives incoming RTMP streams from Argus Ingress Clients

## Docker 

You can find the docker image for this project available in Docker Hub: [https://hub.docker.com/r/brometheus/argus-rtmp-service](https://hub.docker.com/r/brometheus/argus-rtmp-service)

To pull it type:

```
docker pull brometheus/argus-rtmp-service

## Local Usage

Argus RTMP Service accepts connections formatted like:

```
rtmp://{HOST}/{CHANNEL}/{KEY}
```

## Example Usage

docker run -p 1935:1935 brometheus/argus-rtmp-service

ffmpeg -re -i input.mp4 -c:v libx264 -preset fast -b:v 1000k -maxrate 1000k -bufsize 2000k -c:a aac -ar 44100 -b:a 128k -f flv rtmp://localhost/live/stream_key

## TODOs

* we need to finish implementing hsls service

