package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// "ioutil"
const width = 10
const height = 10

func FlashOctos(x int, y int, xMax int, yMax int, octoArray *[width][height]int64, newOctoArray *[width][height]int64) int {
	if (*octoArray)[x][y] == -1 {
		return 0
	}
	(*octoArray)[x][y] += 1
	if (*octoArray)[x][y] >= 9 {
		(*octoArray)[x][y] = -1
		(*newOctoArray)[x][y] = 0
		numberOfFlashes := 1
		for addX := -1; addX <= 1; addX++ {
			for addY := -1; addY <= 1; addY++ {
				newX := x + addX
				newY := y + addY
				if newX >= 0 && newY >= 0 && newX <= xMax && newY <= yMax {

					numberOfFlashes += FlashOctos(newX, newY, xMax, yMax, octoArray, newOctoArray)
				}
			}
		}
		return numberOfFlashes
	} else {
		(*newOctoArray)[x][y] += 1

		return 0
	}
}
func PrintOctoArray(octoArray [width][height]int64) {
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {

			fmt.Print(octoArray[x][y])
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	octoArray := [width][height]int64{}
	rowCounter := 0
	for scanner.Scan() {

		inputString := scanner.Text()
		columnCounter := 0
		for _, numberChar := range inputString {
			number, _ := strconv.ParseInt(string(numberChar), 10, 64)

			octoArray[columnCounter][rowCounter] = number
			columnCounter++

		}
		rowCounter++
	}

	xMax := width - 1
	yMax := height - 1
	numberOfFlashes := 0
	rounds := 100
	for true {
		roundFlashes := 0
		newOctoArray := [width][height]int64{}
		for x := 0; x <= xMax; x++ {
			for y := 0; y <= yMax; y++ {
				newOctoArray[x][y] = octoArray[x][y] + 1
				if octoArray[x][y] >= 9 {

					roundFlashes += FlashOctos(x, y, xMax, yMax, &octoArray, &newOctoArray)

				}
			}
		}

		numberOfFlashes += roundFlashes

		octoArray = newOctoArray

		rounds--
		if rounds == 0 {
			break
		}
	}

	fmt.Println(numberOfFlashes)

}
