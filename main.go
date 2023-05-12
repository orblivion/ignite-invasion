package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const SimulationLength = 10000

func runSimulation(numAliens int, numIterations int, mapInput string) error {
	roadMap, err := parseMap(mapInput)
	if err != nil {
		return err
	}

	roadMap.setInitialAliens(numAliens)
	for i := 0; i < numIterations; i++ {
		// Fight first. This means that if two aliens START in the same
		// city, they all will destroy each other and the city.
		roadMap.fightAliens()
		roadMap.moveAliens()
	}

	fmt.Println(roadMap.outputMap())
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 3 {
		fmt.Printf("Expected 2 arguments: fileName, numAliens (got %d argument[s])\n", len(os.Args)-1)
		os.Exit(1)
	}
	fileName := os.Args[1]
	numAliensStr := os.Args[2]

	mapInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading map file: ", err.Error())
		os.Exit(1)
	}

	numAliens, err := strconv.Atoi(numAliensStr)
	if err != nil {
		fmt.Println("Error parsing number of aliens: ", err.Error())
		os.Exit(1)
	}

	runSimulation(numAliens, SimulationLength, string(mapInput))
}
