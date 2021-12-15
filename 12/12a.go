package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindPaths(cave string, caveMap map[string][]string, visitedCaves map[string]bool, currentPath string) int {
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

			solutions += FindPaths(connectedCave, caveMap, visitedCaves, currentPath+","+connectedCave)
			visitedCaves[connectedCave] = false

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

	solutions := FindPaths("start", caveMap, visitedCaves, "start")

	fmt.Println(solutions)

}
