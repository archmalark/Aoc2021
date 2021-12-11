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
	inputFile, err := os.Open("input.txt")

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

	median := int64(0)
	if numberCount%2 == 0 {
		median = (numberArray[numberCount/2] + numberArray[(numberCount/2)-1]) / 2
	} else {
		median = numberArray[int64(math.Ceil(float64(numberCount)/2))]
	}

	fuelSum := 0
	for _, element := range numberArray {
		fuelSum += int(math.Abs(float64(element - median)))
	}

	fmt.Println("FuelSum ", fuelSum)

}
