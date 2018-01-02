package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	layers := getLayers()

	// Part one
	sum := 0
	for i, l := range layers {
		// Check when scanner is on the zero position
		if i%(2*(l-1)) == 0 {
			sum += i * l
		}
	}
	fmt.Println(sum)

	// Part two
	delay := 0
	done := false
	for !done {
		done = true
		for i, l := range layers {
			if (i+delay)%(2*(l-1)) == 0 {
				delay++
				done = false
				break
			}
		}
	}
	fmt.Println(delay)
}

func getLayers() map[int]int {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	layers := make(map[int]int)

	for s.Scan() {
		parts := strings.Split(s.Text(), ": ")
		layerDepth, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		layerRange, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		layers[layerDepth] = layerRange
	}
	return layers
}

func getScannerPosition(time, depth int) int {
	offset := time % (2 * (depth - 1))
	if offset > depth-1 {
		return 2*(depth-1) - offset
	} else {
		return offset
	}
}
