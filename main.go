package main

import (
	"os"
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
		x -= len(driver.Loads)
		Armada.Drivers = append(Armada.Drivers, driver)
	}
}

func Generate_Route_for_New_Driver() (Driver, int) {
	driver := Init_Driver()
	// find first load
	init_load, _ := find_closest_load_from_depot()
	// if no loads exist then
	if init_load == -1 {
		return driver, -1
	}
	driver.Route = append(driver.Route, Loads[init_load].Number)
	driver.Loads = append(driver.Loads, *Loads[init_load])
	driver.Time_left += Loads[init_load].Distance_from_depot + Loads[init_load].Distance
	Loads[init_load].Complete()
	driver.Curr_Position = Loads[init_load].Dropoff
	// if driver and take another load get next load
	next_load, distance_to_pickup := driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
	//if no next load is avaiable to return
	if next_load == -1 {
		driver.Time_left += driver.Loads[len(driver.Loads)-1].Return_distance
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
		// if no load can fit into their schedule or there is no more loads left then break from this and return to the depot
		if next_load == -1 {
			break
		}
	}
	// return driver to depot
	if driver.Time_left+driver.Loads[len(driver.Loads)-1].Return_distance > 720.0 {
		//remove the load from schedule
		Loads[driver.Loads[len(driver.Loads)-1].Number].UnComplete()
		if len(driver.Loads) > 0 {
			driver.Loads = driver.Loads[:len(driver.Loads)-1]
		}
	}
	// return driver to depot
	driver.Time_left += driver.Loads[len(driver.Loads)-1].Return_distance
	return driver, 0
}

func find_closest_load_from_depot() (int, float64) {
	min_dist := -1.0
	load_num := -1
	for i := 0; i <= len(Loads)-1; i++ {
		if Loads[i].Completed {
			continue
		}
		if min_dist < 0 && !Loads[i].Completed {
			min_dist = Loads[i].Distance_from_depot
			load_num = i
			continue
		}
		if Loads[i].Distance_from_depot < min_dist && !Loads[i].Completed {
			min_dist = Loads[i].Distance_from_depot
			load_num = i
		}
	}
	return load_num, min_dist
}
