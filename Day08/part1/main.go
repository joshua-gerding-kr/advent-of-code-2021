package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	stringifiedFile := string(fileBytes)
	array := strings.Split(stringifiedFile, "\n")

	fmt.Println(len(array))

	var uniqueNumbers int
	for _, line := range array {
		lineArray := strings.Split(line, "|")
		fmt.Println(lineArray)
		for i, outputSignalSplit := range lineArray {
			if i == 1 {
				charcarray := strings.Split(outputSignalSplit, " ")
				for _, charcset := range charcarray {
					if charcset == "|" {
						continue
					} else if len(charcset) == 2 || len(charcset) == 4 || len(charcset) == 3 || len(charcset) == 7 {
						uniqueNumbers++
					}
				}
			}
		}
	}

	fmt.Println(uniqueNumbers)
}
