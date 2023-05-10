package main

type CityName string
type CityRoads struct {
	North &CityRoads
	East &CityRoads
	South &CityRoads
	West &CityRoads
}
type RoadMap map[CityName]CityRoads

func parseMap () City c {
}

func printMap () City c {
}

func destroyCity () {
	// ... follow links, set nils, delete entry from Map. That will kill the aliens inside.
	// Bar has been destroyed by Goomkormonzor and Thublarkorxan!
}


