package logistics

import "log"

type Truck struct {
	Transport
}

func NewTruck(id int, name string, registration string, capacity int) (ITransport, error) {
	return &Truck{
		Transport{
			Id:           id,
			Name:         name,
			Registration: registration,
			Capacity:     capacity,
		},
	}, nil
}

func (t *Truck) Deliver() error {
	born := "ground"
	log.Printf("Delivering goods by %s: id: %d, name: %s, registration: %s, capacity: %d",
		born, t.Id, t.Name, t.Registration, t.Capacity)

	return nil
}
