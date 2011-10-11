package main

import (
	"rand"
	"fmt"
)

type Daycare struct{
	activities   []string
	drop, pickup chan string
	count        chan chan int
	role         chan bool
}

func NewDaycare(activities ...string) *Daycare {
	dc := &Daycare{
		activities: activities,
		drop:   make(chan string),
		pickup: make(chan string),
		count:  make(chan chan int),
		role:   make(chan bool),
	}
	go dc.run()
	return dc
}

func (dc *Daycare) run() {
	children := map[string]*Child{}

	for {
		select {
		case child := <-dc.drop:
			act := rand.Intn(len(dc.activities))
			children[child] = dc.newChild(act)
		case child := <-dc.pickup:
			children[child] = nil, false
		case count := <-dc.count:
			count <- len(children)
		case <-dc.role:
			for name, child := range children {
				fmt.Printf("%5s: %s\n", name, dc.activities[child.activity])
			}
		}
	}
}

func (dc *Daycare) DropOff(children ...string) {
	for _, child := range children {
		dc.drop <- child
	}
}

func (dc *Daycare) PickUp(children ...string) {
	for _, child := range children {
		dc.pickup <- child
	}
}

func (dc *Daycare) RoleCall() {
	dc.role <- true
}

func (dc *Daycare) Count() int {
	ret := make(chan int)
	dc.count <- ret
	return <-ret
}
