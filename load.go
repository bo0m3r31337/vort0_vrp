package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Load
X is pickup location
Y is dropoff location
Distance is the distance between them
*/
type Load struct {
	Number                  int
	Pickup                  Point
	Dropoff                 Point
	Distance                float64
	Same_quad               bool
	Distance_from_depot     float64
	Return_distance         float64
	Dropoff_to_pickup_dists []float64
	Completed               bool
}

// func (l Load) String() string {
// 	str := ""
// 	str += "Number = " + strconv.Atoi(l.Number)
// 	str += "\n"
// 	str += l.Pickup.String()
// 	return str
// }

/*
Construct_Load
args:
num // loadNumber
x, y Points
returns load
*/
func Construct_Load(num int, p, d Point) Load {
	load := Load{
		Number:              num,
		Pickup:              p,
		Dropoff:             d,
		Distance:            Distance_between_points(p, d),
		Same_quad:           p.Quadrant == d.Quadrant,
		Distance_from_depot: Distance_from_depot(p),
		Return_distance:     Distance_from_depot(d),
		// Dropoff_to_pickup_dists: make([]float64, ),
		Completed: false,
	}
	return load
}

func (l *Load) Complete() {
	l.Completed = true
}

func (l *Load) UnComplete() {
	l.Completed = false
}

/*
Populate_Load_distances
Calculates the dropoff of a given loads to the pickups of all loads in Loads
*/

func (l *Load) Populate_Load_distances() {
	// fmt.Println("length of loads = ", len(Loads))
	// fmt.Println("l.Number = ", l.Number)
	l.Dropoff_to_pickup_dists = make([]float64, len(Loads))
	for i := 0; i <= len(Loads)-1; i++ {
		if l.Number == i {
			l.Dropoff_to_pickup_dists[i] = 0.0
		}
		l.Dropoff_to_pickup_dists[i] = Distance_between_points(l.Dropoff, Loads[i].Pickup)
	}
	// l.Dropoff_to_pickup_dists = arr
	// fmt.Println("arr in populate loads = ", l.Dropoff_to_pickup_dists)
}

/*
	Return_min_Load_Not_Completed

returns a load number that isn't completed and is the minimum distance from dropoff
returns -1 if all have been completed or there is no min
*/
func (l *Load) Return_min_Load_Not_Completed(driven_minutes float64) (int, float64) {
	min_dist := -1.0
	load_num := -1
	for i := 0; i <= len(l.Dropoff_to_pickup_dists)-1; i++ {
		// fmt.Println("dropoff distance to pickup and minutes driven = ", l.Dropoff_to_pickup_dists[i]+minutes_driven)
		if Loads[i].Completed {
			continue
		}
		if min_dist < 0 && !Loads[i].Completed && driven_minutes+Loads[i].Return_distance+l.Dropoff_to_pickup_dists[i] < 720 {
			min_dist = l.Dropoff_to_pickup_dists[i]
			load_num = i
			continue
		}
		if l.Dropoff_to_pickup_dists[i] < min_dist && !Loads[i].Completed && driven_minutes+Loads[i].Return_distance+l.Dropoff_to_pickup_dists[i] < 720 {
			min_dist = l.Dropoff_to_pickup_dists[i]
			load_num = i
		}
	}
	// fmt.Println("load num = ", load_num)
	return load_num, min_dist
}

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
		number, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Println("error getting load number")
		}
		load := Construct_Load(number-1, pickup, dropoff)
		Loads = append(Loads, &load)
		Load_Nums = append(Load_Nums, number-1)
		// Armada.Total_Minutes += load.Distance
	}
	// fmt.Println("load size = ", Loads)
	// initial fleet size
	// Fleet_Size = math.Ceil(Armada.Total_Minutes / 720.00)
	// Initial_Proxy_Sweep()
	// for len(Loads) != 0 {
	// 	drivers_maxed := Armada.Sweep_Proxy()
	// 	// fmt.Println("drivers maxed out of function = ", drivers_maxed)
	// 	if drivers_maxed != 0 {
	// 		break
	// 	}
	// }
	// Armada.Return_Home()
}

func Distance_between_points(point_1, point_2 Point) float64 {
	difference_in_x := point_1.x - point_2.x
	difference_in_y := point_1.y - point_2.y
	return math.Sqrt((difference_in_x * difference_in_x) + (difference_in_y * difference_in_y))

}

func Distance_from_depot(point Point) float64 {
	difference_in_x := point.x - 0.0
	difference_in_y := point.y - 0.0
	return math.Sqrt((difference_in_x * difference_in_x) + (difference_in_y * difference_in_y))
}

// func (l *Load) rank_dropoff_pickup_distances(loads []*Load) {
// 	for i := 0; i <= len(loads)-1; i++ {
// 		if loads[i].Number == l.Number {
// 			continue
// 		}

// 	}
// }
