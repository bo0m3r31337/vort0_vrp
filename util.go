package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

/*
Get_loads
reads in the file and creates a slice of loads to carry within the 12 hour period
*/
func Get_loads(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening the file = ", err)
	}
	scanner := bufio.NewScanner(file)
	//burn first line
	scanner.Scan()
	// fmt.Println("before load ingestion = , ", scanner.Text())
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		pickup := Construct_Point(arr[1])
		dropoff := Construct_Point(arr[2])
		number := arr[0]
		load := Construct_Load(number, pickup, dropoff)
		Loads = append(Loads, &load)
		Armada.Total_Minutes += load.Distance
		// fmt.Println("load number = ", load.Number)
		// fmt.Println("distance from depot = ", load.Distance_from_depot)
		// fmt.Println("distance from depot = ", Distance_from_depot(load.Dropoff))
		// fmt.Println("load distance = ", load.Distance)
		// Armada.Total_Minutes += Distance_from_depot(load.Dropoff)
	}
	// initial fleet size
	Fleet_Size = math.Ceil(Armada.Total_Minutes / 720.00)
	// fmt.Println("total minutes = ", Armada.Total_Minutes)
	// fmt.Println("initial fleet size = ", Fleet_Size)
	// fmt.Println("amount of loads = ", len(Loads))
	// fmt.Println("max loads per driver = ", math.Ceil(float64(len(Loads))/Fleet_Size))
	Initial_Proxy_Sweep()
	for len(Loads) != 0 {
		drivers_maxed := Armada.Sweep_Proxy()
		// fmt.Println("drivers maxed out of function = ", drivers_maxed)
		if drivers_maxed != 0 {
			break
		}
	}
	Armada.Return_Home()
	// fmt.Println("cost = ", Armada.get_cost())
	// for j := 0; j <= len(Armada.Drivers)-1; j++ {
	// 	// fmt.Println("driver loads length = ", len(Armada.Drivers[j].Loads))
	// 	fmt.Print("drivers time driven = ", Armada.Drivers[j].Time_left)
	// 	// for i := 0; i <= len(Armada.Drivers[j].Loads)-1; i++ {
	// 	// 	fmt.Print("d.loads = ", Armada.Drivers[j].Loads[i].Number+"\n")
	// 	// 	fmt.Print("\n")
	// 	// }
	// 	// fmt.Println("load number ", d.Loads[i].Number)
	// 	// fmt.Println("distance from depot = ", d.Loads[0].Distance_from_depot)
	// }
}

// func Deploy_Driver() {

// }

func Initial_Proxy_Sweep() {
	for i := 0; i <= int(Fleet_Size)+3; i++ {
		driver := Driver{
			Loads:         make([]Load, 0),
			Time_left:     0.0,
			Curr_Position: Home_Point,
		}
		next_load, load_index := return_closest_load(driver.Curr_Position)
		driver.Loads = append(driver.Loads, next_load)
		driver.Curr_Position = driver.Loads[0].Dropoff
		driver.Time_left += driver.Loads[0].Distance
		driver.Time_left += driver.Loads[0].Distance_from_depot
		if driver.Time_left+Distance_from_depot(next_load.Dropoff) > 720 {
			// fmt.Println("driver is maxed in init")
		}
		// fmt.Println("initial driver time = ", driver.Time_left)
		Remove_Load(load_index)
		Armada.Drivers = append(Armada.Drivers, driver)
	}
}

func Remove_Load(index int) {
	Loads = append(Loads[:index], Loads[index+1:]...)
}

func (f Fleet) Sweep_Proxy() int {
	// fmt.Println("length of loads = ", len(Loads))
	drivers_maxed := 0
	for i := 0; i <= len(f.Drivers)-1; i++ {
		if len(Loads) == 0 {
			break
		}
		next_load, load_index := return_closest_load(f.Drivers[i].Curr_Position)
		if f.Drivers[i].Time_left+next_load.Distance+Distance_from_depot(next_load.Dropoff)+Distance_between_points(f.Drivers[i].Curr_Position, next_load.Pickup) > 720.0 {
			// fmt.Println("drivers maxed =  ", f.Drivers[i].Time_left+next_load.Distance+Distance_from_depot(next_load.Dropoff))
			drivers_maxed++
			continue
		}
		f.Drivers[i].Loads = append(f.Drivers[i].Loads, next_load)
		f.Drivers[i].Time_left += Distance_between_points(f.Drivers[i].Curr_Position, next_load.Pickup)
		f.Drivers[i].Curr_Position = next_load.Dropoff
		f.Drivers[i].Time_left += next_load.Distance
		// fmt.Println("time driving = ", f.Drivers[i].Time_left)
		Remove_Load(load_index)
	}
	// fmt.Println("drivers maxed in function = ", drivers_maxed)
	return drivers_maxed
}

func (f Fleet) Return_Home() {
	Armada.Total_Minutes = 0
	for i := 0; i <= len(f.Drivers)-1; i++ {
		f.Drivers[i].Time_left += Distance_from_depot(f.Drivers[i].Loads[len(f.Drivers[i].Loads)-1].Dropoff)
		// fmt.Println("time driving total = ", f.Drivers[i].Time_left)
		Armada.Total_Minutes += f.Drivers[i].Time_left
	}
	// fmt.Println("Armada total minutes = ", Armada.Total_Minutes)
}

func return_closest_load(current_position Point) (Load, int) {
	var min float64
	load_index := 0
	init := true
	for i := 0; i <= len(Loads)-1; i++ {
		dist := Distance_between_points(current_position, Loads[i].Pickup)
		if init {
			min = dist
			init = false
			load_index = i
			continue
		}
		if dist < min {
			min = dist
			load_index = i
		}
	}
	load_to_return := *Loads[load_index]
	// fmt.Println("load number = ", load_to_return.Number)
	// Loads = append(Loads[:load_index], Loads[load_index+1:]...)
	// fmt.Println("length of loads = ", len(Loads))
	return load_to_return, load_index
}
