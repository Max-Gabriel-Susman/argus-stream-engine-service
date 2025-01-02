# Argus Stream Engine Service

The Argus Stream Engine Service manages RTMP streams between ingress and egress clients. 

## Usage 

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
go run main.go
```

clear binaries:
```
make clean
```