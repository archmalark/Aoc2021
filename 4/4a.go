package main

import (
	"fmt"
	"strconv"
	"strings"

	// "ioutil"
	"bufio"
	"os"
)

const boardXSize = 5
const boardYSize = 5

type Coord struct {
	x int
	y int
}

type Board struct {
	board       [boardXSize][boardYSize]int
	numberMap   map[int64]Coord
	unmarkedSum int64
	bingo       bool
}

func NewCoord(x int, y int) *Coord {
	coord := new(Coord)
	coord.x = x
	coord.y = y
	return coord
}

func NewBoard() *Board {
	board := new(Board)
	board.numberMap = make(map[int64]Coord)
	board.unmarkedSum = 0
	board.bingo = false
	return board
}

func (board *Board) CheckBingo(number int64) Board {

	coord, isMapContainsKey := board.numberMap[number]

	if board.bingo == false && isMapContainsKey == true {
		bingo := true
		for i := 0; i < boardYSize; i++ {
			if board.board[coord.x][i] == 0 {
				bingo = false
				break
			}
		}

		if bingo == false {
			bingo = true
			for i := 0; i < boardXSize; i++ {
				if board.board[i][coord.y] == 0 {
					bingo = false
					break
				}
			}
		}
		board.bingo = bingo
	}

	return *board
}

func (board *Board) AddNumber(number int64, coord *Coord) Board {
	board.numberMap[number] = *coord
	board.unmarkedSum += number
	return *board
}

func (board *Board) Mark(number int64) Board {

	coord, isMapContainsKey := board.numberMap[number]
	if isMapContainsKey == true && board.board[coord.x][coord.y] != 1 {
		board.board[coord.x][coord.y] = 1
		board.unmarkedSum -= number
		board.CheckBingo(number)
	}

	return *board
}

func main() {

	var boards []*Board
	var numberSequence string

	//
	// Parse input
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	var boardCounter = -1
	var rowCounter = 0
	var currentBoard *Board = nil
	var counter = 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		if counter == 0 {
			numberSequence = inputString
		}

		if inputString == "" {
			currentBoard = NewBoard()
			boardCounter++
		} else if boardCounter >= 0 {
			// rowNumbers := strings.Split(inputString, " ")
			rowNumbers := strings.Fields(inputString)

			for index, number := range rowNumbers {
				numberInt, _ := strconv.ParseInt(number, 10, 64)
				currentBoard.AddNumber(numberInt, NewCoord(index, rowCounter))
			}
			rowCounter++

			if rowCounter >= boardYSize {
				if currentBoard != nil {
					boards = append(boards, currentBoard)
				}

				rowCounter = 0
			}
		}

		counter++

	}

	//
	// Solve puzzle
	var bingo = false
	for _, number := range strings.Split(numberSequence, ",") {
		numberInt, _ := strconv.ParseInt(number, 10, 64)
		for _, board := range boards {

			// board.board[coord.x][coord.y] = 1
			board.Mark(numberInt)

			if board.bingo == true {
				bingo = true

				fmt.Println(board.unmarkedSum * numberInt)
				break
			}

		}
		if bingo == true {
			break
		}
	}

}
