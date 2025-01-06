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
make run
```

stream to the service:
```
make stream
```
ingest the stream with VLC:
    1. Media > Open Network Stream 

    2. Use `rtmp://localhost/outgoing/myRestream` as the network URL