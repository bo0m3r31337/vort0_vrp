package main

func Reduce_Drivers_V2() bool {
	var reduction_index []int
	for i := 0; i <= len(Armada.Drivers)-1; i++ {
		if len(Armada.Drivers[i].Loads) == 1 {
			if Insert_Route_To_New_Driver_V2(Armada.Drivers[i].Loads[0], i+1) {
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

func Reduce_Drivers_V3() bool {
	var reduction_index []int
	for i := 0; i <= len(Armada.Drivers)-1; i++ {
		if len(Armada.Drivers[i].Loads) == 1 {
			if Insert_Route_To_New_Driver_V2(Armada.Drivers[i].Loads[0], i+1) {
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

func Insert_Route_To_New_Driver(load_to_insert Load, driver_start_index int) bool {
	// first_load_to_trasfer := Armada.Drivers[0].Loads[0]
	for i := driver_start_index; i <= len(Armada.Drivers)-1; i++ {
		for l := 0; l <= len(Armada.Drivers[i].Loads)-1; l++ {
			if l == 0 {
				insert_here, plus_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
				if insert_here {
					Armada.Drivers[i].Loads = append([]Load{load_to_insert}, Armada.Drivers[i].Loads[0:]...)
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
				insert_here, plus_time := Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[l+1], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
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

func Insert_Route_To_New_Driver_V2(load_to_insert Load, driver_start_index int) bool {
	// first_load_to_trasfer := Armada.Drivers[0].Loads[0]
	for i := driver_start_index; i <= len(Armada.Drivers)-1; i++ {
		length_of_loads := len(Armada.Drivers[i].Loads)
		for l := 0; l <= len(Armada.Drivers[i].Loads)-1; l++ {
			if length_of_loads == 1 {
				front_insert, front_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
				back_insert, back_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
				if front_insert || back_insert {
					if front_time < back_time {
						Armada.Drivers[i].Loads = append([]Load{load_to_insert}, Armada.Drivers[i].Loads[0:]...)
						Armada.Drivers[i].Time_left += front_time
						return true
					} else {
						Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads, load_to_insert)
						Armada.Drivers[i].Time_left += back_time
						return true
					}
				}
			} else if length_of_loads == 2 {
				front_insert, front_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
				back_insert, back_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
				mid_insert, mid_time := Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[length_of_loads-1], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
				if (front_time < back_time && front_time < mid_time) && front_insert {
					insert_here, plus_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
					if insert_here {
						Armada.Drivers[i].Loads = append([]Load{load_to_insert}, Armada.Drivers[i].Loads[0:]...)
						Armada.Drivers[i].Time_left += plus_time
						return true
					}
				} else if (back_time < front_time && back_time < mid_time) && back_insert {
					insert_here, plus_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
					if insert_here {
						Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads, load_to_insert)
						Armada.Drivers[i].Time_left += plus_time
						return true
					}
				} else if (mid_time < front_time && mid_time < back_time) && mid_insert {
					insert_here, plus_time := Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[l+1], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
					if insert_here {
						Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads[:l+1], append([]Load{load_to_insert}, Armada.Drivers[i].Loads[l+1:]...)...)
						Armada.Drivers[i].Time_left += plus_time
						return true
					}
				} else {
					continue
				}
			} else if length_of_loads > 2 {
				var mid_insert bool
				var mid_time float64
				front_insert, front_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
				back_insert, back_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
				if length_of_loads-1 != l {
					mid_insert, mid_time = Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[l+1], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))

				} else {
					mid_insert = false
					mid_time = 0.0
				}
				if (front_time < back_time && front_time < mid_time) && front_insert {
					insert_here, plus_time := Insert_at_beginning(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
					if insert_here {
						Armada.Drivers[i].Loads = append([]Load{load_to_insert}, Armada.Drivers[i].Loads[0:]...)
						Armada.Drivers[i].Time_left += plus_time
						return true
					}
				} else if (back_time < front_time && back_time < mid_time) && back_insert {
					insert_here, plus_time := Insert_Load_at_end(Armada.Drivers[i].Loads[l], load_to_insert, Armada.Drivers[i].Time_left)
					if insert_here {
						Armada.Drivers[i].Loads = append(Armada.Drivers[i].Loads, load_to_insert)
						Armada.Drivers[i].Time_left += plus_time
						return true
					}
				} else if (mid_time < front_time && mid_time < back_time) && mid_insert {
					insert_here, plus_time := Insert_Load_mid_route(Armada.Drivers[i].Loads[l], Armada.Drivers[i].Loads[l+1], load_to_insert, Armada.Drivers[i].Time_left+Distance_from_depot(Armada.Drivers[i].Curr_Position))
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
