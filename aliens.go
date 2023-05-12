package main

import "math/rand"

// Choosing to attach the aliens to the city instead of attaching the city to the aliens because it's simpler to find alien collisions
type AlienName *string
type CityAliens []AlienName

type AlienMoveMap map[*CityRoads]CityAliens

/*
func moveAliens(thisMap AlienMap) (nextMap AlienMap) {
	// In order to keep track of which aliens have already moved, we are making a new AlienMap each time
	// TODO - maybe we map alien to city ... but then we need
	// to have both at least temporarily, for simplicity
	return AlienMap{}
}
*/

func (rm RoadMap) setInitialAliens(numAliens int) {
	// Since we're picking at random we'll want cities integer-indexed
	var cities [](*City)
	for _, city := range rm {
		cities = append(cities, city)
	}

	for x := 0; x < numAliens; x++ {
		index := rand.Intn(len(rm))
		cities[index].Aliens = append(cities[index].Aliens, generateAlienName())
	}
}

// TODO Test - alien names are unique, etc
func generateAlienName() AlienName {
	// Pick 4 different name segments, order matters, that's 3000 alien names.
	// If we need more, we can append numbers

	// Capitalize the first letter
	_ = []string{
		"goom",
		"kor",
		"mon",
		"zor",
		"xan",
		"blax",
		"thu",
		"blar",
		"yaf",
	}

	a := ""
	return AlienName(&a)
}
