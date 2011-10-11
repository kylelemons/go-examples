package main

import (
	"rand"
)

type Child struct {
	daycare   *Daycare
	activity  int
	attention int64

	move chan int
}

func (dc *Daycare) newChild(activity int) *Child {
	c := &Child{
		daycare: dc,
		activity: activity,
		attention: 1*Hour/2 + rand.Int63n(2*Hour),
	}
	go c.play()
	return c
}

func (c *Child) play() {
}
