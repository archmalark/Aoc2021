package main

import (
	"bufio"
	"fmt"
	"os"
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

	numberCount := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Split(inputString, "|")
		// inputDigits := strings.Fields(stringSplit[0])
		outputDigits := strings.Fields(stringSplit[1])

		for _, number := range outputDigits {
			numberLength := len(number)
			if numberLength == 2 || numberLength == 3 || numberLength == 4 || numberLength == 7 {
				// fmt.Println(number)
				numberCount++
			}
		}
		// fmt.Println("INPUT: ", inputDigits)
		// fmt.Println("OUTPUT: ", outputDigits)
		// fmt.Println("_-----------------------")
	}

	fmt.Println(numberCount)

}
