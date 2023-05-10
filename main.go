package main

type CityName string
type CityRoads struct {
	North &CityRoads
	East &CityRoads
	South &CityRoads
	West &CityRoads
}
type RoadMap map[CityName]CityRoads

// Choosing to attach the aliens to the city instead of attaching the city to the aliens because it's simpler to find alien collisions
type AlienName string
type CityAliens []AlienName
type AlienMap map[*CityRoads]CityAliens

func parseMap () City c {
}

func printMap () City c {
}

func moveAliens(thisMap AlienMap) (nextMap AlienMap) {
	// In order to keep track of which aliens have already moved, we are making a new AlienMap each time
	// TODO - maybe we map alien to city ... but then we need to have both at least temporarily, for simplicity
}

func destroyCity () {
	// ... follow links, set nils, delete entry from Map. That will kill the aliens inside.
	// Bar has been destroyed by Goomkormonzor and Thublarkorxan!
}

// TODO Naming convention
SIMULATION_LENGTH = 10000

func generateAlienName {
	// Pick 4 different name segments, order matters, that's 3000 alien names.
	// If we need more, we can append numbers

	// Capitalize the first letter

	"goom"
	"kor"
	"mon"
	"zor"
	"xan"
	"blax"
	"thu"
	"blar"
	"yaf"

}

// Assumptions:

// Are roads automatically two-way? Suppose the map says Foo has north=Bar, but Bar doesn't have south=Foo. Should I assume:
// * Bar south=Foo implicitly
// * A one-way road north from Foo to Bar
// * An invalid file
// For now I will assume the implicit two-way. If I have time maybe I'll change it to allow for one-way roads.

// If 3+ aliens can land on the same city. They all get destroyed in the fight?

// If two aliens START in the same city, they all will destroy each other and the city.

// Each alien *needs* to move every turn, unless it's trapped.
