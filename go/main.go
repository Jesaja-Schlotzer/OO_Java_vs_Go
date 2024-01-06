package main

import (
	"fmt"
	"time"
)

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

	jam := TrafficJam[Car]{}
	jam.addVehicle(car)

	fmt.Printf("Top Speed: %f\n", vehicle.topSpeed)
}
