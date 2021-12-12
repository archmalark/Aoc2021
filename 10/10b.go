package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// "ioutil"

type Stack struct {
	stack []string
	len   int
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.len = 0
	return stack
}

func (stack *Stack) Push(value string) Stack {
	stack.stack = append(stack.stack, value)
	stack.len += 1
	return *stack
}

func (stack *Stack) Pop() (string, bool) {
	if stack.len > 0 {
		returnValue := stack.stack[stack.len-1]
		stack.stack = stack.stack[0 : stack.len-1]
		stack.len -= 1
		return returnValue, true
	}
	return "", false

}

func (stack *Stack) Peek() (string, bool) {
	if stack.len > 0 {
		returnValue := stack.stack[stack.len-1]
		return returnValue, true
	}
	return "", true

}

func main() {

	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	characterMatch := map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
	pointMap := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	pointSumList := []int{}
	pointSumCount := 0
	for scanner.Scan() {
		validLine := true
		lineStack := NewStack()
		inputString := scanner.Text()
		for _, character := range inputString {
			characterString := string(character)

			value, keyExists := characterMatch[characterString]

			if keyExists {
				lineStack.Push(characterString)
			} else {
				openingString, _ := lineStack.Peek()
				value, keyExists = characterMatch[openingString]

				if !keyExists {
					fmt.Errorf("Last stack character is not an opening something is wrong")
				}

				if value == characterString {
					lineStack.Pop()
				} else {
					// pointSum += pointMap[characterString]
					validLine = false
					break
				}

			}

		}

		if validLine {
			pointSum := 0
			for value, ok := lineStack.Pop(); ok; value, ok = lineStack.Pop() {
				completionCharacter := characterMatch[value]
				pointSum *= 5
				pointSum += pointMap[completionCharacter]
			}

			if pointSumCount == 0 {
				pointSumList = append(pointSumList, pointSum)
			} else {

				var index int = 0
				for true {
					if index >= pointSumCount {
						break
					}
					if pointSumList[index] <= pointSum {
						break
					}

					index++
				}

				secondSlice := make([]int, pointSumCount-index)

				var numberOfElements int

				numberOfElements = copy(secondSlice, pointSumList[index:])

				numberOfElements = numberOfElements
				pointSumList = append(pointSumList[:index], pointSum)

				pointSumList = append(pointSumList, secondSlice...)

			}
			pointSumCount++

		}

	}
	middle := int(math.Floor(float64(pointSumCount) / 2))
	fmt.Println(pointSumList[middle])

}
