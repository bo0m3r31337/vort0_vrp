package main

import (
	"sync"
)

/*
	Populate_Matrix

Populates the slice of loads with a slice of distances of their relative dropoffs to pickups
*/
func Populate_Matrix() {
	var waitgroup sync.WaitGroup
	// anon function for thread throttle populates a Load.Dropoff_to_pickup_dist
	populate_distances := func(l *Load) {
		l.Dropoff_to_pickup_dists = make([]float64, len(Loads))
		for i := 0; i <= len(Loads)-1; i++ {
			if l.Number == i {
				l.Dropoff_to_pickup_dists[i] = 0.0
			}
			l.Dropoff_to_pickup_dists[i] = Distance_between_points(l.Dropoff, Loads[i].Pickup)
		}
		defer waitgroup.Done()
	}
	// create routines to populate each load on separate thread
	for i := 0; i <= len(Loads)-1; i++ {
		waitgroup.Add(1)
		go populate_distances(Loads[i])
	}
	//wait before return so all data is populated
	waitgroup.Wait()
}
