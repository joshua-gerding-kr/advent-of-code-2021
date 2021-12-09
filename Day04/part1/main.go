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
	fmt.Println(board)

	for startingIndex < lengthRow {
		lengthBoard := len(board)
		fmt.Println(lengthBoard)
		var newRow []*boardNumber
		for lengthBoard > 0 {
			me := board[lengthBoard-1][startingIndex]
			newRow = append(newRow, me)

			lengthBoard--
		}
		new = append(new, newRow)
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

func isThereAWinner(arrays [][][]*boardNumber) (bool, int) {
	var winningBoard bool
	var boardIndex int
	var completedBoard bool
	for i, board := range arrays {
		completedBoard = checkBoards(board)
		fmt.Println("Completed Board Horizontal: ", completedBoard)
		swappedBoard := swapXY(board)
		completedBoard = checkBoards(swappedBoard)
		if completedBoard {
			winningBoard = true
			boardIndex = i
			break
		}
	}
	return winningBoard, boardIndex
}

func letsPlayBingo(numbers []string, boardsArray [][][]*boardNumber) (int, int) {
	var winningIndex int
	var winningBoard bool
	var winningNumber int
	for _, number := range numbers {
		fmt.Println("Number: ", number)
		for _, board := range boardsArray {
			for _, line := range board {
				for _, entry := range line {
					if entry.n == number {
						entry.marked = true
					}
				}
			}
		}
		winningBoard, winningIndex = isThereAWinner(boardsArray)
		if winningBoard {
			winningNumber, _ = strconv.Atoi(number)
			break
		}
	}
	return winningIndex, winningNumber
}

func unmarkedSummed(board [][]*boardNumber) int {
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

	unmarkedSummed := unmarkedSummed(structuredBoards[winningIndex])

	fmt.Println(unmarkedSummed * winningNumber)

}
