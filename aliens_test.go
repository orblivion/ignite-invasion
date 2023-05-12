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

	allAliens := append(roadMap["Foo"].aliens, roadMap["Bar"].aliens...)
	allAliens = append(allAliens, roadMap["Austin"].aliens...)

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

	foo := City{aliens: []AlienName{alienFromFoo}}
	bar := City{aliens: []AlienName{alienFromBar}}
	austin := City{aliens: []AlienName{alienFromAustin}}
	lonerCity := City{aliens: []AlienName{alienFromLonerCity}}

	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin
	roadMap["LonerCity"] = &lonerCity // no roads in or out

	roadMap["Foo"].roads.north = roadMap["Bar"]
	roadMap["Bar"].roads.south = roadMap["Foo"]
	roadMap["Austin"].roads.east = roadMap["Bar"]
	roadMap["Bar"].roads.east = roadMap["Austin"]

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
			for _, alien := range city.aliens {
				output += fmt.Sprintf(" %s\n", *alien)
			}
		}
		return output
	}

	roadMap.moveAliens()

	// Check that alienFromAustin and alienFromFoo both ended up in Bar
	if numAliens := len(roadMap["Bar"].aliens); numAliens != 2 {
		t.Errorf("Expected 2 aliens to end up in Bar: \n%s", getDebugMap())
	} else {
		// To avoid a panic (harder to read test results), only bother
		// with these if we have the right number of aliens in Bar
		if roadMap["Bar"].aliens[0] != alienFromAustin && roadMap["Bar"].aliens[1] != alienFromAustin {
			t.Errorf("Expected alienFromAustin to end up in Bar: \n%s", getDebugMap())
		}
		if roadMap["Bar"].aliens[0] != alienFromFoo && roadMap["Bar"].aliens[1] != alienFromFoo {
			t.Errorf("Expected alienFromFoo to end up in Bar: \n%s", getDebugMap())
		}
	}

	// Check that alienFromBar ended up in Austin or Foo
	//
	// Combine the city aliens to test for both at the same time conveniently
	austinAndFooAliens := append(roadMap["Foo"].aliens, roadMap["Austin"].aliens...)

	if len(austinAndFooAliens) != 1 || austinAndFooAliens[0] != alienFromBar {
		t.Errorf("Expected alienFromBar to end up in (only) Austin or Foo: \n%s", getDebugMap())
	}

	// Check that alienFromLonerCity stayed put
	if len(roadMap["LonerCity"].aliens) != 1 || roadMap["LonerCity"].aliens[0] != alienFromLonerCity {
		t.Errorf("Expected alienFromLonerCity to stay alone in LonerCity: \n%s", getDebugMap())
	}
}

// Test that cities get destroyed when two or more aliens inhabit it.
// Testing that the correct roads get destroyed is part of the destroyCity test
func TestFightAliens(t *testing.T) {

	roadMap := make(RoadMap)

	// 0 aliens
	lol := City{aliens: []AlienName{}}

	// 1 alien
	foo := City{aliens: []AlienName{generateAlienName(0)}}

	// 2 aliens
	bar := City{aliens: []AlienName{generateAlienName(1), generateAlienName(2)}}

	// 3 aliens
	austin := City{aliens: []AlienName{generateAlienName(3), generateAlienName(4), generateAlienName(5)}}

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

func TestGenerateAlienName(t *testing.T) {
	tests := map[string]AlienName{
		"Goomkormonzor": generateAlienName(0),
		"Korgoommonzor": generateAlienName(1),

		// Exhaust all combinations, start over with numbers
		"Goomkormonzor 2": generateAlienName(9 * 8 * 7 * 6),
		"Korgoommonzor 2": generateAlienName(9*8*7*6 + 1),
		"Goomkormonzor 3": generateAlienName(9 * 8 * 7 * 6 * 2),
		"Korgoommonzor 3": generateAlienName(9*8*7*6*2 + 1),
	}
	for want, got := range tests {
		if want != *got {
			t.Errorf("Expected %s got %s", want, *got)
		}
	}

	// Test that alien names are unique

	// Using `string` since AlienName is a pointer to string. We want to know
	// whether the underlying strings are unique
	uniqueNames := make(map[string]bool)
	for x := 0; x < 50000; x++ {
		newName := generateAlienName(x)
		if _, ok := uniqueNames[*newName]; ok {
			t.Errorf("Duplicated name %s with seed %d", *newName, x)
		}
		uniqueNames[*newName] = true
	}
}

func TestAliensString(t *testing.T) {
	a := "a"
	b := "b"
	c := "c"
	d := "d"

	tests := map[string]CityAliens{
		"":              CityAliens{},
		"a":             CityAliens{&a},
		"a and b":       CityAliens{&a, &b},
		"a, b and c":    CityAliens{&a, &b, &c},
		"a, b, c and d": CityAliens{&a, &b, &c, &d},
	}

	for want, aliens := range tests {
		got := aliens.englishList()
		if want != got {
			t.Errorf("Expected %s got %s", want, got)
		}
	}
}
