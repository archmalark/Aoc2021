package main

import (
	"bufio"
	"fmt"
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

func (stack *Stack) Pop() string {
	returnValue := stack.stack[stack.len-1]
	stack.stack = stack.stack[0 : stack.len-1]
	stack.len -= 1
	return returnValue
}

func (stack *Stack) Peek() string {
	returnValue := stack.stack[stack.len-1]
	return returnValue
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	characterMatch := map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
	pointMap := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	pointSum := 0
	for scanner.Scan() {
		lineStack := NewStack()
		inputString := scanner.Text()
		for _, character := range inputString {
			characterString := string(character)

			value, keyExists := characterMatch[characterString]

			if keyExists {
				lineStack.Push(characterString)
			} else {
				openingString := lineStack.Peek()
				value, keyExists = characterMatch[openingString]

				if !keyExists {
					fmt.Errorf("Last stack character is not an opening something is wrong")
				}

				if value == characterString {
					lineStack.Pop()
				} else {
					pointSum += pointMap[characterString]
					break
				}

			}

		}

	}

	fmt.Println(pointSum)

}
