package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	groupLevel := 0
	groupSum := 0
	garbageCount := 0
	var garbage bool
	for s.Scan() {
		char := s.Text()
		//fmt.Println(char)
		switch char {
		case "{":
			if !garbage {
				groupLevel++
				groupSum += groupLevel
			} else {
				garbageCount++
			}
		case "}":
			if !garbage {
				groupLevel--
			} else {
				garbageCount++
			}
		case "<":
			if !garbage {
				garbage = true
			} else {
				garbageCount++
			}
		case ">":
			garbage = false
		case "!":
			s.Scan() // Skip next character
		default:
			if garbage {
				garbageCount++
			}
		}
	}
	fmt.Println(groupSum)
	fmt.Println(garbageCount)
}
