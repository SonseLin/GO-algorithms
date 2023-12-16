package main

import (
	"fmt"
	"math"
)

type currentNavigationRoot struct {
	location            string
	directions          map[string]*currentNavigationRoot
	distance            int
	latitude, longitude float64
}

func distance(from, to currentNavigationRoot) float64 {
	lat1Rad := degToRad(from.latitude)
	lon1Rad := degToRad(from.longitude)
	lat2Rad := degToRad(to.latitude)
	lon2Rad := degToRad(to.longitude)

	dlon := lon2Rad - lon1Rad

	distance := math.Acos(math.Sin(lat1Rad)*math.Sin(lat2Rad)+math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(dlon)) * 6371

	return distance
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func getGraph() *currentNavigationRoot {
	var navGraph *currentNavigationRoot
	navGraph = &currentNavigationRoot{
		location:   "piter",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   59.93,
		longitude:  30.36,
	}
	moscow := &currentNavigationRoot{
		location:   "moscow",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   55.75,
		longitude:  37.62,
	}
	kazan := &currentNavigationRoot{
		location:   "kazan",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   55.79,
		longitude:  49.12,
	}
	navGraph.directions["moscow"] = moscow
	moscow.directions["piter"] = navGraph
	moscow.directions["kazan"] = kazan
	kazan.directions["moscow"] = moscow
	return navGraph
}

func main() {
	// setup
	totalTravelDistance := 0.0
	var destinationCity string
	currentLocation := getGraph()
	ok := true
	fmt.Println(currentLocation.directions["kazan"])
	// cycle
	for ok {
		fmt.Printf("from %s you can go on to:\n", currentLocation.location)
		for _, city := range currentLocation.directions {
			fmt.Println(city.location)
		}
		_, err := fmt.Scanf("%s", &destinationCity)
		if err == nil {
			destCity := currentLocation.directions[destinationCity]
			if destinationCity == currentLocation.location {
				ok = false
			}
			if destCity != nil {
				fmt.Println("travelled to", destinationCity)
				totalTravelDistance += distance(*currentLocation, *destCity)
				currentLocation = destCity
			} else {
				fmt.Println("no city to travel", destinationCity)
			}
		}
	}
	fmt.Println("total travelled distance is", totalTravelDistance)
}
