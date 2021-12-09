package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func process(numberArray []int, outerRing int, medianValue int) int {
	var highAmount int
	var lowerAmount int
	var currentBest int
	for _, number := range numberArray {
		if outerRing == 0 {
			if medianValue >= number {
				highAmount += (medianValue - number)
			}
			if medianValue <= number {
				highAmount += (number - medianValue)
			}
		} else {
			highRing := medianValue + outerRing
			lowRing := medianValue - outerRing
			if highRing >= number {
				highAmount += (highRing - number)
			}
			if highRing < number {
				highAmount += (number - highRing)
			}
			if lowRing > number {
				lowerAmount += (lowRing - number)
			}
			if lowRing <= number {
				lowerAmount += (number - lowRing)
			}
		}
	}
	if highAmount <= lowerAmount {
		currentBest = highAmount
	}
	if lowerAmount <= highAmount {
		currentBest = highAmount
	}

	return currentBest
}

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

	sort.Ints(numberArray)
	median := len(numberArray) / 2
	medianValue := numberArray[median]
	outerRing := 0
	currentBest := 0
	for range numberArray {
		localBest := process(numberArray, outerRing, medianValue)
		if localBest <= currentBest || currentBest == 0 {
			currentBest = localBest
		}
		outerRing++
	}
	fmt.Println(currentBest)
}
