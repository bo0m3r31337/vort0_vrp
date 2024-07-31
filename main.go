package main

import (
	"os"
	"sort"
)

// slice of loads to complete
var Loads []*Load

// init Fleet
var Armada Fleet

// Home_Point is the (0,0) depot point
var Home_Point Point

func main() {
	// create fleet of drivers
	Armada = Fleet{
		Drivers:       make([]Driver, 0),
		Total_Minutes: 0,
	}
	// set home point
	Home_Point = Make_Home_Point()
	// get loads
	Get_loads(os.Args[1])
	// populate loads with dropoff to pickup distances
	Populate_Matrix()
	// solve Vehicle Routing Problem
	VRP_Solve()
	// sort and reduce drivers with 1 load if possible
	for x := 3; x != 0; x-- {
		sort.Sort(Armada.Drivers)
		Reduce_Drivers()
		// permute existing schedule if the schedule per driver is lower
		Armada.Permute_Schedule()
	}
	// PRINT TO STD OUT
	Armada.Print()
}

func VRP_Solve() {
	x := len(Loads)
	for x != 0 {
		driver, err := Generate_Route_for_New_Driver()
		if err != 0 {
			break
		}
		x -= len(driver.Route)
		Armada.Docked = append(Armada.Docked, driver)
		Armada.Drivers = append(Armada.Drivers, driver)
		sort.Sort(Armada.Drivers)
		Reduce_Drivers()
	}
}

func Generate_Route_for_New_Driver() (Driver, int) {
	driver := Init_Driver()
	// find first load
	for i := 0; i <= len(Loads)-1; i++ {
		if Loads[i].Completed {
			continue
		}
		driver.Route = append(driver.Route, Loads[i].Number)
		driver.Loads = append(driver.Loads, *Loads[i])
		driver.Time_left += Loads[i].Distance_from_depot + Loads[i].Distance
		Loads[i].Complete()
		driver.Curr_Position = Loads[i].Dropoff
		break
	}
	if len(driver.Loads) == 0 {
		return driver, -1
	}
	// if driver and take another load get next load
	next_load, distance_to_pickup := driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
	// if no load is available return
	if next_load == -1 {
		return driver, 0
	}
	// else add more loads
	for driver.Time_left+distance_to_pickup+Loads[next_load].Return_distance+Loads[next_load].Distance < 720 {
		driver.Route = append(driver.Route, Loads[next_load].Number)
		driver.Loads = append(driver.Loads, *Loads[next_load])
		driver.Time_left += Loads[next_load].Distance + distance_to_pickup
		Loads[next_load].Complete()
		driver.Curr_Position = Loads[next_load].Dropoff
		next_load, distance_to_pickup = driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
		// if no load exists break and return
		if next_load == -1 {
			break
		}
	}
	return driver, 0
}
