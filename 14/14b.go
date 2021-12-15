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
	pairMap := map[string]int{}
	// elementString := ""
	counter := 0
	rounds := 40
	for scanner.Scan() {

		inputString := scanner.Text()
		if counter == 0 {

			for index := 0; index < len(inputString)-1; index++ {
				elementCounterMap[string(inputString[index])] += 1
				pairString := inputString[index : index+2]
				pairMap[pairString] += 1

			}
			elementCounterMap[string(inputString[len(inputString)-1])] += 1
		} else if strings.Trim(inputString, "\n") != "" {
			fields := strings.Fields(inputString)
			match := string(fields[0])
			replacement := string(fields[2])
			replacementMap[match] = replacement

		}
		counter++
	}
	for i := rounds; i > 0; i-- {
		newPairMap := map[string]int{}
		for key, value := range pairMap {
			replacementChar := string(replacementMap[key])
			if replacementChar != "" {
				newPairOne := string(key[0]) + replacementChar
				newPairTwo := replacementChar + string(key[1])
				newPairMap[newPairOne] += value
				newPairMap[newPairTwo] += value
				elementCounterMap[replacementChar] += value

			}
		}
		pairMap = newPairMap
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

	// fmt.Println(pairMap)
	fmt.Println(highest - lowest)
	// fmt.Println(replacementMap)
}
