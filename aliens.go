package main

import (
	"fmt"
	"math/rand"
)

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
		if city.roads.north != nil {
			neighborCities = append(neighborCities, city.roads.north)
		}
		if city.roads.east != nil {
			neighborCities = append(neighborCities, city.roads.east)
		}
		if city.roads.south != nil {
			neighborCities = append(neighborCities, city.roads.south)
		}
		if city.roads.west != nil {
			neighborCities = append(neighborCities, city.roads.west)
		}

		// If there are no roads out of this city, all aliens in this
		// city are trapped. Indicate that the list of aliens does not
		// change.
		if len(neighborCities) == 0 {
			alienMoves[city] = city.aliens
			continue
		}

		// If there are roads, pick a neighbor city at random for each
		// alien and indicate that they will move there.
		for _, alien := range city.aliens {
			nextCity := neighborCities[rand.Intn(len(neighborCities))]
			alienMoves[nextCity] = append(alienMoves[nextCity], alien)
		}
	}

	// Actually move the aliens to their new destinations
	for _, city := range rm {
		// If city is not in alienMoves, alienMoves[city].aliens will
		// be empty, which is what we want here anyway.
		city.aliens = alienMoves[city]
	}
}

func (rm RoadMap) fightAliens() {
	for cityName, city := range rm {
		// If 2 or more aliens can land on the same city, they destroy the city (and
		// each other).
		if len(city.aliens) > 1 {
			rm.destroyCity(cityName)
		}
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
		cities[index].aliens = append(cities[index].aliens, generateAlienName(x))
	}
}

func generateAlienName(seed int) AlienName {
	// Pick 4 different name segments. With order mattering, that's 3000+ alien names.
	// If we need more, we start appending numbers.

	segments := []string{
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

	name := ""
	for x := 0; x < 4; x++ {
		index := seed % len(segments)
		seed /= len(segments)
		name += segments[index]
		segments = append(segments[0:index], segments[index+1:]...)
	}

	// If we really need so many names, start appending numbers
	// Start with 2 (un-numbered would be 1)
	if seed > 0 {
		name += fmt.Sprintf(" %d", seed+1)
	}

	// Capitalize the first letter
	firstLetter := name[0] + 'A' - 'a'
	name = string(firstLetter) + name[1:]

	return AlienName(&name)
}
