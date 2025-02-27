package transport

type Submarine struct {
	Transport
}

func NewSubmarine(id int, name string, registration string, capacity int) (ITransport, error) {
	return &Submarine{
		Transport{
			Id:           id,
			Name:         name,
			Registration: registration,
			Capacity:     capacity,
		},
	}, nil
}

//// Deliver is intentionally not implemented method to showcase that the Deliver() method from
//// common transport is called instead if not overridden in a specific transport implementation.
//func (t *Submarine) Deliver() error {
//	born := "undersea"
//	log.Printf("Delivering goods by %s: id: %d, name: %s, registration: %s, capacity: %d",
//		born, t.Id, t.Name, t.Registration, t.Capacity)
//
//	return nil
//}
