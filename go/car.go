package main

import (
	"fmt"
)

// Car is a type in Go that extends Vehicle and implements FuelDependent
type Car struct {
	Vehicle        // Embedding Vehicle to extend its fields and methods
	doorCount      int
	passengerCount int
	fuelLevel      float32
}

// Implementing refuel method for Car (required by FuelDependent interface)
func (c *Car) refuel(liter float32) {
	// Implement refuel logic here
	c.fuelLevel += liter
	fmt.Printf("Refueled. Current fuel level: %f liters\n", c.fuelLevel)
}

// Override move method for Car
func (c *Car) move() {
	fmt.Println("Car is moving...")
}
