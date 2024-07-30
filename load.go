package main

import (
	"math"
)

/*
Load
X is pickup location
Y is dropoff location
Distance is the distance between them
*/
type Load struct {
	Number              string
	Pickup              Point
	Dropoff             Point
	Distance            float64
	Same_quad           bool
	Distance_from_depot float64
}

func (l Load) String() string {
	str := ""
	str += "Number = " + l.Number
	str += "\n"
	str += l.Pickup.String()
	return str
}

/*
Construct_Load
args:
num // loadNumber
x, y Points
returns load
*/
func Construct_Load(num string, p, d Point) Load {
	load := Load{
		Number:              num,
		Pickup:              p,
		Dropoff:             d,
		Distance:            Distance_between_points(p, d),
		Same_quad:           p.Quadrant == d.Quadrant,
		Distance_from_depot: Distance_from_depot(p),
	}
	return load
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

func (l *Load) rank_dropoff_pickup_distances(loads []*Load) {
	for i := 0; i <= len(loads)-1; i++ {
		if loads[i].Number == l.Number {
			continue
		}

	}
}
