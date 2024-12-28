# Argus Stream Engine Service

The Argus Stream Engine Service manages RTMP streams between ingress and egress clients. 

## Usage 

install dependencies:
```
make deps
```

build pipeline library binaries:
```
make build
```

run application: 
```
go run main.go
```

clear binaries:
```
make clean
```

## TODOs

1. Upgrade build system to a CMake Implementation

2. Dockerize the service

3. Provide instructions for deployment on and consumption from EKS