/*
*****************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

******************************************************************************
*/
package main

import (
	"fmt"
	"time"
)

type Vehicle struct {
	topSpeed          float32
	weight            float32
	manufacturingDate time.Time
}

// Move is an abstract method in Go
func (v *Vehicle) move() {
	fmt.Println("Moving...")
}

type FuelDependent interface {
	refuel(liter float32)
}

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

// Bike is a type in Go that extends Vehicle
type Bike struct {
	Vehicle     // Embedding Vehicle to extend its fields and methods
	gearCount   int
	currentGear int
}

// Override move method for Bike
func (b *Bike) move() {
	fmt.Println("Bike is moving...")
}

// ringBell is a method specific to Bike
func (b *Bike) ringBell() {
	fmt.Println("Ring, ring!")
}

// gearShiftUp is a method to shift gears up in Bike
func (b *Bike) gearShiftUp() {
	b.currentGear++
}

// gearShiftUpWithCount is a method to shift gears up by a specified count in Bike
func (b *Bike) gearShiftUpWithCount(gearCount int) {
	b.currentGear += gearCount
}

func main() {
	var vehicle = Vehicle{
		300.0,
		2500.0,
		time.Now(),
	}

	var car = Car{
		Vehicle: Vehicle{
			topSpeed:          160,
			weight:            2000,
			manufacturingDate: time.Now(),
		},
		doorCount:      4,
		passengerCount: 4,
		fuelLevel:      0,
	}

	fmt.Printf("Top Speed: %f\n", vehicle.topSpeed)
}
