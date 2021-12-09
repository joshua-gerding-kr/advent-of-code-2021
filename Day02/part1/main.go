package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
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

	var depth int
	var hoizontalLocation int
	for _, line := range lines {
		r := regexp.MustCompile("[^\\s]+")
		array := r.FindAllString(line, -1)
		direction := array[0]
		number, _ := strconv.Atoi(array[1])

		if direction == "up" {
			depth -= number
		} else if direction == "down" {
			depth += number
		} else if direction == "forward" {
			hoizontalLocation += number
		} else if direction == "backward" {
			hoizontalLocation -= number
		}
	}
	fmt.Println(depth * hoizontalLocation)
}
