package main

import (
	"github.com/eceglov/go-design-patterns.git/internal/creational/factory-method/logistics"

	"github.com/gookit/goutil/dump"
	"log"
)

func main() {
	var logisticsFactory logistics.ITransportFactory
	var transport []logistics.ITransport

	logisticsFactory = &logistics.ShipFactory{}
	ship, _ := logisticsFactory.NewTransport()
	if err := ship.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}

	logisticsFactory = &logistics.PlaneFactory{}
	plane, _ := logisticsFactory.NewTransport()
	if err := plane.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}

	logisticsFactory = &logistics.TruckFactory{}
	truck, _ := logisticsFactory.NewTransport()
	if err := truck.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}

	logisticsFactory = &logistics.SubmarineFactory{}
	submarine, _ := logisticsFactory.NewTransport()
	if err := submarine.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}

	transport = append(transport, ship, plane, truck, submarine)

	dump.P(transport)
}
