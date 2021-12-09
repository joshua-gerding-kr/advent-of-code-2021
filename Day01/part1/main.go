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
		fmt.Println("Error: %s\n", err)
		return
	}

	var firstInput int
	var secondInput int
	var decrease int
	for _, line := range lines {
		secondInput, _ = strconv.Atoi(line)
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
