package main

import (
	"fmt"
	"math"

	// "ioutil"
	"bufio"
	"os"
	"strconv"
)

func main() {

	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// input := ioutil.ReadAll(inputFile)
	scanner := bufio.NewScanner(inputFile)
	const numberSize = 12

	var accumulatorList [numberSize]int64
	var counter int64 = 0
	for scanner.Scan() {
		inputString := scanner.Text()
		inputNumber, _ := strconv.ParseInt(inputString, 2, 64)

		var mask int64 = 1
		for i := 0; i < numberSize; i++ {
			maskShifted := mask << i
			accumulatorList[i] += inputNumber & maskShifted

		}
		counter++

	}
	var mask int64 = 1
	binaryString := ""
	for index, element := range accumulatorList {
		maskShifted := mask << index
		halfMax := maskShifted * counter / 2
		if element > halfMax {
			binaryString = "1" + binaryString
		} else {
			binaryString = "0" + binaryString
		}
		// fmt.Println("MaskShifted: ", strconv.FormatInt(maskShifted, 2))
		// fmt.Println("HalfMax: ", strconv.FormatInt(, 2))
		// fmt.Println("Acc: ", strconv.FormatInt(element, 2))
		// fmt.Println("------------")

	}
	gamma, _ := strconv.ParseInt(binaryString, 2, 64)
	mask = int64(math.Pow(2, numberSize)) - 1

	epsilon := (^gamma) & mask

	// fmt.Println(gamma)
	// fmt.Println(epsilon)
	fmt.Println(epsilon * gamma)
}

// 1100000
// 	 10000
