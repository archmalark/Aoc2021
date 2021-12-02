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

	directionDict := map[string]Coord{
		"forward": *NewCoord(1, 0),
		"up":      *NewCoord(0, -1),
		"down":    *NewCoord(0, 1),
	}

	coord := NewCoord(0, 0)

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
		directionCoord := directionDict[directionString]
		speed, _ := strconv.ParseInt(stringSplit[1], 10, 64)
		move := directionCoord.Multiply(int(speed))
		coord.Add(&move)
	}

	fmt.Println(coord.x * coord.y)

}
