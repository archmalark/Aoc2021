package main

import (
	"fmt"
	"strconv"

	// "ioutil"

	"bufio"
	"os"
)

type NumberTree struct {
	numberArray [2]int
	next        *NumberTree
}

func NewNumberTree() *NumberTree {
	numberTree := new(NumberTree)
	numberTree.numberArray[0] = 0
	numberTree.numberArray[1] = 0
	numberTree.next = nil
	return numberTree
}

func GetMaxNumber(numberStrings []string, number string) string {
	var leadingZeroes []string
	var leadingOnes []string
	var numberString string

	for i := len(numberStrings) - 1; i >= 0; i-- {

		numberString = string(numberStrings[i])

		if string(numberString)[0] == '0' {
			leadingZeroes = append(leadingZeroes, numberString[1:])
		} else {
			leadingOnes = append(leadingOnes, numberString[1:])
		}

	}

	if len(numberStrings) == 1 {

		number = number + numberStrings[0]

	} else if len(leadingOnes) >= len(leadingZeroes) {
		number = number + "1"

	} else {
		number = number + "0"

	}

	if len(numberStrings) == 1 {
		return number
	}
	if len(numberString) == 1 {
		return number
	}

	if len(leadingOnes) >= len(leadingZeroes) {

		return GetMaxNumber(leadingOnes, number)

	} else {

		return GetMaxNumber(leadingZeroes, number)
	}

	return "nil"

}

func GetMinNumber(numberStrings []string, number string) string {
	var leadingZeroes []string
	var leadingOnes []string
	var numberString string

	for i := len(numberStrings) - 1; i >= 0; i-- {

		numberString = string(numberStrings[i])

		if string(numberString)[0] == '0' {
			leadingZeroes = append(leadingZeroes, numberString[1:])
		} else {
			leadingOnes = append(leadingOnes, numberString[1:])
		}

	}

	if len(numberStrings) == 1 {

		number = number + numberStrings[0]

	} else if len(leadingZeroes) <= len(leadingOnes) {
		number = number + "0"

	} else {
		number = number + "1"

	}

	if len(numberStrings) == 1 {
		return number
	}
	if len(numberString) == 1 {
		return number
	}

	if len(leadingZeroes) <= len(leadingOnes) {

		return GetMinNumber(leadingZeroes, number)

	} else {

		return GetMinNumber(leadingOnes, number)
	}

	return "nil"

}

func main() {

	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	const numberSize = 12

	var numberStrings []string

	for scanner.Scan() {
		inputString := scanner.Text()

		numberStrings = append(numberStrings, inputString)

	}

	c02, _ := strconv.ParseInt(GetMinNumber(numberStrings, ""), 2, 64)
	oxy, _ := strconv.ParseInt(GetMaxNumber(numberStrings, ""), 2, 64)

	fmt.Println(c02 * oxy)

}
