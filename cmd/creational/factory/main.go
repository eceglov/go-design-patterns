package main

import (
	"github.com/eceglov/go-design-patterns.git/internal/creational/factory/transport"

	"github.com/gookit/goutil/dump"
	"log"
)

func main() {
	var transportFactory transport.ITransportFactory
	var transportation []transport.ITransport

	transportFactory = &transport.ShipFactory{}
	ship, _ := transportFactory.NewTransport()
	if err := ship.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}
	transportation = append(transportation, ship)

	transportFactory = &transport.PlaneFactory{}
	plane, _ := transportFactory.NewTransport()
	if err := plane.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}
	transportation = append(transportation, plane)

	transportFactory = &transport.TruckFactory{}
	truck, _ := transportFactory.NewTransport()
	if err := truck.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}
	transportation = append(transportation, truck)

	transportFactory = &transport.SubmarineFactory{}
	submarine, _ := transportFactory.NewTransport()
	if err := submarine.Deliver(); err != nil {
		log.Println("Failed to deliver:", err)
	}
	transportation = append(transportation, submarine)

	dump.P(transportation)
}
