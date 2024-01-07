package main

import (
	"fmt"
)

// Train is a type in Go that extends Vehicle and implements FuelDependent
type Train struct {
	Vehicle    // Embedding Vehicle to extend its fields and methods
	wagonCount int
}

// Implementing move method for Train
func (t *Train) move() {
	fmt.Println("Train is moving...")
}

// PassengerTrain is a type in Go that extends Train
type PassengerTrain struct {
	Train          // Embedding Train to extend its fields and methods
	passengerCount int
}

// Implementing refuel method for PassengerTrain (required by FuelDependent interface)
func (pt *PassengerTrain) refuel(liter float32) {
	// Implement refuel logic here
}

// Override move method for PassengerTrain
func (pt *PassengerTrain) move() {
	fmt.Println("Passenger train is moving...")
}

// FreightTrain is a type in Go that extends Train
type FreightTrain struct {
	Train         // Embedding Train to extend its fields and methods
	carryCapacity float32
}

// Implementing refuel method for FreightTrain (required by FuelDependent interface)
func (ft *FreightTrain) refuel(liter float32) {
	// Implement refuel logic here
}

// Override move method for FreightTrain
func (ft *FreightTrain) move() {
	fmt.Println("Freight train is moving...")
}
