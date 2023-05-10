package main

// TODO Naming convention
const SimulationLength = 10000

// Assumptions:

// Are roads automatically two-way? Suppose the map says Foo has north=Bar, but Bar doesn't have south=Foo. Should I assume:
// * Bar south=Foo implicitly
// * A one-way road north from Foo to Bar
// * An invalid file
// For now I will assume the implicit two-way. If I have time maybe I'll change it to allow for one-way roads.

// If 3+ aliens can land on the same city. They all get destroyed in the fight?

// If two aliens START in the same city, they all will destroy each other and the city.

// Each alien *needs* to move every turn, unless it's trapped.

func main() {
}
