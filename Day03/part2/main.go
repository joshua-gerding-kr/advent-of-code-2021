package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type counts struct {
	zeroCount int
	onesCount int
}

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

func looptheloop(mainArray []string, index int, comp string) (array []string) {
	var onesArray []string
	var zeroArray []string

	for _, line := range mainArray {
		if string(line[index]) == "1" {
			onesArray = append(onesArray, line)
		}
		if string(line[index]) == "0" {
			zeroArray = append(zeroArray, line)
		}
	}
	if comp == "greater" && len(onesArray) >= len(zeroArray) {
		return onesArray
	} else if comp == "greater" && len(onesArray) < len(zeroArray) {
		return zeroArray
	} else if comp == "lesser" && len(onesArray) < len(zeroArray) {
		return onesArray
	} else if comp == "lesser" && len(onesArray) >= len(zeroArray) {
		return zeroArray
	} else {
		return zeroArray
	}
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		return
	}

	var theArray []string = lines
	var index int
	for index < 12 {
		theArray = looptheloop(theArray, index, "greater")
		if len(theArray) == 1 {
			break
		}
		index++
	}

	var theArray2 []string = lines
	var index2 int
	for index2 < 12 {
		theArray2 = looptheloop(theArray2, index2, "lesser")
		if len(theArray2) == 1 {
			break
		}
		index2++
	}

	oxygenvalue, _ := strconv.ParseInt(theArray[0], 2, 64)
	co2value, _ := strconv.ParseInt(theArray2[0], 2, 64)
	fmt.Println(oxygenvalue * co2value)
}
