package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// "ioutil"

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

		// fmt.Println("INPUT: ", inputDigits)
		// fmt.Println("OUTPUT: ", outputDigits)
		// fmt.Println("_-----------------------")
	}

	width := len(numberArray[0])
	height := len(numberArray)
	lowPointSum := int64(0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			digit, _ := strconv.ParseInt(string(numberArray[y][x]), 10, 64)
			lowPoint := true
			adjacent := int64(0)

			if x > 0 {
				adjacent, _ = strconv.ParseInt(string(numberArray[y][x-1]), 10, 64)
				if digit >= adjacent {
					lowPoint = false
				}
			}

			if x < width-1 {

				adjacent, _ = strconv.ParseInt(string(numberArray[y][x+1]), 10, 64)
				if digit >= adjacent {
					lowPoint = false
				}
			}

			if y > 0 {
				adjacent, _ = strconv.ParseInt(string(numberArray[y-1][x]), 10, 64)
				if digit >= adjacent {
					lowPoint = false
				}
			}

			if y < height-1 {
				adjacent, _ = strconv.ParseInt(string(numberArray[y+1][x]), 10, 64)
				if digit >= adjacent {
					lowPoint = false
				}
			}
			if lowPoint {

				lowPointSum += digit + 1
			}

		}
	}

	fmt.Println(lowPointSum)

}
