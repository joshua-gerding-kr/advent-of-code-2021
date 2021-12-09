package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	stringifiedFile := string(fileBytes)
	array := strings.Split(stringifiedFile, "\n")

	fmt.Println(array[0])
	var grid [][]int
	for _, row := range array {
		var gridRow []int
		for _, charac := range row {
			number, _ := strconv.Atoi(string(charac))
			gridRow = append(gridRow, number)
		}
		grid = append(grid, gridRow)
	}

	fmt.Println(grid)
	heightOfGrid := len(grid)
	var riskLevel int
	for heightIndex, row := range grid {
		for corrdIndex, cord := range row {
			low := true
			if corrdIndex != 0 {
				low = cord < row[corrdIndex-1]
				if !low {
					continue
				}
			}
			if corrdIndex != len(row)-1 {
				low = cord < row[corrdIndex+1]
				if !low {
					continue
				}
			}
			if heightIndex != 0 {
				low = cord < grid[heightIndex-1][corrdIndex]
				if !low {
					continue
				}
			}
			if heightIndex != heightOfGrid-1 {
				low = cord < grid[heightIndex+1][corrdIndex]
				if !low {
					continue
				}
			}
			if low {
				riskLevel += cord + 1
			}
		}
	}

	fmt.Println(riskLevel)
}
