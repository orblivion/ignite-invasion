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
			expectedErr: "Invalid line in initial map (expecting = in road description): Bar northFoo west=Austin",
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
				"Bar west=Foo",
			},
			expectedErr: "Invalid line in initial map (road to undefined city): Bar west=Foo",
		},
		{
			inputLines: []string{
				"Bar%",
			},

			// TODO - comment in implementation, alphanumeric and -_
			expectedErr: "Invalid line in initial map (invalid city name): Bar%",
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

			expectedErr: "Invalid initial map (not connected in both directions)",
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
