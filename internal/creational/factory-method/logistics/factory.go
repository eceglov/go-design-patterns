package logistics

type ITransportFactory interface {
	NewTransport() (ITransport, error)
}

type PlaneFactory struct{}
type ShipFactory struct{}
type TruckFactory struct{}
type SubmarineFactory struct{}

func (pf *PlaneFactory) NewTransport() (ITransport, error) {
	return NewPlane(1, "Twitter", "TW22010", 200)
}

func (sf *ShipFactory) NewTransport() (ITransport, error) {
	return NewShip(1, "Elena", "EL1010", 1000)
}

func (tf *TruckFactory) NewTransport() (ITransport, error) {
	return NewTruck(1, "Trucker", "TR90013", 10)
}

func (sf *SubmarineFactory) NewTransport() (ITransport, error) {
	return NewSubmarine(1, "K150 Alligator", "K150A112", 350)
}
