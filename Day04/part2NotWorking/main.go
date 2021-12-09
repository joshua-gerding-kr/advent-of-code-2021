package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

type boardNumber struct {
	n      string
	marked bool
}

func createBoardStruct(fileBytes []byte) ([]string, [][][]*boardNumber) {
	var structuredBoards [][][]*boardNumber
	stringifiedFile := string(fileBytes)
	destructuredFile := strings.Split(stringifiedFile, "\n\n")
	bingoNumbers := strings.Split(destructuredFile[0], ",")
	for _, board := range destructuredFile[1:] {
		var boardArray [][]*boardNumber
		for _, line := range strings.Split(board, "\n") {
			var lineArray []*boardNumber
			for _, number := range strings.Split(line, " ") {
				if len(number) == 0 {
					continue
				}
				boardNumber := &boardNumber{n: number, marked: false}
				lineArray = append(lineArray, boardNumber)
			}
			boardArray = append(boardArray, lineArray)
		}
		structuredBoards = append(structuredBoards, boardArray)
	}

	return bingoNumbers, structuredBoards
}

func swapXY(board [][]*boardNumber) (swappedBoard [][]*boardNumber) {
	lengthRow := 5
	startingIndex := 0
	var new [][]*boardNumber
	var test [][]boardNumber

	for startingIndex < lengthRow {
		lengthBoard := len(board)
		var newRow []*boardNumber
		var newTest []boardNumber
		for lengthBoard > 0 {
			me := board[lengthBoard-1][startingIndex]
			me2 := board[lengthBoard-1][startingIndex]
			newRow = append(newRow, me)
			newTest = append(newTest, *me2)

			lengthBoard--
		}
		new = append(new, newRow)
		test = append(test, newTest)
		startingIndex++
	}
	return board
}

func checkBoards(board [][]*boardNumber) bool {
	var completedBoard bool
	for _, line := range board {
		var completedNumbers int
		for _, entry := range line {
			if entry.marked {
				completedNumbers++
			}
		}
		if completedNumbers == 5 {
			completedBoard = true
		}
	}
	return completedBoard
}

func isThereAWinner(board [][]*boardNumber) (bool, [][]boardNumber) {
	var winningBoard bool
	var boardInstance [][]boardNumber
	var completedBoard bool
	completedBoard = checkBoards(board)
	if completedBoard {
		winningBoard = true
		for _, line := range board {
			var lines []boardNumber
			for _, entry := range line {
				lines = append(lines, *entry)
			}
			boardInstance = append(boardInstance, lines)
		}
		return winningBoard, boardInstance

	}
	swappedBoard := swapXY(board)
	completedBoard = checkBoards(swappedBoard)
	if completedBoard {
		winningBoard = true

		for _, line := range board {
			var lines []boardNumber
			for _, entry := range line {
				lines = append(lines, *entry)
			}
			boardInstance = append(boardInstance, lines)
		}
	}
	return winningBoard, boardInstance
}

func letsPlayBingo(numbers []string, boardsArray [][][]*boardNumber) ([][]boardNumber, int) {
	var boardWon [][]boardNumber
	var winningNumber int
	for _, number := range numbers {
		for _, board := range boardsArray {
			var winningBoard bool
			for _, line := range board {
				for _, entry := range line {
					if entry.n == number {
						entry.marked = true
					}
				}
			}
			winningBoard, winningIndex := isThereAWinner(board)
			if winningBoard {
				intNumber, _ := strconv.Atoi(number)
				winningNumber = intNumber
				boardWon = winningIndex
			}
		}
	}
	return boardWon, winningNumber
}

func unmarkedSummed(board [][]boardNumber) int {
	var sum int
	for _, row := range board {
		for _, entry := range row {
			if !entry.marked {
				number, _ := strconv.Atoi(entry.n)
				sum += number
			}
		}
	}
	return sum
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	bingoNumbers, structuredBoards := createBoardStruct(fileBytes)

	winningIndex, winningNumber := letsPlayBingo(bingoNumbers, structuredBoards)
	fmt.Println(winningIndex)
	fmt.Println(winningNumber)

	unmarkedSummed := unmarkedSummed(winningIndex)

	fmt.Println(unmarkedSummed * winningNumber)

}
