package main

import "fmt"

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
}
