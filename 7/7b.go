package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// "ioutil"

func main() {

	// Parse input
	inputFile, err := os.Open("input_test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	var sum int64 = 0
	var numberCount int64 = 0
	numberArray := []int64{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Split(inputString, ",")
		for _, number := range stringSplit {
			inputNumber, _ := strconv.ParseInt(number, 10, 64)

			if numberCount == 0 {
				numberArray = append(numberArray, inputNumber)
			} else {

				var index int64 = 0
				for true {
					if index >= numberCount {
						break
					}
					if numberArray[index] >= inputNumber {
						break
					}

					index++
				}

				secondSlice := make([]int64, numberCount-index)
				var numberOfElements int

				numberOfElements = copy(secondSlice, numberArray[index:])

				numberOfElements = numberOfElements
				numberArray = append(numberArray[:index], inputNumber)

				numberArray = append(numberArray, secondSlice...)

			}

			sum += inputNumber
			numberCount++

		}
	}

	averageRounded := int64(math.Round(float64(sum) / float64(numberCount)))

	fuelSum := -1
	for i := -1; i < 2; i++ {
		temporaryFuleSum := 0
		test := averageRounded - int64(i)
		for _, element := range numberArray {
			steps := int(math.Abs(float64(element - test)))
			fuelCost := (steps * (steps + 1)) / 2
			temporaryFuleSum += fuelCost
		}

		if fuelSum == -1 || fuelSum >= temporaryFuleSum {
			fuelSum = temporaryFuleSum
		}
	}

	fmt.Println("FuelSum ", fuelSum)

}
