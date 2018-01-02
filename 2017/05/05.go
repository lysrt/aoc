package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Part one
	ints := ReadIntLines("input")
	fmt.Println(jump(ints, partOne))

	// Part two
	fmt.Println(jump(ints, partTwo))
}

func ReadIntLines(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var ints []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func jump(lines []int, update func([]int, int)) int {
	ints := make([]int, len(lines))
	copy(ints, lines)

	current := 0
	nbJumps := 0
	for current >= 0 && current < len(ints) {
		step := ints[current]
		update(ints, current)

		current += step
		nbJumps++
	}
	return nbJumps
}

func partOne(ints []int, index int) {
	ints[index]++
}

func partTwo(ints []int, index int) {
	if ints[index] >= 3 {
		ints[index]--
	} else {
		ints[index]++
	}
}
