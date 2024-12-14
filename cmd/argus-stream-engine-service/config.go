package main

import "os"

// Config encapsulates the fields used to configure the service.
type Config struct {
	WithRelay                string
	RelayIngressAddress      string
	RelayEgressAddress       string
	EgressDestinationAddress string
}

// NewConfig constructs a new Configuration for the service.
func NewConfig() Config {
	return Config{
		WithRelay:                os.Getenv("WITH_RELAY"),
		RelayIngressAddress:      os.Getenv("RELAY_INGRESS_ADDRESS"),
		RelayEgressAddress:       os.Getenv("RELAY_EGRESS_ADDRESS"),
		EgressDestinationAddress: os.Getenv("EGRESS_DESTINATION_ADDRESS"),
	}
}
