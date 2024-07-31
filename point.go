package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x                  float64
	y                  float64
	Quadrant           int
	Distance_from_home float64
}

func Construct_Point(vals string) Point {
	vals_split := strings.Split(vals, ",")
	x_val := strings.Replace(vals_split[0], "(", "", -1)
	y_val := strings.Replace(vals_split[1], ")", "", -1)
	x, err := strconv.ParseFloat(x_val, 64)
	// fmt.Println("float64 fx = ", x)
	if err != nil {
		fmt.Println("error creating x val in point constructor error = ", err)
	}
	y, err := strconv.ParseFloat(y_val, 64)
	// fmt.Println("float64 fy = ", y)
	if err != nil {
		fmt.Println("error creating y val in point constructor error = ", err)
	}
	point := Point{
		x: x,
		y: y,
	}
	if x > 0 && y > 0 {
		point.Quadrant = 1
	} else if x < 0 && y > 0 {
		point.Quadrant = 2
	} else if x < 0 && y < 0 {
		point.Quadrant = 3
	} else if x == 0 && y == 0 {
		point.Quadrant = 0
	} else {
		point.Quadrant = 4
	}
	point.Distance_from_home = Distance_from_depot(point)
	return point
}

func (p Point) String() string {
	return fmt.Sprintf("x = %f\ny= %f\n", p.x, p.y)
}

func Make_Home_Point() Point {
	home := Point{
		x:                  0.0,
		y:                  0.0,
		Quadrant:           0,
		Distance_from_home: 0,
	}
	return home
}
