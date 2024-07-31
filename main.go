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
	for x := 3; x != 0; x-- {
		sort.Sort(Armada.Drivers)
		Reduce_Drivers()
		Armada.Permute_Schedule()
	}
	Armada.Print()
	// Armada.Print_Cost_Per_Driver()

}

func VRP_Solve() {
	// loads := make([]*Load, len(Loads))
	// _ = copy(loads, Loads)
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
		// Armada.Permute_Schedule()
	}
}

func VRP_Solve_v2() {
	// loads := make([]*Load, len(Loads))
	// _ = copy(loads, Loads)
	x := len(Loads)
	// var fleet_size_min int
	for x != 0 {
		driver, err := Dispatch_New_Driver()
		if err != 0 {
			break
		}
		driver, err = driver.Driver_Continue_Lowest_Delta()
		if err != 0 {
			break
		}
		x -= len(driver.Route)
		// fmt.Println("stuck in this loop1")
		Armada.Docked = append(Armada.Docked, driver)
		Armada.Drivers = append(Armada.Drivers, driver)
		sort.Sort(Armada.Drivers)
		Reduce_Drivers()
		Armada.Permute_Schedule()
	}
	// fleet_size_min = len(Armada.Drivers)
	previous_fleet_size := Armada.Drivers.Len()
	fleet_size := 0
	for fleet_size != previous_fleet_size {
		sort.Sort(Armada.Drivers)
		previous_fleet_size = Armada.Drivers.Len()
		// fmt.Println("previous fleet size = ", previous_fleet_size)
		Reduce_Drivers_V2()
		// Find_Driver_and_Route_Schedule()
		// fmt.Println("fleet size after reduction = ", fleet_size)
		fleet_size = Armada.Drivers.Len()
		// fmt.Println("stuck in this loop2")
	}

}

func Calculate_New_Route_For_Existing_Driver(driver Driver) (Driver, int) {
	// if driver and take another load get next load
	next_load, distance_to_pickup := driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
	// if no load available return
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

func (driver Driver) Driver_Continue_Lowest_Delta() (Driver, int) {
	next_load, distance_to_pickup := driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
	if next_load == -1 {
		return driver, -1
	}
	for driver.Time_left+distance_to_pickup+Loads[next_load].Return_distance+Loads[next_load].Distance < 720 {
		new_route, delta := driver.Search_Driver(Loads[next_load])
		if new_route == nil {
			driver.Route = append(driver.Route, Loads[next_load].Number)
			driver.Loads = append(driver.Loads, *Loads[next_load])
			driver.Time_left += Loads[next_load].Distance + distance_to_pickup
			Loads[next_load].Complete()
			driver.Curr_Position = Loads[next_load].Dropoff

		} else {
			driver.Loads = new_route
			driver.Time_left += delta
			Loads[next_load].Complete()
			driver.Curr_Position = new_route[len(new_route)-1].Dropoff
		}
		next_load, distance_to_pickup = driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
		// if no load exists break and return
		if next_load == -1 {
			break
		}
	}
	// new_route := driver.Search_Driver(Loads[next_load])
	return driver, 0
}

func Dispatch_New_Driver() (Driver, int) {
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
	if driver.Time_left == 0.0 {
		return driver, -1
	}
	return driver, 0
}

// func (d Driver) Best_Insertion() {

// }
