package main

import (
	"fmt"
	"math"
)

type currentNavigationRoot struct {
	location            string
	directions          map[string]*currentNavigationRoot
	distance            float64
	latitude, longitude float64
	shortestPath        string
	isVisited           bool
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
	perm := &currentNavigationRoot{
		location:   "perm",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   58.01,
		longitude:  56.23,
	}
	tver := &currentNavigationRoot{
		location:   "tver",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   56.86,
		longitude:  35.92,
	}
	ebaniiGorod := &currentNavigationRoot{
		location:   "ebanii",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   100,
		longitude:  100,
		isVisited:  false,
	}
	ekata := &currentNavigationRoot{
		location:   "ekata",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   56.84,
		longitude:  60.65,
	}
	ufa := &currentNavigationRoot{
		location:   "ufa",
		directions: make(map[string]*currentNavigationRoot),
		distance:   0,
		latitude:   54.74,
		longitude:  55.96,
	}
	navGraph.directions["moscow"] = moscow

	moscow.directions["piter"] = navGraph
	moscow.directions["kazan"] = kazan

	kazan.directions["moscow"] = moscow
	kazan.directions["ebanii"] = ebaniiGorod
	kazan.directions["perm"] = perm
	kazan.directions["tver"] = tver

	ebaniiGorod.directions["tver"] = tver

	tver.directions["kazan"] = kazan
	tver.directions["perm"] = perm
	tver.directions["ekata"] = ekata
	tver.directions["ufa"] = ufa

	perm.directions["ekata"] = ekata
	perm.directions["kazan"] = kazan
	perm.directions["tver"] = tver

	ekata.directions["perm"] = perm
	ekata.directions["tver"] = tver

	ufa.directions["tver"] = tver
	return navGraph
}

func interactiveNavigation(totalTravelDistance *float64, currentLocation *currentNavigationRoot) {
	ok := true
	var destinationCity string
	for ok {
		fmt.Printf("from %s you can go on to:\n", currentLocation.location)
		count := 0
		for _, city := range currentLocation.directions {
			count++
			fmt.Printf("%d) %s\n", count, city.location)
		}
		_, err := fmt.Scanf("%s", &destinationCity)
		if err == nil {
			destCity := currentLocation.directions[destinationCity]
			if destinationCity == currentLocation.location {
				ok = false
			}
			if destCity != nil {
				*totalTravelDistance += distance(*currentLocation, *destCity)
				fmt.Println("travelled to", destinationCity, "\nTravelled distance is", *totalTravelDistance)
				currentLocation = destCity
			} else {
				fmt.Println("no city to travel", destinationCity)
			}
		}
	}
}

func interactiveMode(currentLocation *currentNavigationRoot) {
	totalTravelDistance := 0.0
	interactiveNavigation(&totalTravelDistance, currentLocation)
	fmt.Println("total travelled distance is", totalTravelDistance)
}

func DFS(cities *currentNavigationRoot, destPoint **currentNavigationRoot, curr, dest string, totalDistance *float64) bool {
	if curr == dest {
		if *totalDistance == 0 {
			*totalDistance = cities.distance
		} else if *totalDistance != 0 && *totalDistance > cities.distance {
			*totalDistance = cities.distance
		}
		*destPoint = cities
		return true
	}
	if cities.isVisited {
		return false
	}
	cities.isVisited = true
	for _, city := range cities.directions {
		potentialDistance := cities.distance + distance(*cities, *city)
		// если новое расстояние меньше старого или не определено
		if city.distance == 0 || city.distance > potentialDistance {
			city.shortestPath = cities.location
			city.distance = potentialDistance
		}
	}
	for _, city := range cities.directions {
		if !city.isVisited {
			endpoint := DFS(city, destPoint, city.location, dest, totalDistance)
			if endpoint {
				return true
			}
		}
	}
	return false
}

func searchShortPathMode(currentLocation *currentNavigationRoot) {
	var destination string
	var shortestPath []string
	_, err := fmt.Scanf("%s", &destination)
	if err == nil {
		totalDistance := 0.0
		var endPoint *currentNavigationRoot
		ok := DFS(currentLocation, &endPoint, currentLocation.location, destination, &totalDistance)
		if ok == true {
			fmt.Println("shortest path to", endPoint.location, "is", totalDistance, "includes cities:")
			for endPoint != currentLocation {
				shortestPath = append(shortestPath, endPoint.location)
				endPoint = endPoint.directions[endPoint.shortestPath]
			}
			shortestPath = append(shortestPath, currentLocation.location)
		} else {
			fmt.Printf("%s city not found\n", destination)
		}
	} else {
		panic(err)
	}
	for _, i := range shortestPath {
		fmt.Println(i)
	}
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("follow code rules, input int")
		}
	}()
	currentLocation := getGraph()
	var mode int
	_, err := fmt.Scanf("%d", &mode)
	if err == nil {
		switch mode {
		case 1:
			interactiveMode(currentLocation)
		case 2:
			searchShortPathMode(currentLocation)
		}
	} else {
		panic(err)
	}
}
