package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeFish(numberArray map[int]int, days *int) map[int]int {
	tempMap := make(map[int]int)
	for i, num := range numberArray {
		if i == 0 {
			tempMap[6] = num + tempMap[6]
			tempMap[8] = num + tempMap[8]
		} else {
			tempMap[i-1] = num + tempMap[i-1]
		}
	}
	return tempMap
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	stringifiedFile := string(fileBytes)
	array := strings.Split(stringifiedFile, ",")

	numberArray := make(map[int]int)
	for _, stringNumber := range array {
		number, _ := strconv.Atoi(stringNumber)
		if _, ok := numberArray[number]; ok {
			numberArray[number]++
		} else {
			numberArray[number] = 1
		}
	}

	var days = 1
	tempMap := numberArray
	for days <= 256 {
		tempMap = makeFish(tempMap, &days)
		days++
	}

	var total int
	for _, v := range tempMap {
		total += v
	}

	fmt.Println(total)
}
