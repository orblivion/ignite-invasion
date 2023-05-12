package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// We can use *CityRoads as an identifier for the city. We can compare pointers sometimes.

type CityName string
type CityRoads struct {
	north *City
	east  *City
	south *City
	west  *City
}

type City struct {
	aliens CityAliens
	roads  CityRoads
}

type RoadMap map[CityName]*City

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
	// check for duplicate cities or later check for references to nonexistent
	// cities.
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		segments := strings.Split(line, " ")

		// Allow empty lines
		if len(segments) == 1 && len(segments[0]) == 0 {
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
		newCity := City{roads: CityRoads{}}
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

			if subject.roads.north == destination ||
				subject.roads.east == destination ||
				subject.roads.south == destination ||
				subject.roads.west == destination {
				err = fmt.Errorf("Invalid line in initial map (two roads to the same town): " + line)
				return
			}

			switch direction {
			case "north":
				if subject.roads.north != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.roads.north = destination
			case "east":
				if subject.roads.east != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.roads.east = destination
			case "south":
				if subject.roads.south != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.roads.south = destination
			case "west":
				if subject.roads.west != nil {
					err = fmt.Errorf("Invalid line in initial map (road direction repeated): " + line)
					return
				}
				subject.roads.west = destination
			default:
				err = fmt.Errorf("Invalid line in initial map (invalid direction): " + line)
				return
			}
		}
	}

	// Loop over the results to make sure no one-way roads

	// This function checks whether a given city is connected to another city
	isHalfConnectedCity := func(subject *City, destination *City) bool {
		// If the destination is nil in the first place, it's not a
		// city, therefore not a half-connected city.
		if destination == nil {
			return false
		}

		// If none of the directions from destination go back to the
		// subject, destination is half-connected
		return !(destination.roads.north == subject ||
			destination.roads.east == subject ||
			destination.roads.south == subject ||
			destination.roads.west == subject)
	}
	for _, city := range roadMap {
		// Confirm that every connected city connects back
		if isHalfConnectedCity(city, city.roads.north) ||
			isHalfConnectedCity(city, city.roads.east) ||
			isHalfConnectedCity(city, city.roads.south) ||
			isHalfConnectedCity(city, city.roads.west) {
			err = fmt.Errorf("Invalid initial map (some roads are not connected on both ends)")
			return
		}
	}
	if len(roadMap) == 0 {
		err = fmt.Errorf("Invalid initial map (empty)")
		return
	}
	return
}

func (rm RoadMap) outputMap() (output string) {
	// Sort the city names so we get predictable output
	var orderedCityNames sort.StringSlice
	for cityName, _ := range rm {
		// Cast it to string so that we can use the string sorting function
		orderedCityNames = append(orderedCityNames, string(cityName))
	}
	sort.Sort(orderedCityNames)

	// A rare case that we want a city name given a city
	nameByCity := make(map[*City]string)
	for cityName, city := range rm {
		// We'll use it for output so let's just cast it to string now
		nameByCity[city] = string(cityName)
	}

	for i, cn := range orderedCityNames {
		output += cn

		// Cast it back to CityName
		cityName := CityName(cn)
		roads := rm[cityName].roads

		if roads.north != nil {
			output += " north=" + nameByCity[roads.north]
		}
		if roads.east != nil {
			output += " east=" + nameByCity[roads.east]
		}
		if roads.south != nil {
			output += " south=" + nameByCity[roads.south]
		}
		if roads.west != nil {
			output += " west=" + nameByCity[roads.west]
		}

		// Newline if it's not the last line
		if i < len(orderedCityNames)-1 {
			output += "\n"
		}
	}
	return
}

func (rm RoadMap) destroyCity(cityName CityName) {
	city := rm[cityName]

	message := fmt.Sprintf("%s has been destroyed by %s!", cityName, city.aliens.englishList())

	// This function deletes a connection from the given city to another city
	deleteConnections := func(subject *City, destination *City) {
		// If the destination is nil in the first place, it's not a
		// city, therefore nothing to destroy.
		if destination == nil {
			return
		}

		// Find the road back to the subject and delete the connection
		// in that direction
		if destination.roads.north == subject {
			destination.roads.north = nil
		}
		if destination.roads.east == subject {
			destination.roads.east = nil
		}
		if destination.roads.south == subject {
			destination.roads.south = nil
		}
		if destination.roads.west == subject {
			destination.roads.west = nil
		}
	}

	// Delete all roads back the city, then delete the city itself.
	deleteConnections(city, city.roads.north)
	deleteConnections(city, city.roads.east)
	deleteConnections(city, city.roads.south)
	deleteConnections(city, city.roads.west)
	delete(rm, cityName)

	fmt.Println(message)
}
