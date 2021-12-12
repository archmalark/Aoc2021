package main

import (
	"bufio"
	"fmt"
	"os"
)

// "ioutil"

func GetBasinSize(x int, y int, inputArray []string, xMax int, yMax int, expandedCoords *map[int]map[int]int) int {
	if (*expandedCoords)[y] == nil {
		(*expandedCoords)[y] = make(map[int]int)
	}

	(*expandedCoords)[y][x] = 1
	if string(inputArray[y][x]) == "9" {

		return 0
	}

	returnValue := 1

	for addX := -1; addX <= 1; addX++ {
		for addY := -1; addY <= 1; addY++ {
			if (addY == 0 || addX == 0) && (*expandedCoords)[y+addY][x+addX] == 0 {
				newX := x + addX
				newY := y + addY
				if newX >= 0 && newY >= 0 && newX <= xMax && newY <= yMax {

					returnValue += GetBasinSize(newX, newY, inputArray, xMax, yMax, expandedCoords)
				}
			}
		}
	}
	return returnValue
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	numberArray := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()

		numberArray = append(numberArray, inputString)

	}
	xMax := len(numberArray[0]) - 1
	yMax := len(numberArray) - 1
	expandedCoords := map[int]map[int]int{}

	basinSizeArray := []int{}
	basinCount := 0
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			if string(numberArray[y][x]) != "9" && expandedCoords[y][x] != 1 {
				basinSize := GetBasinSize(x, y, numberArray, xMax, yMax, &expandedCoords)

				if basinCount == 0 {
					basinSizeArray = append(basinSizeArray, basinSize)
				} else {

					var index int = 0
					for true {
						if index >= basinCount {
							break
						}
						if basinSizeArray[index] <= basinSize {
							break
						}

						index++
					}

					secondSlice := make([]int, basinCount-index)

					var numberOfElements int

					numberOfElements = copy(secondSlice, basinSizeArray[index:])

					numberOfElements = numberOfElements
					basinSizeArray = append(basinSizeArray[:index], basinSize)

					basinSizeArray = append(basinSizeArray, secondSlice...)

				}
				basinCount++
			}
		}
	}

	fmt.Println(basinSizeArray[0] * basinSizeArray[1] * basinSizeArray[2])

}
