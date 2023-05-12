package main

import (
	"strings"
	"testing"
)

func TestParseMapSuccess(t *testing.T) {
	input := strings.Join([]string{
		"Bar north=Foo east=Qux south=Lol west=Austin",
		"Foo south=Bar",
		"Austin north=Bar",
		"Lol east=Bar",
		"Qux south=Bar",
	}, "\n")
	roadMap, err := parseMap(input)

	if err != nil {
		t.Fatalf("Unexpected error: `%s`", err.Error())
	}

	if l := len(roadMap); l != 5 {
		t.Fatalf("Expected 5 cities, got %d. %+v", l, roadMap)
	}

	tests := []struct {
		subject CityName
		north   CityName
		south   CityName
		east    CityName
		west    CityName
	}{
		{
			subject: "Bar",
			north:   "Foo",
			east:    "Qux",
			south:   "Lol",
			west:    "Austin",
		}, {
			subject: "Foo",
			south:   "Bar",
		}, {
			subject: "Qux",
			south:   "Bar",
		}, {
			subject: "Lol",
			east:    "Bar",
		}, {
			subject: "Austin",
			north:   "Bar",
		},
	}

	for _, tt := range tests {
		// nil by default, which is what we want if a name wasn't provided
		var north *City
		var east *City
		var south *City
		var west *City

		if tt.north != "" {
			north = roadMap[tt.north]
		}
		if tt.east != "" {
			east = roadMap[tt.east]
		}
		if tt.south != "" {
			south = roadMap[tt.south]
		}
		if tt.west != "" {
			west = roadMap[tt.west]
		}

		if roadMap[tt.subject].roads.north != north ||
			roadMap[tt.subject].roads.east != east ||
			roadMap[tt.subject].roads.south != south ||
			roadMap[tt.subject].roads.west != west {
			t.Errorf("%s doesn't have expected neighbor cities. Got %+v", tt.subject, roadMap[tt.subject])
		}
	}
}

func TestParseMapInvalid(t *testing.T) {
	tests := []struct {
		inputLines  []string
		expectedErr string
	}{
		{
			inputLines: []string{
				"Austin",
				"Foo",
				"Bar northFoo west=Austin",
			},
			expectedErr: "Invalid line in initial map (expecting one = in road description): Bar northFoo west=Austin",
		},
		{
			inputLines: []string{
				"Austin",
				"Foo",
				"Bar west=Austin west=Foo",
			},
			expectedErr: "Invalid line in initial map (road direction repeated): Bar west=Austin west=Foo",
		},
		{
			inputLines: []string{
				"Foo",
				"Bar up=Foo",
			},
			expectedErr: "Invalid line in initial map (invalid direction): Bar up=Foo",
		},

		{
			inputLines: []string{
				"Bar west=Foo",
			},
			expectedErr: "Invalid line in initial map (road to undefined city): Bar west=Foo",
		},
		{
			inputLines: []string{
				"Bar#",
			},

			expectedErr: "Invalid line in initial map (invalid city name): Bar#",
		},
		{
			inputLines: []string{
				"Bar",
				"Foo",
				"Bar west=Foo",
			},

			expectedErr: "Invalid line in initial map (repeated city): Bar west=Foo",
		},
		{
			inputLines: []string{
				"Bar west=Bar",
			},

			expectedErr: "Invalid line in initial map (road from city to itself): Bar west=Bar",
		},
		{
			inputLines: []string{
				"Bar west=Foo east=Foo",
				"Foo",
			},

			expectedErr: "Invalid line in initial map (two roads to the same town): Bar west=Foo east=Foo",
		},
		{
			inputLines: []string{
				"Bar west=Foo",
				"Foo",
			},

			expectedErr: "Invalid initial map (some roads are not connected on both ends)",
		},
		{
			inputLines: []string{},

			expectedErr: "Invalid initial map (empty)",
		},
	}

	for _, tt := range tests {
		input := strings.Join(tt.inputLines, "\n")

		_, err := parseMap(input)

		if err == nil {
			t.Errorf("Expected err to be `%s` got nil", tt.expectedErr)
		} else if err.Error() != tt.expectedErr {
			t.Errorf("Expected err to be `%s` got: `%s`", tt.expectedErr, err.Error())
		}
	}
}

func TestDestroyCity(t *testing.T) {
	roadMap := make(RoadMap)

	foo := City{}
	bar := City{}
	austin := City{}

	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin

	roadMap["Foo"].roads.north = roadMap["Bar"]
	roadMap["Bar"].roads.south = roadMap["Foo"]
	roadMap["Austin"].roads.east = roadMap["Bar"]
	roadMap["Bar"].roads.east = roadMap["Austin"]

	roadMap.destroyCity("Foo")

	if len(roadMap) != 2 {
		t.Errorf("Expected Bar and Austin only to remain after destroying Foo")
	}
	if _, ok := roadMap["Bar"]; !ok {
		t.Errorf("Expected Bar and Austin only to remain after destroying Foo")
	}
	if _, ok := roadMap["Austin"]; !ok {
		t.Errorf("Expected Bar and Austin only to remain after destroying Foo")
	}

	expectedBar := CityRoads{east: roadMap["Austin"]}
	if roadMap["Bar"].roads != expectedBar {
		t.Errorf("Expected Bar to have one road, east to Austin")
	}
}

func TestOutputMap(t *testing.T) {
	roadMap := make(RoadMap)

	foo := City{}
	bar := City{}
	austin := City{}

	roadMap["Foo"] = &foo
	roadMap["Bar"] = &bar
	roadMap["Austin"] = &austin

	roadMap["Foo"].roads.north = roadMap["Bar"]
	roadMap["Bar"].roads.south = roadMap["Foo"]
	roadMap["Austin"].roads.east = roadMap["Bar"]
	roadMap["Bar"].roads.east = roadMap["Austin"]

	expected := strings.Join([]string{
		"Austin east=Bar",
		"Bar east=Austin south=Foo",
		"Foo north=Bar",
	}, "\n")

	output := roadMap.outputMap()
	if output != expected {
		t.Errorf("Map output not as expected. Got:\n%s", output)
	}
}
