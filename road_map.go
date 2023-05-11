package main

import (
	"fmt"
	"regexp"
	"strings"
)

// TODO - lowercase everything, we're not leaving the main module

// We can use *CityRoads as an identifier for the city. We can compare pointers sometimes.

type CityName string
type CityRoads struct {
	North *CityRoads
	East  *CityRoads
	South *CityRoads
	West  *CityRoads
}

// TODO can we get rid of the *, and refer to &roadMap[name] and get the right pointer?
type RoadMap map[CityName]*CityRoads

// Are roads automatically two-way? Suppose the map says Foo has north=Bar, but
// Bar doesn't have south=Foo. Should I assume:
//
// * Bar south=Foo implicitly
// * A one-way road north from Foo to Bar
// * An invalid file
//
// I don't want to assume the implicit two-way, because then I'd need to assume
// that the map is actually laid out on a nice grid, roads don't curve, etc. I
// don't want to assume one-way roads, that doesn't make as much sense. I will
// assume the file is invalid if there isn't a reference back, even if it's not
// coming from the opposite direction. I.e. a road north from Foo can come in
// from the east of Bar.

// Each city name should be alphanumeric or underscore

func parseMap(input string) (roadMap RoadMap, err error) {
	roadMap = make(RoadMap)

	// Loop over input once to just look at the cities being defined, to create
	// the empty city stucts. Having this as a separate step makes it easier to
	// check later for duplicate cities or later check for references to
	// nonexistent cities.
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		segments := strings.Split(line, " ")

		// Allow empty lines
		if len(segments) < 1 {
			continue
		}

		subjectName := CityName(segments[0])

		re := regexp.MustCompile(`^\w+$`)
		if !re.Match([]byte(subjectName)) {
			err = fmt.Errorf("Invalid line in initial map (invalid city name): " + line)
			return
		}

		if _, ok := roadMap[subjectName]; ok {
			err = fmt.Errorf("Invalid line in initial map (repeated city): " + line)
			return
		}

		// Initialize a new city at this name in the roadMap
		newCity := CityRoads{}
		roadMap[subjectName] = &newCity
	}

	// Loop over input a second time to make the connections between cities
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		segments := strings.Split(line, " ")

		// Allow empty lines
		if len(segments) < 1 {
			continue
		}

		// The subject is the city we're making roads to and from
		subjectName := CityName(segments[0])
		subject := roadMap[subjectName]

		for _, road := range segments[1:] {
			// Each road is in the form of "direction=destinationName"
			roadParts := strings.Split(road, "=")

			if len(roadParts) != 2 {
				err = fmt.Errorf("Invalid line in initial map (expecting one = in road description): " + line)
				return
			}
			direction := roadParts[0]

			destinationName := CityName(roadParts[1])
			if _, ok := roadMap[destinationName]; !ok {
				err = fmt.Errorf("Invalid line in initial map (road to undefined city): " + line)
				return
			}
			destination := roadMap[destinationName]

			if subject == destination {
				err = fmt.Errorf("Invalid line in initial map (road from city to itself): " + line)
				return
			}

			if subject.North == destination ||
				subject.East == destination ||
				subject.South == destination ||
				subject.West == destination {
				err = fmt.Errorf("Invalid line in initial map (two roads to the same town): " + line)
				return
			}

			switch direction {
			case "north":
				if subject.North != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.North = destination
			case "east":
				if subject.East != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.East = destination
			case "south":
				if subject.South != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.South = destination
			case "west":
				if subject.West != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.West = destination
			default:
				err = fmt.Errorf("Invalid line in initial map (invalid direction): " + line)
				return
			}
		}
	}

	// Loop over the results to make sure no one-way roads

	// This function checks whether a given city is connected to another city
	isHalfConnectedCity := func(subject *CityRoads, destination *CityRoads) bool {
		// If the destination is nil in the first place, it's not a
		// city, therefore not a half-connected city.
		if destination == nil {
			return false
		}

		// If none of the directions from destination go back to the
		// subject, destination is half-connected
		return !(destination.North == subject ||
			destination.East == subject ||
			destination.South == subject ||
			destination.West == subject)
	}
	for _, city := range roadMap {
		// Confirm that every connected city connects back
		if isHalfConnectedCity(city, city.North) ||
			isHalfConnectedCity(city, city.East) ||
			isHalfConnectedCity(city, city.South) ||
			isHalfConnectedCity(city, city.West) {
			err = fmt.Errorf("Invalid initial map (some roads are not connected on both ends)")
			return
		}
	}
	return
}

func printMap(rm RoadMap) {
}

func destroyCity() {
	// ... follow links, set nils, delete entry from Map. That will kill the aliens inside.
	// Bar has been destroyed by Goomkormonzor and Thublarkorxan!
}
