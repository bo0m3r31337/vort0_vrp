package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Loads []Load

func main() {
	fmt.Println(os.Args[1])
	get_loads(os.Args[1])
}

func get_loads(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening the file = ", err)
	}
	scanner := bufio.NewScanner(file)
	//burn first line
	scanner.Scan()
	fmt.Println("before load ingestion = , ", scanner.Text())
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		pickup := Construct_Point(arr[1])
		dropoff := Construct_Point(arr[2])
		number, err := strconv.ParseInt(arr[0], 10, 64)
		if err != nil {
			fmt.Println("error parsing loadNumber err = ", err)
		}
		load := Construct_Load(number, pickup, dropoff)
		Loads = append(Loads, load)
	}
	for i := 0; i <= len(Loads)-1; i++ {
		fmt.Println("Load number = ", Loads[i].Number)
		Loads[i].Pickup.print_point()
		Loads[i].Dropoff.print_point()
		fmt.Println("distance between points = ", Loads[i].Distance)

	}
	// fmt.Println(Loads)
}
