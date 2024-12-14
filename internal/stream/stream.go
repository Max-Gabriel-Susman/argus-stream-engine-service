package stream

import "fmt"

type Ingress struct {
}

// Separate ingress and egress address enables pipelining image processing services.
//
// Destination encapsulates the fields used to interact with a streaming destination.
type Destination struct {
	DestinationIngressAddress string
	DestinationEgressAddress  string
	DestinationResponsePath   string
}

type Egress struct {
	EgressDestination Destination
}

type Relay struct {
	RelayIngress     Ingress
	RelayDestination Destination
	RelayEgress      Egress
}

// TODO: we're gonna need to parameterize these address args with environment variables.
//
// NewRelay Constructs a new Relay.
func NewRelay(destinationIngressAddress, destinationEgressAddress string) Relay {
	return Relay{
		RelayDestination: Destination{
			DestinationIngressAddress: destinationIngressAddress,
			DestinationEgressAddress:  destinationEgressAddress,
		},
	}
}

func (r Relay) IngestStream() error {
	return nil
}

func (r Relay) EgestStream() error {
	return nil
}

func (r Relay) HandleStreamRelay() error {
	if err := r.IngestStream(); err != nil {
		return fmt.Errorf("failure to listen to relay response: %w", err)
	}

	if err := r.EgestStream(); err != nil {
		return fmt.Errorf("failure to initiate relay: %w", err)
	}

	return nil
}
