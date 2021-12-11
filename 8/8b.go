package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// "ioutil"

func GetBinaryRep(numberString string, digitMap map[string]int) int {
	returnNumber := 0
	for _, char := range numberString {
		number := digitMap[string(char)]
		returnNumber = returnNumber | number
	}

	return returnNumber
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	digitMap := map[string]int{"a": 0b1, "b": 0b10, "c": 0b100, "d": 0b1000, "e": 0b10000, "f": 0b100000, "g": 0b1000000}
	numberSum := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Split(inputString, "|")
		inputDigits := strings.Fields(stringSplit[0])
		outputDigits := strings.Fields(stringSplit[1])

		segmentDecoder := map[string]int{}
		binaryToInt := map[int]int{}

		intToBinary := map[int]int{}

		lengthToBinary := map[int][]int{}
		for _, numberString := range inputDigits {
			numberLength := len(numberString)

			if lengthToBinary[numberLength] == nil {
				lengthToBinary[numberLength] = make([]int, 0)
			}
			lengthToBinary[numberLength] = append(lengthToBinary[numberLength], GetBinaryRep(numberString, digitMap))

		}

		// Numbers known by length of segment string
		intToBinary[1] = lengthToBinary[2][0] // The one
		binaryToInt[intToBinary[1]] = 1

		intToBinary[4] = lengthToBinary[4][0] // The four
		binaryToInt[intToBinary[4]] = 4

		intToBinary[7] = lengthToBinary[3][0] // The seven
		binaryToInt[intToBinary[7]] = 7

		intToBinary[8] = lengthToBinary[7][0] // The eight
		binaryToInt[intToBinary[8]] = 8

		// Find numbers with length six
		for _, number := range lengthToBinary[6] {
			if number == number|intToBinary[4]|intToBinary[7] {
				// Found the nine
				binaryToInt[number] = 9
				intToBinary[9] = number
				segmentDecoder["e"] = number ^ intToBinary[8]
			} else if number == number|intToBinary[1] {
				// Found the zero
				binaryToInt[number] = 0
				intToBinary[0] = number
				segmentDecoder["d"] = number ^ intToBinary[8]
			} else {
				// Found the six
				binaryToInt[number] = 6
				intToBinary[6] = number
				segmentDecoder["c"] = number ^ intToBinary[8]
			}
		}

		// Decode some egments to get the rest of the numbers
		segmentDecoder["a"] = intToBinary[7] ^ intToBinary[1]
		segmentDecoder["g"] = (intToBinary[7] | segmentDecoder["e"] | intToBinary[4]) ^ intToBinary[8]

		intToBinary[5] = intToBinary[8] ^ segmentDecoder["c"] ^ segmentDecoder["e"]
		binaryToInt[intToBinary[5]] = 5

		intToBinary[3] = segmentDecoder["a"] | segmentDecoder["d"] | segmentDecoder["g"] | intToBinary[1]
		binaryToInt[intToBinary[3]] = 3

		intToBinary[2] = segmentDecoder["a"] | segmentDecoder["c"] | segmentDecoder["d"] | segmentDecoder["e"] | segmentDecoder["g"]
		binaryToInt[intToBinary[2]] = 2

		// Decode numbers
		number := 0
		for _, numberString := range outputDigits {
			digitBinary := GetBinaryRep(numberString, digitMap)
			digitInt := binaryToInt[digitBinary]
			number *= 10

			number += digitInt

		}
		numberSum += number

	}

	fmt.Println(numberSum)

}
