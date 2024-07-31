package main

func Reduce_Drivers() bool {
	// reduction_index := -1
	for i := 0; i <= len(Armada.Drivers)-1; i++ {
		if len(Armada.Drivers[i].Loads) == 1 {
			if Insert_Route_To_New_Driver(Armada.Drivers[i].Loads[0], i+1) {
				Armada.Drivers = append(Armada.Drivers[:i], Armada.Drivers[i+1:]...)
				i--
			}
		} else {
			continue
		}
	}
	return true
}

func Insert_Route_To_New_Driver(load_to_insert Load, driver_start_index int) bool {
	// first_load_to_trasfer := Armada.Drivers[0].Loads[0]
	for i := driver_start_index; i <= len(Armada.Drivers)-1; i++ {
		for l := 0; l <= len(Armada.Drivers[i].Loads)-1; l++ {
			if l == 0 {
				insert_here, plus_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
				if insert_here {
					Armada.Drivers[i].Loads = append([]Load{load_to_insert}, Armada.Drivers[i].Loads...)
					Armada.Drivers[i].Time_left += plus_time
					return true
				}
			} else if l == len(Armada.Drivers[i].Loads)-1 {
				insert_here, plus_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
				if insert_here {
					Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads, load_to_insert)
					Armada.Drivers[i].Time_left += plus_time
					return true
				}
			} else if l != 0 {
				insert_here, plus_time := Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[l+1], load_to_insert, Armada.Drivers[i].Time_left)
				if insert_here {
					Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads[:l+1], append([]Load{load_to_insert}, Armada.Drivers[i].Loads[l+1:]...)...)
					Armada.Drivers[i].Time_left += plus_time
					return true
				}
			} else {
				continue
			}
		}
	}
	return false
}

func Insert_Load_mid_route(stop_before, stop_after, load_to_insert Load, driving_time float64) (bool, float64) {
	time := stop_before.Dropoff_to_pickup_dists[load_to_insert.Number] + load_to_insert.Distance + load_to_insert.Dropoff_to_pickup_dists[stop_after.Number] + driving_time
	return time < 720, time
}

func Insert_Load_at_end(stop_before, load_to_insert Load, driving_time float64) (bool, float64) {
	time := stop_before.Dropoff_to_pickup_dists[load_to_insert.Number] + load_to_insert.Distance + load_to_insert.Return_distance + driving_time
	return time < 720, time
}

func Insert_at_beginning(stop_after, load_to_insert Load, driving_time float64) (bool, float64) {
	time := load_to_insert.Distance_from_depot + load_to_insert.Dropoff_to_pickup_dists[stop_after.Number] + load_to_insert.Distance + driving_time
	return time < 720.0, time
}
