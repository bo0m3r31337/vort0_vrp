package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func Construct_Point(vals string) Point {
	vals_split := strings.Split(vals, ",")
	x_val := strings.Replace(vals_split[0], "(", "", -1)
	y_val := strings.Replace(vals_split[1], ")", "", -1)
	x, err := strconv.ParseFloat(x_val, 64)
	fmt.Println("float64 fx = ", x)
	if err != nil {
		fmt.Println("error creating x val in point constructor error = ", err)
	}
	y, err := strconv.ParseFloat(y_val, 64)
	fmt.Println("float64 fy = ", y)
	if err != nil {
		fmt.Println("error creating y val in point constructor error = ", err)
	}
	point := Point{
		x: x,
		y: y,
	}
	// point.print_point()
	return point
}

func (p Point) print_point() {
	fmt.Println("x =", p.x)
	fmt.Println("y = ", p.y)
}
