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
	var lastInput int = -1
	var counter = 0
	for scanner.Scan() {
		inputValue, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if lastInput == -1 {
			fmt.Println(inputValue)
		} else {
			if int(inputValue) > lastInput {
				fmt.Println("Increase.", lastInput, inputValue)
				counter++
			}
		}

		lastInput = int(inputValue)
	}
	fmt.Println(counter)
}
