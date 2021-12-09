package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 1 == len(2)
// 4 == len(4)
// 7 == len(3)
// 8 == len(7)
// 2,3,5 == len(5)
// 0,6,9 == len(6)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	stringifiedFile := string(fileBytes)
	array := strings.Split(stringifiedFile, "\n")

	var uniqueNumbers int
	for _, line := range array {
		lineArray := strings.Split(line, "|")
		uniqueNumbersMap := make(map[int]string)
		for i, outputSignalSplit := range lineArray {
			charcarray := strings.Split(outputSignalSplit, " ")
			if i == 0 {
				for len(uniqueNumbersMap) != 10 {
					for _, charcset := range charcarray {
						if len(charcset) == 2 {
							uniqueNumbersMap[1] = charcset
						} else if len(charcset) == 3 {
							uniqueNumbersMap[7] = charcset
						} else if len(charcset) == 4 {
							uniqueNumbersMap[4] = charcset
						} else if len(charcset) == 7 {
							uniqueNumbersMap[8] = charcset
						} else if len(charcset) == 6 {
							var res string = charcset
							for _, v := range uniqueNumbersMap[1] {
								res = strings.ReplaceAll(res, string(v), "")
							}
							if len(res) == 5 {
								uniqueNumbersMap[6] = charcset
							} else {
								for _, v := range uniqueNumbersMap[4] {
									res = strings.ReplaceAll(res, string(v), "")
								}
								if len(res) == 2 {
									uniqueNumbersMap[9] = charcset
								} else {
									uniqueNumbersMap[0] = charcset
								}
							}
						} else if len(charcset) == 5 {
							var res string = charcset
							for _, v := range uniqueNumbersMap[1] {
								res = strings.ReplaceAll(res, string(v), "")
							}
							if len(res) == 3 {
								uniqueNumbersMap[3] = charcset
							} else {
								for _, v := range uniqueNumbersMap[4] {
									res = strings.ReplaceAll(res, string(v), "")
								}
								if len(res) == 2 {
									uniqueNumbersMap[5] = charcset
								} else {
									uniqueNumbersMap[2] = charcset
								}
							}
						}
					}
				}
			} else {
				optimizedMap := make(map[string]string)
				for k, v := range uniqueNumbersMap {
					me := SortString(v)
					stringNumber := strconv.Itoa(k)
					optimizedMap[me] = stringNumber
				}
				var numberString string
				charcarray := strings.Split(outputSignalSplit, " ")
				for _, output := range charcarray {
					sortString := SortString(string(output))
					numberString += optimizedMap[sortString]
				}
				convertNumberString, _ := strconv.Atoi(numberString)
				uniqueNumbers += convertNumberString
			}
		}
	}
	fmt.Println(uniqueNumbers)
}
