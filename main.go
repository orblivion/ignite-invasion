package main

const SimulationLength = 10000

// Assumptions:

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

// If 3+ aliens can land on the same city. They all get destroyed in the fight?

// If two aliens START in the same city, they all will destroy each other and the city.

// Each alien *needs* to move every turn, unless it's trapped.

func main() {
}
