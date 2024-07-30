package main

import (
	"os"
)

var Loads []*Load
var Armada Fleet
var Fleet_Size float64

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
	Armada.Print()
	// end_time := time.Now()
	// exec_time := end_time.Sub(start_time)
	// fmt.Println("exec time = ", exec_time)
}
