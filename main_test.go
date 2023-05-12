package main

import (
	"testing"
)

func TestSimulation(t *testing.T) {
	// We have no particular expected output. We'll just try running it a bunch of
	// times to make sure it doesn't crash.

	for x := 0; x < 100; x++ {
		roadMap := make(RoadMap)

		foo := City{}
		bar := City{}
		qux := City{}
		austin := City{}

		roadMap["Foo"] = &foo
		roadMap["Bar"] = &bar
		roadMap["Qux"] = &qux
		roadMap["Austin"] = &austin

		roadMap["Foo"].roads.north = roadMap["Bar"]
		roadMap["Bar"].roads.south = roadMap["Foo"]

		roadMap["Austin"].roads.east = roadMap["Bar"]
		roadMap["Bar"].roads.west = roadMap["Austin"]

		roadMap["Qux"].roads.west = roadMap["Foo"]
		roadMap["Foo"].roads.east = roadMap["Qux"]

		mapInput := roadMap.outputMap()

		// 3 aliens, 10 iterations. Surely something gets destroyed.
		runSimulation(3, 10, mapInput)
	}
}
