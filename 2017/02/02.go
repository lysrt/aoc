package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ints := ReadTabSeparatedLinesInts("input")

	fmt.Println(partOne(ints))
	fmt.Println(partTwo(ints))
}

func ReadTabSeparatedLinesInts(fileName string) [][]int {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var ints [][]int
	for s.Scan() {
		line := lineToInts(s.Text(), "\t")
		ints = append(ints, line)
	}

	return ints
}

func lineToInts(line, separator string) []int {
	var ints []int

	parts := strings.Split(line, separator)
	for _, a := range parts {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func partOne(ints [][]int) int {
	sum := 0
	for _, line := range ints {
		min, max := math.MaxUint32, 0
		for _, i := range line {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		diff := max - min
		sum += diff
	}
	return sum
}

func partTwo(ints [][]int) int {
	sum := 0
	for _, line := range ints {
	numberLabel:
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				a, b := line[i], line[j]
				if a%b == 0 {
					sum += a / b
					break numberLabel
				} else if b%a == 0 {
					sum += b / a
					break numberLabel
				}
			}
		}
	}
	return sum
}
