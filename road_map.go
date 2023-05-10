package main

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

func parseMap(s string) (rm RoadMap, err error) {
	rm = make(RoadMap)
	return rm, nil
}

func printMap(rm RoadMap) {
}

func destroyCity() {
	// ... follow links, set nils, delete entry from Map. That will kill the aliens inside.
	// Bar has been destroyed by Goomkormonzor and Thublarkorxan!
}
