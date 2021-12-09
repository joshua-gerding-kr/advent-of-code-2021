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

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	// // var depth int
	// var hoizontalLocation int
	m := make(map[int]*counts)
	for _, line := range lines {
		for i, c := range line {
			if string(c) == "1" {
				if m[i] != nil {
					m[i].onesCount++
				} else {
					m[i] = &counts{
						onesCount: 1,
						zeroCount: 0,
					}
				}
			}
			if string(c) == "0" {
				if m[i] != nil {
					m[i].zeroCount++
				} else {
					m[i] = &counts{
						onesCount: 0,
						zeroCount: 1,
					}
				}
			}

		}
	}
	i := 0
	var gamma string
	var epsilon string
	for i < len(m) {
		if m[i].onesCount > m[i].zeroCount {
			gamma += "1"
			epsilon += "0"
		}
		if m[i].onesCount < m[i].zeroCount {
			gamma += "0"
			epsilon += "1"
		}
		i++
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaInt)
	fmt.Println(epsilonInt)
	fmt.Println(gammaInt * epsilonInt)
}
