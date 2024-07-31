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

func Reduce_Drivers() bool {
	var reduction_index []int
	for i := 0; i <= len(Armada.Drivers)-1; i++ {
		if len(Armada.Drivers[i].Loads) == 1 {
			if Insert_Route_To_New_Driver(Armada.Drivers[i].Loads[0], i+1) {
				// fmt.Println("remve driver =", i)
				// Armada.Drivers = append(Armada.Drivers[:i], Armada.Drivers[i:]...)
				reduction_index = append(reduction_index, i)
			}
		} else {
			continue
		}
	}
	reduced_index := 0
	for i := 0; i <= len(reduction_index)-1; i++ {
		Armada.Drivers = Armada.Drivers.Remove_driver(reduction_index[i] - reduced_index)
		reduced_index++
	}
	// fmt.Println("reduction index = ", reduction_index)
	// Armada.Drivers = Armada.Drivers.Remove_drivers(reduction_index)
	return true
}

func (f Fleet) Permute_Schedule() {
	for i := 0; i <= f.Drivers.Len()-1; i++ {
		f.Drivers[i] = f.Drivers[i].Swap_Schedule()
	}
}

func (d Driver) Swap_Schedule() Driver {
	initial_cost := d.Get_Cost()
	prev_cost := d.Get_Cost()
	for i := 0; i <= len(d.Loads)-1; i++ {
		vary_loads := make(Load_Schedule, len(d.Loads))
		copy(vary_loads, d.Loads)
		for j := 0; j <= len(d.Loads)-1; j++ {
			vary_loads[i], vary_loads[j] = vary_loads[j], vary_loads[i]
			if vary_loads.Load_Cost() < initial_cost && vary_loads.Load_Cost() < prev_cost {
				// fmt.Println("cheaper arrangment")
				prev_cost = vary_loads.Load_Cost()
				copy(d.Loads, vary_loads)
			}
		}
	}
	return d
}

// func Recursive_Permute()

/*
	Driver

is a struct with a slice of loads and time left and current position
*/
type Driver struct {
	Loads         Load_Schedule
	Time_left     float64
	Curr_Position Point
	Route         []int
}

type Drivers []Driver

func (d Drivers) Remove_first() Drivers {
	return d[1:]
}

func (d Drivers) Remove_last() Drivers {
	return d[:len(d)-1]
}

func (d Driver) Get_Cost() float64 {
	cost := 500.0
	cost += d.Loads[0].Distance_from_depot
	for i := 0; i <= len(d.Loads)-2; i++ {
		cost += d.Loads[i].Distance
		cost += d.Loads[i].Dropoff_to_pickup_dists[d.Loads[i+1].Number]
	}
	cost += d.Loads[len(d.Loads)-1].Return_distance + d.Loads[len(d.Loads)-1].Distance
	// fmt.Println("cost for driver = ", cost)
	if cost == d.Time_left+d.Loads[len(d.Loads)-1].Return_distance {
		fmt.Println("cost and time left are equal")
	}
	// fmt.Println("cost = ", cost)
	// fmt.Println("time left = ", d.Time_left+d.Loads[len(d.Loads)-1].Return_distance)
	return cost
}

func (d Drivers) Remove_driver(index int) Drivers {
	if index == 0 {
		return d.Remove_first()
	} else if index == len(d)-1 {
		return d.Remove_last()
	} else {
		return append(d[:index], d[index+1:]...)
	}
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

func (f Fleet) Print_Cost_Per_Driver() {
	for i := 0; i <= len(f.Drivers)-1; i++ {
		fmt.Println(f.Drivers[i].Get_Cost())
	}
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
