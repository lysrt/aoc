package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	ints := InputBytesToInts("input")

	fmt.Println(partOne(ints))
	fmt.Println(partTwo(ints))
}

func InputBytesToInts(fileName string) []int {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var ints []int

	for _, s := range input {
		v, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}
		ints = append(ints, v)
	}
	return ints
}

func partOne(ints []int) int {
	sum := 0

	for i := 0; i < len(ints); i++ {
		current := ints[i]
		next := ints[(i+1)%len(ints)]

		if current == next {
			sum += current
		}
	}

	return sum
}

func partTwo(ints []int) int {
	// List has an even number of elements
	steps := len(ints) / 2

	sum := 0

	for i := 0; i < len(ints); i++ {
		current := ints[i]
		next := ints[(i+steps)%len(ints)]

		if current == next {
			sum += current
		}
	}

	return sum
}
