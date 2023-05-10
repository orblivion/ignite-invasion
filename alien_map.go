package main

// Choosing to attach the aliens to the city instead of attaching the city to the aliens because it's simpler to find alien collisions
type AlienName string
type CityAliens []AlienName
type AlienMap map[*CityRoads]CityAliens

func moveAliens(thisMap AlienMap) (nextMap AlienMap) {
	// In order to keep track of which aliens have already moved, we are making a new AlienMap each time
	// TODO - maybe we map alien to city ... but then we need to have both at least temporarily, for simplicity
}


