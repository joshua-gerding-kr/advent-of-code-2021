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
		fmt.Println("Error: %s\n", err)
		return
	}

	stringifiedFile := string(fileBytes)
	array := strings.Split(stringifiedFile, ",")

	var numberArray []int
	for _, stringNumber := range array {
		number, _ := strconv.Atoi(stringNumber)
		numberArray = append(numberArray, number)
	}

	var days = 1
	for days <= 80 {
		for i, num := range numberArray {
			if num == 0 {
				numberArray[i] = 6
				numberArray = append(numberArray, 8)
			} else {
				numberArray[i]--
			}
		}
		days++
	}
	fmt.Println(len(numberArray))
}
