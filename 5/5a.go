package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	// "ioutil"
	"bufio"
	"os"
)

type Coord struct {
	x int64
	y int64
}

type CoordPair struct {
	one *Coord
	two *Coord
}

func NewCoord(x int64, y int64) *Coord {
	coord := new(Coord)
	coord.x = x
	coord.y = y
	return coord
}

func NewCoordPair(one *Coord, two *Coord) *CoordPair {
	coordPair := new(CoordPair)
	if math.Sqrt(math.Pow(float64(one.x), 2)+math.Pow(float64(one.y), 2)) <= math.Sqrt(math.Pow(float64(two.x), 2)+math.Pow(float64(two.y), 2)) {
		coordPair.one = one
		coordPair.two = two
	} else {
		coordPair.one = two
		coordPair.two = one
	}

	return coordPair
}

func CheckOverlap(first *CoordPair, second *CoordPair) bool {

	if first.one.x <= second.one.x && second.one.x <= first.two.x && second.one.y <= first.one.y && first.one.y <= second.two.y {

		return true
	}

	if second.one.x <= first.one.x && first.one.x <= second.two.x && first.one.y <= second.one.y && second.one.y <= first.two.y {

		return true
	}
	return false
}

func CountOverlap(first *CoordPair, second *CoordPair, overlapCounter *int, overlapMap map[int64]map[int64]int64) {

	var overlapAtY int64
	var overlapAtX int64
	if first.one.x == first.two.x {
		if second.one.x == second.two.x {

			if first.one.y <= second.one.y {
				for i := second.one.y; i <= second.two.y; i++ {
					if i <= first.two.y {
						overlapAtY = i
						overlapAtX = second.one.x
						if overlapMap[overlapAtX] == nil {
							overlapMap[overlapAtX] = make(map[int64]int64)
						}
						overlapMap[overlapAtX][overlapAtY] += 1
						if overlapMap[overlapAtX][overlapAtY] == 1 {
							*overlapCounter++
						}
					}
				}
			} else {

				for i := first.one.y; i <= first.two.y; i++ {
					if i <= second.two.y {
						overlapAtY = i
						overlapAtX = first.one.x
						if overlapMap[overlapAtX] == nil {
							overlapMap[overlapAtX] = make(map[int64]int64)
						}
						overlapMap[overlapAtX][overlapAtY] += 1
						if overlapMap[overlapAtX][overlapAtY] == 1 {
							*overlapCounter++
						}
					}
				}

			}

		} else {
			overlapAtY = second.one.y
			overlapAtX = first.one.x
			if overlapMap[overlapAtX] == nil {
				overlapMap[overlapAtX] = make(map[int64]int64)
			}
			overlapMap[overlapAtX][overlapAtY] += 1
			if overlapMap[overlapAtX][overlapAtY] == 1 {
				*overlapCounter++
			}

		}
	} else if first.one.y == first.two.y {
		if second.one.y == second.two.y {

			if first.one.x <= second.one.x {
				for i := second.one.x; i <= second.two.x; i++ {
					if i <= first.two.x {
						overlapAtY = second.one.y
						overlapAtX = i
						if overlapMap[overlapAtX] == nil {
							overlapMap[overlapAtX] = make(map[int64]int64)
						}
						overlapMap[overlapAtX][overlapAtY] += 1
						if overlapMap[overlapAtX][overlapAtY] == 1 {
							*overlapCounter++
						}
					}
				}
			} else {

				for i := first.one.x; i <= first.two.x; i++ {
					if i <= second.two.x {
						overlapAtY = first.one.y
						overlapAtX = i
						if overlapMap[overlapAtX] == nil {
							overlapMap[overlapAtX] = make(map[int64]int64)
						}
						overlapMap[overlapAtX][overlapAtY] += 1
						if overlapMap[overlapAtX][overlapAtY] == 1 {
							*overlapCounter++
						}
					}
				}

			}

		} else {
			overlapAtY = first.one.y
			overlapAtX = second.one.x
			if overlapMap[overlapAtX] == nil {
				overlapMap[overlapAtX] = make(map[int64]int64)
			}
			overlapMap[overlapAtX][overlapAtY] += 1
			if overlapMap[overlapAtX][overlapAtY] == 1 {
				*overlapCounter++
			}
		}
	} else {
		// Is this a case
		fmt.Println("Is this a case")
	}

}

func PrintCoordPair(name string, coordPair CoordPair) {
	fmt.Println(name, coordPair.one.x, coordPair.one.y, "->", coordPair.two.x, coordPair.two.y)
}

func main() {

	//
	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	var totalOverlaps = 0
	var coordPairList []*CoordPair

	var overlapMap = map[int64]map[int64]int64{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		stringSplit := strings.Fields(inputString)
		stringSplitOne := strings.Split(stringSplit[0], ",")
		stringSplitTwo := strings.Split(stringSplit[2], ",")
		xOne, _ := strconv.ParseInt(stringSplitOne[0], 10, 64)
		yOne, _ := strconv.ParseInt(stringSplitOne[1], 10, 64)
		coordOne := NewCoord(xOne, yOne)

		xTwo, _ := strconv.ParseInt(stringSplitTwo[0], 10, 64)
		yTwo, _ := strconv.ParseInt(stringSplitTwo[1], 10, 64)
		coordTwo := NewCoord(xTwo, yTwo)
		coordPair := NewCoordPair(coordOne, coordTwo)

		if coordPair.one.x == coordPair.two.x || coordPair.one.y == coordPair.two.y {

			for i := len(coordPairList) - 1; i >= 0; i-- {
				if CheckOverlap(coordPair, coordPairList[i]) {
					CountOverlap(coordPair, coordPairList[i], &totalOverlaps, overlapMap)

				}
			}
			coordPairList = append(coordPairList, coordPair)
		}

	}

	fmt.Println(totalOverlaps)

}
