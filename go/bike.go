package main

import (
	"fmt"
)

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
