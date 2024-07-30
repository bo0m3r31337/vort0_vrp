package main

import (
	"fmt"
	"os"
)

var Loads []*Load
var Load_Nums []int
var Armada Fleet
var Fleet_Size float64
var Distance_Matrix [][]float64

// Home_Point is the (0,0) depot point
var Home_Point Point

func main() {
	Armada = Fleet{
		Drivers:       make([]Driver, 0),
		Total_Minutes: 0,
	}
	Home_Point = Make_Home_Point()
	// fmt.Println(os.Args[1])
	// start_time := time.Now()
	Get_loads(os.Args[1])
	// Init_Matrix()
	Populate_Matrix()
	VRP_Solve()
	Armada.Print()
	// end_time := time.Now()
	// exec_time := end_time.Sub(start_time)
	// fmt.Println("exec time = ", exec_time)
}

func VRP_Solve() {
	// loads := make([]*Load, len(Loads))
	// _ = copy(loads, Loads)
	x := len(Loads)
	for x != 0 {
		driver, err := Generate_Route_for_Driver()
		if err != 0 {
			break
		}
		x -= len(driver.Route)
		Armada.Docked = append(Armada.Docked, driver)
		Armada.Drivers = append(Armada.Drivers, driver)
	}
}

func Generate_Route_for_Driver() (Driver, int) {
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
	if len(driver.Route) == 0 {
		return driver, -1
	}
	// // check if initial load distance exceeds drivers driving limit
	// if driver.Time_left+Distance_from_depot(driver.Curr_Position) > 720 {
	// 	return driver
	// }
	next_load, distance_to_pickup := driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
	if next_load == -1 {
		return driver, 0
	}
	// else add more loads
	for driver.Time_left+distance_to_pickup+Loads[next_load].Return_distance+Loads[next_load].Distance < 720 {
		// fmt.Println("driver time left = ", driver.Time_left)
		// find closest pickup
		driver.Route = append(driver.Route, Loads[next_load].Number)
		driver.Loads = append(driver.Loads, *Loads[next_load])
		driver.Time_left += Loads[next_load].Distance + distance_to_pickup
		Loads[next_load].Complete()
		driver.Curr_Position = Loads[next_load].Dropoff
		next_load, distance_to_pickup = driver.Loads[len(driver.Loads)-1].Return_min_Load_Not_Completed(driver.Time_left)
		if next_load == -1 {
			break
		}

	}
	// driver can't carry anymore non-completed loads
	// fmt.Println("driver route = ", driver.Route)
	return driver, 0

}

func Find_route_for_new_driver() Driver {
	driver := Init_Driver()
	for i := 0; i <= len(Loads)-1; i++ {
		fmt.Println(Distance_Matrix)
		if driver.Time_left == 0 && !Loads[i].Completed {
			driver.Route = append(driver.Route, Loads[i].Number)
			driver.Loads = append(driver.Loads, *Loads[i])
			driver.Time_left += Loads[i].Distance_from_depot + Loads[i].Distance
			Loads[i].Completed = true
			driver.Curr_Position = Loads[i].Dropoff
		} else if Loads[i].Completed {
			continue
		} else if driver.Time_left < 720 {
			next_load, dist := Find_closest_load_pickup_from_dropoff_of(i, driver.Time_left)
			if next_load == -1 {
				// no possible pickups for driver
				driver.Time_left += Distance_from_depot(driver.Curr_Position)
				break
			}
			driver.Route = append(driver.Route, next_load)
			driver.Loads = append(driver.Loads, *Loads[next_load])
			driver.Time_left += dist + Loads[i].Distance
			driver.Curr_Position = Loads[next_load].Dropoff
		}
	}
	return driver
}
