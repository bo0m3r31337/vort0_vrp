package main

import (
	"fmt"
)

/*
	Init_Matrix

initializes distance matrix
*/
func Init_Matrix() {
	Distance_Matrix = make([][]float64, len(Loads))
	for i := 0; i <= len(Loads)-1; i++ {
		Distance_Matrix[i] = make([]float64, len(Loads))
	}
}

/*
	Populate_Matrix

Populates the slice of loads with a slice of distances of their relative dropoffs to pickups
*/
func Populate_Matrix() {
	// var waitgroup sync.WaitGroup

	for i := 0; i <= len(Loads)-1; i++ {
		Loads[i].Populate_Load_distances()
	}
	// var waitgroup sync.WaitGroup
	// populate_cell := func(load_num_init, load_num_dest int, init, dest Load) {
	// 	Distance_Matrix[load_num_init][load_num_dest] = Distance_between_points(init.Dropoff, dest.Pickup)
	// 	Distance_Matrix[load_num_dest][load_num_init] = Distance_between_points(dest.Dropoff, init.Pickup)
	// 	defer waitgroup.Done()
	// }
	// for i := 0; i <= len(Loads)-1; i++ {
	// 	for j := i; j <= len(Loads)-1; j++ {
	// 		if i == j {
	// 			continue
	// 		}
	// 		waitgroup.Add(1)
	// 		go populate_cell(Loads[i].Number, Loads[j].Number, *Loads[i], *Loads[j])
	// 	}
	// }
	// waitgroup.Wait()
	// fmt.Println(Distance_Matrix)
}
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
		// if driver.Time_left+Distance_from_depot(next_load.Dropoff) > 720 {
		// 	// fmt.Println("driver is maxed in init")
		// }
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

func Find_closest_load_pickup_from_dropoff_of(load_num int, distance_already_traveled float64) (int, float64) {
	fmt.Println("load num = ", load_num)
	fmt.Println("distance already travelled = ", distance_already_traveled)
	closest_next_pickup := -1
	var min_dist float64
	init := true
	for i := 0; i <= len(Distance_Matrix[load_num])-1; i++ {
		if Distance_Matrix[load_num][i] == 0 {
			continue
		}
		if init {
			min_dist = Distance_Matrix[load_num][i]
			init = false
			closest_next_pickup = i
			continue
		}
		if Distance_Matrix[load_num][i] < min_dist && Distance_Matrix[load_num][i] > 0 && distance_already_traveled+Distance_from_depot(*&Loads[load_num].Dropoff) < 720 {
			min_dist = Distance_Matrix[load_num][i]
			closest_next_pickup = i
		}
	}
	if closest_next_pickup == -1 {
		return -1, 0.0
	}
	Distance_Matrix[load_num][closest_next_pickup] = -1.0
	return closest_next_pickup, min_dist
}
