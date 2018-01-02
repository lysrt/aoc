package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ints := ReadTabSeparatedInts("input")
	fmt.Println(partOne(ints))
	fmt.Println(partOne(ints)) // Trick to get part 2
}

func ReadTabSeparatedInts(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

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

func partOne(ints []int) int {
	var pastStates [][]int
	nbDistributions := 0
	for {
		if contains(pastStates, ints) {
			break
		}

		past := make([]int, len(ints))
		copy(past, ints)
		pastStates = append(pastStates, past)

		i, v := getMax(ints)
		ints[i] = 0
		distribute(ints, i+1, v)
		nbDistributions++
	}
	return nbDistributions
}

func getMax(ints []int) (index, value int) {
	maxI, max := 0, 0
	for i, v := range ints {
		if v > max {
			maxI, max = i, v
		}
	}
	return maxI, max
}

func distribute(ints []int, startIndex, amount int) {
	for i := startIndex; amount > 0; i++ {
		ints[i%len(ints)]++
		amount--
	}
}

func contains(states [][]int, new []int) bool {
past:
	for _, pastState := range states {
		for i, value := range new {
			if pastState[i] != value {
				continue past
			}
		}
		return true
	}
	return false
}
