package stream

import "fmt"

type Ingress struct {
}

type Destination struct {
	DestinationAddress      string
	DestinationResponsePath string
}

type Egress struct {
	EgressDestination Destination
}

type Relay struct {
	RelayIngress     Ingress
	RelayDestination Destination
	RelayEgress      Egress
}

func NewRelay(destinationAddress string) Relay {
	return Relay{
		RelayDestination: Destination{
			DestinationAddress: destinationAddress,
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
