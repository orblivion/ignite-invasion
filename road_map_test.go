package main

import (
	"strings"
	"testing"
)

// Except, I need to have specific tests for specific situations. but we'll see.
// This one is good for:
// * making sure the output and input work the same? Hmm. Ordering tho.
// * fuzz/crash testing
// func generateTestMap() {
// }

func TestParseMapSuccess(t *testing.T) {
	// Incomplete on purpose. south=Bar is explicitly reciprocal, the rest should
	// be inferred.
	input := strings.Join([]string{
		"Bar north=Foo east=Qux south=Lol west=Austin",
		"Foo south=Bar",
	}, "\n")
	roadMap, err := parseMap(input)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
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
			west:    "Bar",
		}, {
			subject: "Lol",
			north:   "Bar",
		}, {
			subject: "Austin",
			east:    "Bar",
		},
	}

	for _, tt := range tests {
		// nil by default, which is what we want if a name wasn't provided
		var north *CityRoads
		var east *CityRoads
		var south *CityRoads
		var west *CityRoads

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

		if roadMap[tt.subject].North != north ||
			roadMap[tt.subject].East != east ||
			roadMap[tt.subject].South != south ||
			roadMap[tt.subject].West != west {
			t.Errorf("%s doesn't have expected neighbor cities. Got %+v", tt.subject, roadMap[tt.subject])
		}
	}
}

func TestParseMapInvalid(t *testing.T) {
	input := strings.Join([]string{
		"Bar northFoo east=Qux up=Lol west=Austin",
		"Foo south=Bar",
	}, "\n")

	_, err := parseMap(input)

	if err == nil {
		t.Fatalf("Expected err to be TODO got nil")
	}

	if err.Error() != "TODO" {
		t.Fatalf("Expected err to be TODO got: %s", err.Error())
	}
}
