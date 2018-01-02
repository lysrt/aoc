package main

import (
	"fmt"
)

// From instructions
var multA = 16807
var multB = 48271

func main() {
	/*
		Input:
		Generator A starts with 512
		Generator B starts with 191
	*/
	seedA := 512 //512
	seedB := 191 //191

	// Part One
	var count int
	for i := 0; i < 40000000; i++ {
		seedA = round(seedA, multA)
		seedB = round(seedB, multB)
		if match(seedA, seedB) {
			count++
		}
	}
	fmt.Println(count)

	// Part Two
	var queueA []int
	var queueB []int

	seedA = 512 //512
	seedB = 191 //191
	count = 0
	for {
		if len(queueA) < 5000000 {
			seedA = round(seedA, multA)
			if seedA%4 == 0 {
				queueA = append(queueA, seedA)
			}
		}

		if len(queueB) < 5000000 {
			seedB = round(seedB, multB)
			if seedB%8 == 0 {
				queueB = append(queueB, seedB)
			}
		}

		if len(queueA) >= 5000000 && len(queueB) >= 5000000 {
			break
		}
	}

	la := len(queueA)
	lb := len(queueB)
	min := la
	if lb < min {
		min = lb
	}
	for i := 0; i < min; i++ {
		a := queueA[i]
		b := queueB[i]

		if match(a, b) {
			count++
		}

	}
	fmt.Println(count)
}

func round(seed, mult int) int {
	return seed * mult % 2147483647
}

func match(a, b int) bool {
	// Get the 16 trailing bits. 0xffff = 1 << 16 -1
	return a&0xffff == b&0xffff
}
