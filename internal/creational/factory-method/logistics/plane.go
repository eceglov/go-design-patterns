package logistics

import "log"

type Plane struct {
	Transport
}

func NewPlane(id int, name string, registration string, capacity int) (ITransport, error) {
	return &Plane{
		Transport{
			Id:           id,
			Name:         name,
			Registration: registration,
			Capacity:     capacity,
		},
	}, nil
}

func (t *Plane) Deliver() error {
	born := "air"
	log.Printf("Delivering goods by %s: id: %d, name: %s, registration: %s, capacity: %d",
		born, t.Id, t.Name, t.Registration, t.Capacity)

	return nil
}
