package main

import "fmt"

/*Fleet
is a struct of total distances in driving minutes between pickup dropoff loads
and a slice of active drivers
*/
type Fleet struct {
	Drivers       []Driver
	Total_Minutes float64
	Docked        []Driver
}

// func Print_Fleet() string{
// 	for _, d := range Armada.Drivers {
// 		str := "["
// 		for i := 0; i <= len(d)
// 	}
// }

/* Driver
is a struct with a slice of loads and time left and current position
*/
type Driver struct {
	Loads         []Load
	Time_left     float64
	Curr_Position Point
}

func (d *Driver) dock() {
	d.Time_left += Distance_from_depot(d.Curr_Position)
}

func (f Fleet) get_cost() float64 {
	return f.Total_Minutes + (500.00 * float64(len(f.Drivers)))
}

func (f Fleet) Print() {
	for i := 0; i <= len(f.Drivers)-1; i++ {
		fmt.Print("[")
		for j := 0; j <= len(f.Drivers[i].Loads)-1; j++ {
			fmt.Print(f.Drivers[i].Loads[j].Number)
			if j != len(f.Drivers[i].Loads)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]\n")
	}
}
