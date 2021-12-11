package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// "ioutil"

const numberOfDays int = 256
const cycleLength int = 9

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	var lanternfishList [cycleLength]int64
	var fishCount int64 = 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Split(inputString, ",")
		for _, number := range stringSplit {
			inputNumber, _ := strconv.ParseInt(number, 10, 64)
			lanternfishList[inputNumber] += 1
			fishCount += 1
		}
	}

	// Solve puzzle
	fmt.Println(lanternfishList)

	for day := 0; day < numberOfDays; day++ {
		newFishes := lanternfishList[0]
		for i := 1; i <= cycleLength-1; i++ {
			lanternfishList[i-1] = lanternfishList[i]
		}
		lanternfishList[6] += newFishes
		lanternfishList[8] = newFishes
		fishCount += newFishes
	}

	fmt.Println(fishCount)
}
