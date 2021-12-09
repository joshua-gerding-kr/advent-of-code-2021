package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

type coordinates struct {
	x int
	y int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	var coordinatesArray []string
	for _, line := range lines {
		me := make(map[string]coordinates)
		stringArray := strings.Split(line, " -> ")
		for i, stringCoords := range stringArray {
			coords := strings.Split(stringCoords, ",")
			xCords, _ := strconv.Atoi(coords[0])
			yCords, _ := strconv.Atoi(coords[1])
			structCords := coordinates{x: xCords, y: yCords}
			if i == 0 {
				me["start"] = structCords

			} else {
				me["end"] = structCords
			}
		}
		fmt.Println(me)
		if me["start"].x == me["end"].x {
			startingPoint := me["start"].y
			if me["start"].y >= me["end"].y {
				for startingPoint >= me["end"].y {
					s1 := strconv.FormatInt(int64(startingPoint), 10)
					s2 := strconv.FormatInt(int64(me["end"].x), 10)
					location := s2 + "," + s1
					coordinatesArray = append(coordinatesArray, location)
					startingPoint--
				}
			}
			if me["start"].y <= me["end"].y {
				for startingPoint <= me["end"].y {
					s1 := strconv.FormatInt(int64(startingPoint), 10)
					s2 := strconv.FormatInt(int64(me["end"].x), 10)
					location := s2 + "," + s1
					coordinatesArray = append(coordinatesArray, location)
					startingPoint++
				}
			}
		}
		if me["start"].y == me["end"].y {
			startingPoint := me["start"].x
			if me["start"].x >= me["end"].x {
				for startingPoint >= me["end"].x {
					s1 := strconv.FormatInt(int64(startingPoint), 10)
					s2 := strconv.FormatInt(int64(me["end"].y), 10)
					location := s1 + "," + s2
					coordinatesArray = append(coordinatesArray, location)
					startingPoint--
				}
			}
			if me["start"].x <= me["end"].x {
				for startingPoint <= me["end"].x {
					s1 := strconv.FormatInt(int64(startingPoint), 10)
					s2 := strconv.FormatInt(int64(me["end"].y), 10)
					location := s1 + "," + s2
					coordinatesArray = append(coordinatesArray, location)
					fmt.Println(location)
					startingPoint++
				}
			}
		}
		// Up and left
		if me["start"].y > me["end"].y && me["start"].x > me["end"].x {
			startingPointX := me["start"].x
			startingPointY := me["start"].y
			for startingPointX >= me["end"].x && startingPointY >= me["end"].y {
				s1 := strconv.FormatInt(int64(startingPointX), 10)
				s2 := strconv.FormatInt(int64(startingPointY), 10)
				location := s1 + "," + s2
				coordinatesArray = append(coordinatesArray, location)
				startingPointX--
				startingPointY--
			}
		}
		// down and left
		if me["start"].y < me["end"].y && me["start"].x > me["end"].x {
			startingPointX := me["start"].x
			startingPointY := me["start"].y
			fmt.Println("I am here")
			for startingPointX >= me["end"].x && startingPointY <= me["end"].y {
				s1 := strconv.FormatInt(int64(startingPointX), 10)
				s2 := strconv.FormatInt(int64(startingPointY), 10)
				location := s1 + "," + s2
				coordinatesArray = append(coordinatesArray, location)
				startingPointX--
				startingPointY++
			}
		}
		// Up and right
		if me["start"].y > me["end"].y && me["start"].x < me["end"].x {
			startingPointX := me["start"].x
			startingPointY := me["start"].y
			for startingPointX <= me["end"].x && startingPointY >= me["end"].y {
				s1 := strconv.FormatInt(int64(startingPointX), 10)
				s2 := strconv.FormatInt(int64(startingPointY), 10)
				location := s1 + "," + s2
				coordinatesArray = append(coordinatesArray, location)
				startingPointX++
				startingPointY--
			}
		}
		// down and right
		if me["start"].y < me["end"].y && me["start"].x < me["end"].x {
			startingPointX := me["start"].x
			startingPointY := me["start"].y
			for startingPointX <= me["end"].x && startingPointY <= me["end"].y {
				s1 := strconv.FormatInt(int64(startingPointX), 10)
				s2 := strconv.FormatInt(int64(startingPointY), 10)
				location := s1 + "," + s2
				coordinatesArray = append(coordinatesArray, location)
				startingPointX++
				startingPointY++
			}
		}
	}

	fmt.Println(coordinatesArray)

	coordsMap := make(map[string]int)
	for _, cords := range coordinatesArray {
		if val, ok := coordsMap[cords]; ok {
			fmt.Println(val)
			coordsMap[cords]++
		} else {
			coordsMap[cords] = 1
		}
	}

	var overlap int
	for _, v := range coordsMap {
		if v >= 2 {
			overlap++
		}
	}

	fmt.Println(overlap)
}
