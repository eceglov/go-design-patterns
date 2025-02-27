package transport

import "log"

type Ship struct {
	Transport
}

func NewShip(id int, name string, registration string, capacity int) (ITransport, error) {
	return &Ship{
		Transport{
			Id:           id,
			Name:         name,
			Registration: registration,
			Capacity:     capacity,
		},
	}, nil
}

func (t *Ship) Deliver() error {
	born := "sea"
	log.Printf("Delivering goods by %s: id: %d, name: %s, registration: %s, capacity: %d",
		born, t.Id, t.Name, t.Registration, t.Capacity)

	return nil
}
