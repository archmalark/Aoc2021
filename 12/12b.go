package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindPaths(cave string, caveMap map[string][]string, visitedCaves map[string]bool, currentPath string, doubleVisit bool) int {
	if cave == "end" {
		// fmt.Println(currentPath)
		return 1
	}

	if strings.ToLower(cave) == cave {
		(visitedCaves)[cave] = true
	}

	solutions := 0
	for _, connectedCave := range caveMap[cave] {
		if (visitedCaves)[connectedCave] != true {

			solutions += FindPaths(connectedCave, caveMap, visitedCaves, currentPath+","+connectedCave, doubleVisit)
			visitedCaves[connectedCave] = false

		} else if doubleVisit == false && connectedCave != "start" {
			solutions += FindPaths(connectedCave, caveMap, visitedCaves, currentPath+","+connectedCave, true)
		}

	}
	return solutions
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	caveMap := map[string][]string{}
	visitedCaves := map[string]bool{}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		inputString := scanner.Text()

		stringSplit := strings.Split(inputString, "-")

		caveOne := stringSplit[0]
		caveTwo := stringSplit[1]

		caveMap[caveOne] = append(caveMap[caveOne], caveTwo)
		caveMap[caveTwo] = append(caveMap[caveTwo], caveOne)
	}

	solutions := FindPaths("start", caveMap, visitedCaves, "start", false)

	fmt.Println(solutions)

}
