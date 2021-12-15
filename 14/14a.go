package main

import (
	"bufio"
	"fmt"
	"os"
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
	replacementMap := map[string]string{}
	elementCounterMap := map[string]int{}
	elementString := ""
	counter := 0
	rounds := 10
	for scanner.Scan() {

		inputString := scanner.Text()
		if counter == 0 {
			elementString = inputString
			for _, character := range inputString {
				elementCounterMap[string(character)] += 1
			}
		} else if strings.Trim(inputString, "\n") != "" {
			fields := strings.Fields(inputString)
			match := string(fields[0])
			replacement := string(fields[2])
			replacementMap[match] = replacement

		}
		counter++
	}
	for i := rounds; i > 0; i-- {
		newElementString := ""
		for index := 0; index < len(elementString)-1; index++ {
			pairString := elementString[index : index+2]
			newElementString += string(elementString[index])

			replacementChar := replacementMap[pairString]
			elementCounterMap[replacementChar] += 1
			newElementString += replacementChar

		}
		newElementString += string(elementString[len(elementString)-1])
		elementString = newElementString

	}

	// fmt.Println(elementCounterMap)
	// fmt.Println(elementString)
	// fmt.Println(elementCounterMap)
	lowest := -1
	highest := -1
	for _, element := range elementCounterMap {
		if lowest == -1 || element <= lowest {
			lowest = element
		}
		if highest == -1 || element >= highest {
			highest = element
		}
	}

	fmt.Println(highest - lowest)
	// fmt.Println(replacementMap)
}
