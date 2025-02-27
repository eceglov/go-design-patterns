package transport

import "log"

type ITransport interface {
	Deliver() error
}

type Transport struct {
	Id           int
	Name         string
	Registration string
	Capacity     int
}

func (t *Transport) Deliver() error {
	log.Printf("Deliver() method is not implemented by the transport: %s", t.Name)

	return nil
}
