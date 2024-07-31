package main

import (
	"slices"
	"sync"
)

/*
	Populate_Matrix

Populates the slice of loads with a slice of distances of their relative dropoffs to pickups
*/
func Populate_Matrix() {
	var waitgroup sync.WaitGroup
	// anon function for thread throttle populates a Load.Dropoff_to_pickup_dist
	populate_distances := func(l *Load) {
		l.Dropoff_to_pickup_dists = make([]float64, len(Loads))
		for i := 0; i <= len(Loads)-1; i++ {
			if l.Number == i {
				l.Dropoff_to_pickup_dists[i] = 0.0
			}
			l.Dropoff_to_pickup_dists[i] = Distance_between_points(l.Dropoff, Loads[i].Pickup)
		}
		defer waitgroup.Done()
	}
	// create routines to populate each load on separate thread
	for i := 0; i <= len(Loads)-1; i++ {
		waitgroup.Add(1)
		go populate_distances(Loads[i])
	}
	//wait before return so all data is populated
	waitgroup.Wait()
}

func Find_Driver_and_Route_Schedule(load_num int) {
	// var waitgroup sync.WaitGroup
	get_load_info := Loads[load_num]
	for i := 0; i <= len(Armada.Drivers)-1; i++ {
		if slices.Contains(Armada.Drivers[i].Route, load_num) {
			continue
		}
		if Armada.Drivers[i].Time_left+get_load_info.Distance > 720 || Armada.Drivers[i].Time_left+get_load_info.Return_distance > 720 {
			continue
		} else {
			// waitgroup.Add(1)
			new_load_schedule, _ := Armada.Drivers[i].Search_Driver(get_load_info)
			if new_load_schedule != nil {
				Armada.Drivers[i].Loads = new_load_schedule
			} else {
				continue
			}
		}
	}
}

func (driver Driver) Search_Driver(load *Load) ([]Load, float64) {
	var insert_index int
	var min_delta float64

	// check insert at front delta
	start_delta := load.Distance_from_depot + load.Distance + load.Dropoff_to_pickup_dists[driver.Loads[0].Number] + driver.Curr_Position.Distance_from_home
	min_delta = start_delta
	insert_index = 0
	for i := 1; i <= len(driver.Loads)-2; i++ {
		delta := driver.Loads[i].Dropoff_to_pickup_dists[load.Number] + load.Distance + load.Dropoff_to_pickup_dists[driver.Loads[i+1].Number] + driver.Curr_Position.Distance_from_home
		if delta < min_delta {
			min_delta = delta
			insert_index = i
		}
	}
	end_delta := load.Return_distance + load.Distance + driver.Loads[len(driver.Loads)-1].Dropoff_to_pickup_dists[load.Number]
	if end_delta < min_delta {
		min_delta = end_delta
		insert_index = len(driver.Loads) - 1
	}
	if min_delta+driver.Time_left > 720 {
		return nil, 0.0
	} else {
		return append(driver.Loads[:insert_index+1], append([]Load{*load}, driver.Loads[insert_index+1:]...)...), min_delta + driver.Time_left
	}
}
