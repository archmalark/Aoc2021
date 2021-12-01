package main

import (
	"fmt"
	"strconv"

	// "ioutil"

	"bufio"
	"os"
)

func main() {

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// input := ioutil.ReadAll(inputFile)
	scanner := bufio.NewScanner(inputFile)
	const windowSize = 3
	var counter = 0
	var increaseCounter = 0
	var inputList [windowSize]int

	// inputList[0] = 0
	// inputList[1] = 0
	// inputList[2] = 0
	var lastWindow = -1
	for scanner.Scan() {
		inputValue, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		for i := 0; i < 3; i++ {

			inputList[i] += int(inputValue)

		}

		counter++
		if counter > windowSize-1 {
			if lastWindow > -1 && lastWindow < inputList[counter%windowSize] {
				increaseCounter += 1

			}
			lastWindow = inputList[counter%windowSize]

		}
		inputList[counter%windowSize] = 0

	}

	fmt.Println(increaseCounter)
}
