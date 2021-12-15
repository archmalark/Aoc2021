package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	xToYMap := map[int64]map[int64]bool{}
	yToXMap := map[int64]map[int64]bool{}
	numberOfDots := 0
	numberOfFolds := 1
	for scanner.Scan() {

		inputString := scanner.Text()

		if strings.Contains(inputString, "fold along") {
			fields := strings.Fields(inputString)
			split := strings.Split(fields[2], "=")
			direction := split[0]
			position, _ := strconv.ParseInt(split[1], 10, 64)

			// fmt.Println("Direction: ", direction, " Position: ", position)
			if direction == "y" {
				for yIndex, yElement := range yToXMap {
					if yIndex > position {
						// fmt.Println(yIndex, yElement)
						newY := position + position - yIndex
						for xIndex, xElement := range yElement {
							if yToXMap[newY] == nil {
								yToXMap[newY] = make(map[int64]bool)
							}
							if xToYMap[xIndex] == nil {
								xToYMap[xIndex] = make(map[int64]bool)
							}
							if xElement == true && yToXMap[newY][xIndex] == true {
								numberOfDots--
							}
							yToXMap[newY][xIndex] = xElement
							xToYMap[xIndex][newY] = xElement
							xToYMap[xIndex][yIndex] = false
							yToXMap[yIndex][xIndex] = false
						}
					}
				}
			} else if direction == "x" {
				for xIndex, xElement := range xToYMap {
					if xIndex > position {
						newX := position + position - xIndex
						for yIndex, yElement := range xElement {
							if xToYMap[newX] == nil {
								xToYMap[newX] = make(map[int64]bool)
							}
							if yToXMap[yIndex] == nil {
								yToXMap[yIndex] = make(map[int64]bool)
							}

							if yElement == true && xToYMap[newX][yIndex] == true {
								numberOfDots--
							}
							xToYMap[newX][yIndex] = yElement
							yToXMap[yIndex][newX] = yElement
							xToYMap[xIndex][yIndex] = false
							yToXMap[yIndex][xIndex] = false
						}
					}
				}
			}
			numberOfFolds--

		} else if strings.Trim(inputString, "\n") != "" {
			// fmt.Println(inputString)
			stringSplit := strings.Split(inputString, ",")
			x := stringSplit[0]
			y := stringSplit[1]
			xInt, _ := strconv.ParseInt(x, 10, 64)
			yInt, _ := strconv.ParseInt(y, 10, 64)
			if xToYMap[xInt] == nil {
				xToYMap[xInt] = make(map[int64]bool)
			}
			if yToXMap[yInt] == nil {
				yToXMap[yInt] = make(map[int64]bool)
			}
			yToXMap[yInt][xInt] = true
			xToYMap[xInt][yInt] = true
			numberOfDots++
		}

		if numberOfFolds == 0 {
			break
		}

	}

	// fmt.Println(yToXMap)
	fmt.Println(numberOfDots)

}
