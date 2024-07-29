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
	Number   int64
	Pickup   Point
	Dropoff  Point
	Distance float64
}

/*
Construct_Load
args:
num // loadNumber
x, y Points
returns load
*/
func Construct_Load(num int64, p, d Point) Load {
	load := Load{
		Number:   num,
		Pickup:   p,
		Dropoff:  d,
		Distance: Distance_between_points(p, d),
	}
	return load
}

func Distance_between_points(point_1, point_2 Point) float64 {
	difference_in_x := point_1.x - point_2.x
	difference_in_y := point_1.y - point_2.y
	return math.Sqrt((difference_in_x * difference_in_x) + (difference_in_y * difference_in_y))

}
