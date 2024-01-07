package main

type TrafficJam[T any] struct {
	queue []T
}

func (t *TrafficJam[T]) addVehicle(vehicle T) {
	t.queue = append(t.queue, vehicle)
}

func (t *TrafficJam[T]) jamLength() int {
	return len(t.queue)
}
