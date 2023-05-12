package main

import "math/rand"

// Choosing to attach the aliens to the city instead of attaching the city to the aliens because it's simpler to find alien collisions
type AlienName *string
type CityAliens []AlienName

type AlienMoveMap map[*CityRoads]CityAliens

func (rm RoadMap) moveAliens() {
	// We can't just loop through every city and move the aliens around; we
	// would end up moving some aliens multiple times. Instead, we do one
	// sweep where we determine where to move to, and a second sweep to
	// actually move them.
	alienMoves := make(map[*City]CityAliens)

	for _, city := range rm {
		// Put neighbor cities in a slice so we can easily pick on at random
		var neighborCities [](*City)
		if city.Roads.North != nil {
			neighborCities = append(neighborCities, city.Roads.North)
		}
		if city.Roads.East != nil {
			neighborCities = append(neighborCities, city.Roads.East)
		}
		if city.Roads.South != nil {
			neighborCities = append(neighborCities, city.Roads.South)
		}
		if city.Roads.West != nil {
			neighborCities = append(neighborCities, city.Roads.West)
		}

		// If there are no roads out of this city, all aliens in this
		// city are trapped. Indicate that the list of aliens does not
		// change.
		if len(neighborCities) == 0 {
			alienMoves[city] = city.Aliens
			continue
		}

		// If there are roads, pick a neighbor city at random for each
		// alien and indicate that they will move there.
		for _, alien := range city.Aliens {
			nextCity := neighborCities[rand.Intn(len(neighborCities))]
			alienMoves[nextCity] = append(alienMoves[nextCity], alien)
		}
	}

	// Actually move the aliens to their new destinations
	for _, city := range rm {
		// If city is not in alienMoves, alienMoves[city].Aliens will
		// be empty, which is what we want here anyway.
		city.Aliens = alienMoves[city]
	}
}

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
