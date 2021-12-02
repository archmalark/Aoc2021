package main

import (
	"fmt"
	"strconv"
	"strings"

	// "ioutil"

	"bufio"
	"os"
)

type Coord struct {
	x int
	y int
}

func NewCoord(x int, y int) *Coord {
	coord := new(Coord)
	coord.x = x
	coord.y = y
	return coord
}

func (coord1 *Coord) Add(coord2 *Coord) Coord {
	coord1.x += coord2.x
	coord1.y += coord2.y
	return *coord1
}

func (coord1 *Coord) Multiply(m int) Coord {
	coord1.x *= m
	coord1.y *= m
	return *coord1
}

func main() {

	coord := NewCoord(0, 0)
	aim := 0

	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// input := ioutil.ReadAll(inputFile)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Split(inputString, " ")
		directionString := stringSplit[0]
		speed, _ := strconv.ParseInt(stringSplit[1], 10, 64)

		switch directionString {
		case "forward":
			coord.Add(NewCoord(int(speed), int(speed)*aim))
		case "up":
			aim -= int(speed)
		case "down":
			aim += int(speed)
		}

	}

	fmt.Println(coord.x * coord.y)

}
