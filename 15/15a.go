package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coord struct {
	x int64
	y int64
}

func NewCoord(x int64, y int64) Coord {
	coord := new(Coord)
	coord.x = x
	coord.y = y
	return *coord
}

type CoordList struct {
	len       int64
	memberMap map[Coord]bool
	slice     []Coord
}

func NewCoordList() CoordList {
	coordList := new(CoordList)
	coordList.memberMap = make(map[Coord]bool)
	coordList.len = 0
	return *coordList
}

func (coordList *CoordList) InsertSorted(coord Coord, fmap map[Coord]int64) CoordList {

	if coordList.len == 0 {
		coordList.slice = append(coordList.slice, coord)
		coordList.len += 1
	} else {
		index := int64(0)
		for true {
			if index >= coordList.len {
				break
			}
			currentElement := coordList.slice[index]
			if fmap[currentElement] >= fmap[coord] {
				break
			}
			index++
		}

		secondSlice := make([]Coord, coordList.len-index)
		numberOfElements := copy(secondSlice, coordList.slice[index:])
		numberOfElements = numberOfElements
		coordList.slice = append(coordList.slice[:index], coord)
		coordList.slice = append(coordList.slice, secondSlice...)
		coordList.len += 1

	}
	coordList.memberMap[coord] = true
	return *coordList
}

func (coordList *CoordList) Pop() Coord {
	elementToReturn := coordList.slice[0]
	coordList.slice = coordList.slice[1:]
	coordList.memberMap[elementToReturn] = false
	coordList.len -= 1
	return elementToReturn

}

func GetH(coord Coord, endX int64, endY int64) int64 {
	returnValue := int64(0)
	returnValue += int64(math.Abs(float64(endY - coord.y)))
	returnValue += int64(math.Abs(float64(endX - coord.x)))
	return returnValue
}

func main() {

	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	caveMap := map[int64]map[int64]int64{}
	fMap := map[Coord]int64{}
	gMap := map[Coord]int64{}

	openList := NewCoordList()
	scanner := bufio.NewScanner(inputFile)

	rowCounter := int64(0)
	exitX := int64(0)
	exitY := int64(0)
	for scanner.Scan() {

		inputString := scanner.Text()

		for columnCounter, numberChar := range inputString {
			number, _ := strconv.ParseInt(string(numberChar), 10, 64)
			if caveMap[int64(columnCounter)] == nil {
				caveMap[int64(columnCounter)] = make(map[int64]int64)

			}
			caveMap[int64(columnCounter)][rowCounter] = number
			exitX = int64(columnCounter)
		}

		exitY = rowCounter
		rowCounter++

	}

	startCoord := NewCoord(0, 0)
	gMap[startCoord] = 0
	fMap[startCoord] = 0
	openList.InsertSorted(startCoord, fMap)

	for true {

		currentNode := openList.Pop()
		if currentNode.x == exitX && currentNode.y == exitY {
			break
		}

		for addX := int64(-1); addX <= 1; addX++ {
			for addY := int64(-1); addY <= 1; addY++ {

				if (addX == 0 || addY == 0) && addX != addY && currentNode.x+addX >= 0 && currentNode.x+addX <= exitX && currentNode.y+addY >= 0 && currentNode.y+addY <= exitY {
					neighbourNode := NewCoord(currentNode.x+addX, currentNode.y+addY)
					newG := gMap[currentNode] + caveMap[neighbourNode.x][neighbourNode.y]

					if newG < gMap[neighbourNode] || gMap[neighbourNode] == 0 {
						gMap[neighbourNode] = newG
						fMap[neighbourNode] = newG + GetH(neighbourNode, exitX, exitY)
						if openList.memberMap[neighbourNode] != true {
							openList.InsertSorted(neighbourNode, fMap)
						}

					}
				}
			}
		}

	}
	fmt.Println(gMap[NewCoord(exitX, exitY)])

}
