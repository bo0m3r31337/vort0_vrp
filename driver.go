package main

import (
	"fmt"
)

/*
Fleet
is a struct of total distances in driving minutes between pickup dropoff loads
and a slice of active drivers
*/
type Fleet struct {
	Drivers       Drivers
	Total_Minutes float64
	Docked        []Driver
}

/*
	Driver

is a struct with a slice of loads and time left and current position
*/
type Driver struct {
	Loads         []Load
	Time_left     float64
	Curr_Position Point
	Route         []int
}

type Drivers []Driver

func (d Drivers) Remove_first() Drivers {
	return d[1:]
}

func (d Drivers) Len() int {
	return len(d)
}

func (d Drivers) Less(i, j int) bool {
	return d[i].Time_left < d[j].Time_left
}

func (d Drivers) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Init_Driver creates a driver struct
func Init_Driver() Driver {
	return Driver{
		Loads:         make([]Load, 0),
		Time_left:     0.0,
		Curr_Position: Home_Point,
		Route:         make([]int, 0),
	}
}

// Gets cost of fleet and driving time
func (f Fleet) get_cost() float64 {
	return f.Total_Minutes + (500.00 * float64(len(f.Drivers)))
}

func (f Fleet) Print() {
	for i := 0; i <= len(f.Drivers)-1; i++ {
		fmt.Print("[")
		for j := 0; j <= len(f.Drivers[i].Loads)-1; j++ {
			fmt.Print(f.Drivers[i].Loads[j].Number + 1)
			if j != len(f.Drivers[i].Loads)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]\n")
	}
}

func (d Driver) Calculates_Route_Time() float64 {
	// sort.Ints(d.Route)
	// d.Loads[1], d.Loads[2] = d.Loads[2], d.Loads[1]
	init := d.Loads[0].Distance_from_depot
	r_dist := Loads[len(d.Loads)-1].Return_distance
	travel := 0.0
	for i := 0; i <= len(d.Loads)-2; i++ {
		travel += d.Loads[i].Distance + d.Loads[i].Dropoff_to_pickup_dists[d.Route[i+1]]
	}
	travel += d.Loads[len(d.Loads)-1].Distance
	return init + travel + r_dist
}

func (f *Fleet) Retire_Driver() {
	for i := 0; i <= len(f.Drivers[0].Loads)-1; i++ {
		f.Drivers[0].Loads[i].UnComplete()
	}
	f.Drivers = f.Drivers.Remove_first()
}
