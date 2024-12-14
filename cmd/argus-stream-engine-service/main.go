package main

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
)

func main() {
	rtmp.LogInfo("Initializing Argus Stream Engine...")

	config := NewConfig()
	if config.WithRelay {
		rtmp.LogInfo("Relay enabled.")
		rtmp.LogInfo("Relay Ingress Address: " + config.RelayIngressAddress)
		rtmp.LogInfo("Relay Egress Address: " + config.RelayEgressAddress)
		rtmp.LogInfo("Egress Destination Address: " + config.EgressDestinationAddress)

	}

	server := rtmp.NewServer()
	if server != nil {
		server.Start()
	}

	rtmp.LogInfo("Argus Stream Engine Online.")
}
