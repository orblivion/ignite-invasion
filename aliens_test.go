package main

import (
	"fmt"
	"testing"
)

func TestSetInitialAliens(t *testing.T) {
	roadMap := make(RoadMap)

	foo := City{}
	bar := City{}
	austin := City{}

	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin

	roadMap.setInitialAliens(70)

	dupeCheck := make(map[AlienName]bool)

	allAliens := append(roadMap["Foo"].Aliens, roadMap["Bar"].Aliens...)
	allAliens = append(allAliens, roadMap["Austin"].Aliens...)

	for _, alien := range allAliens {
		if _, ok := dupeCheck[alien]; ok {
			t.Errorf("An alien showed up twice somewhere")
		}
		dupeCheck[alien] = true
	}

	if len(dupeCheck) != 70 {
		t.Errorf("Expected 70 aliens to get created")
	}
}

func TestMoveAliens(t *testing.T) {
	roadMap := make(RoadMap)

	// Instead of generating names, let's make convenient ones for
	// debugging.
	alienFromAustinStr := "alienFromAustin"
	alienFromBarStr := "alienFromBar"
	alienFromFooStr := "alienFromFoo"
	alienFromLonerCityStr := "alienFromLonerCity"

	alienFromAustin := AlienName(&alienFromAustinStr)
	alienFromBar := AlienName(&alienFromBarStr)
	alienFromFoo := AlienName(&alienFromFooStr)
	alienFromLonerCity := AlienName(&alienFromLonerCityStr)

	foo := City{Aliens: []AlienName{alienFromFoo}}
	bar := City{Aliens: []AlienName{alienFromBar}}
	austin := City{Aliens: []AlienName{alienFromAustin}}
	lonerCity := City{Aliens: []AlienName{alienFromLonerCity}}

	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin
	roadMap["LonerCity"] = &lonerCity // no roads in or out

	roadMap["Foo"].Roads.North = roadMap["Bar"]
	roadMap["Bar"].Roads.South = roadMap["Foo"]
	roadMap["Austin"].Roads.East = roadMap["Bar"]
	roadMap["Bar"].Roads.East = roadMap["Austin"]

	// The map looks like this:
	//
	// Austin <-> Bar <-> Foo     LonerCity
	//
	// We expect:
	// * The aliens in Austin and Foo to travel to Bar
	// * The alien in Bar to travel to Austin or Foo
	// * The alien in LonerCity to stay put

	// Since just about every error output could benefit from this as debug
	// output, make it into a function. Don't waste time generating the
	// text unless we need it.
	getDebugMap := func() string {
		output := ""
		for name, city := range roadMap {
			output += fmt.Sprintf("%s:\n", name)
			for _, alien := range city.Aliens {
				output += fmt.Sprintf(" %s\n", *alien)
			}
		}
		return output
	}

	roadMap.moveAliens()

	// Check that alienFromAustin and alienFromFoo both ended up in Bar
	if numAliens := len(roadMap["Bar"].Aliens); numAliens != 2 {
		t.Errorf("Expected 2 aliens to end up in Bar: \n%s", getDebugMap())
	} else {
		// To avoid a panic (harder to read test results), only bother
		// with these if we have the right number of aliens in Bar
		if roadMap["Bar"].Aliens[0] != alienFromAustin && roadMap["Bar"].Aliens[1] != alienFromAustin {
			t.Errorf("Expected alienFromAustin to end up in Bar: \n%s", getDebugMap())
		}
		if roadMap["Bar"].Aliens[0] != alienFromFoo && roadMap["Bar"].Aliens[1] != alienFromFoo {
			t.Errorf("Expected alienFromFoo to end up in Bar: \n%s", getDebugMap())
		}
	}

	// Check that alienFromBar ended up in Austin or Foo
	//
	// Combine the city aliens to test for both at the same time conveniently
	austinAndFooAliens := append(roadMap["Foo"].Aliens, roadMap["Austin"].Aliens...)

	if len(austinAndFooAliens) != 1 || austinAndFooAliens[0] != alienFromBar {
		t.Errorf("Expected alienFromBar to end up in (only) Austin or Foo: \n%s", getDebugMap())
	}

	// Check that alienFromLonerCity stayed put
	if len(roadMap["LonerCity"].Aliens) != 1 || roadMap["LonerCity"].Aliens[0] != alienFromLonerCity {
		t.Errorf("Expected alienFromLonerCity to stay alone in LonerCity: \n%s", getDebugMap())
	}
}

// Test that cities get destroyed when two or more aliens inhabit it.
// Testing that the correct roads get destroyed is part of the destroyCity test
func TestFightAliens(t *testing.T) {

	roadMap := make(RoadMap)

	// 0 aliens
	lol := City{Aliens: []AlienName{}}

	// 1 alien
	foo := City{Aliens: []AlienName{generateAlienName()}}

	// 2 aliens
	bar := City{Aliens: []AlienName{generateAlienName(), generateAlienName()}}

	// 3 aliens
	austin := City{Aliens: []AlienName{generateAlienName(), generateAlienName(), generateAlienName()}}

	roadMap["Lol"] = &lol
	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin

	roadMap.fightAliens()

	if len(roadMap) != 2 {
		t.Errorf("Expected two cities to remain")
	}
	if _, ok := roadMap["Lol"]; !ok {
		t.Errorf("Expected Lol to remain")
	}
	if _, ok := roadMap["Foo"]; !ok {
		t.Errorf("Expected Foo to remain")
	}
}
