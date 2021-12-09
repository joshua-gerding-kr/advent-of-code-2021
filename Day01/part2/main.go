package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
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

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var index int
	var windowSize int = 3
	var windowArray []int
	for index < len(lines)-1 {
		var sum int
		for _, line := range lines[index : index+windowSize] {
			input, _ := strconv.Atoi(line)
			sum += input
		}
		windowArray = append(windowArray, sum)
		index++
	}
	fmt.Println(windowArray)

	var firstInput int
	var secondInput int
	var decrease int
	for _, value := range windowArray {
		secondInput = value
		fmt.Println("First: ", firstInput)
		fmt.Println("Second: ", secondInput)
		if firstInput != 0 && secondInput > firstInput {
			fmt.Println("i am lower")
			decrease++
		}
		firstInput = secondInput
	}
	fmt.Println(decrease)
}
