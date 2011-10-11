package main

import (
	"time"
)

const Hour = 10e6 // 1 hour = 10ms

func main() {
	dc := NewDaycare("Arts & Crafts", "Playground", "Reading", "Learning")
	dc.DropOff("Billy", "Carla", "Joe", "Amy", "Sam")

	go func() {
		<-time.After(2*Hour)
		dc.DropOff("Pam")
	}()

	go func() {
		<-time.After(4*Hour)
		//dc.LunchTime()
	}()

	go func() {
		<-time.After(6*Hour)
		dc.PickUp("Billy", "Carla")
	}()

	go func() {
		<-time.After(7*Hour)
		dc.PickUp("Sam")
	}()

	go func() {
		<-time.After(8*Hour)
		dc.PickUp("Amy", "Pam")
	}()

	go func() {
		<-time.After(9*Hour)
		dc.PickUp("Joe")
	}()

	dc.RoleCall()
	for dc.Count() > 0 {
		<-time.After(1*Hour + 1*Hour/5)
		dc.RoleCall()
	}
}
