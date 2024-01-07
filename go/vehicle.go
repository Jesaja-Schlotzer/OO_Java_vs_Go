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
