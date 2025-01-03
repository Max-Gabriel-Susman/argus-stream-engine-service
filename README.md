# Argus Stream Engine Service

The Argus Stream Engine Service manages RTMP streams between ingress and egress clients.

## Build and run As a Container

build docker image for the service:
```
docker build -t argus-stream-engine-service .
```

run docker container for the service:
```
docker run --name argus-stream-engine-service -p 1935:1935 argus-stream-engine-service
```

## Build and run from Source 

install dependencies:
```
make deps 
```

NOTE: The desired nginx.conf is included at the root level of this repository for reference; check or edit /etc/nginx/nginx.conf accordingly. Then:
```
make prep
```

build pipeline library binaries:
```
make build
```

run application: 
```
make run
```

clear binaries:
```
make clean
```