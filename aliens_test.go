package main

import (
	"testing"
)

func TestSetInitialAliens(t *testing.T) {
	roadMap := make(RoadMap)

	newCity1 := City{}
	newCity2 := City{}
	newCity3 := City{}

	roadMap["Foo"] = &newCity1
	roadMap["Bar"] = &newCity2
	roadMap["Austin"] = &newCity3

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
