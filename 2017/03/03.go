package main

import (
	"fmt"
	"math"
)

func main() {
	input := 347991
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input int) int {
	i := 1
	for i*i < input {
		i += 2
	}
	square := i / 2
	return (square * 2) - ((i*i - input) % square)
}

type node struct {
	X, Y, Value int
}

func partTwo(input int) int {
	values := []node{node{0, 0, 1}}
	directions := []node{
		node{1, 0, 0},  // Right
		node{0, 1, 0},  // Up
		node{-1, 0, 0}, // Left
		node{0, -1, 0}, // Down
	}

	steps := 0
	directionIndex := 0
	for {
		// Increase the step (segment length) every two segment
		if directionIndex%2 == 0 {
			steps++
		}

		// Compute one segment value
		for i := 0; i < steps; i++ {
			lastValue := values[len(values)-1]
			direction := directions[directionIndex%4]

			nextX := lastValue.X + direction.X
			nextY := lastValue.Y + direction.Y

			neighboursSum := 0
			for _, n := range values {
				// If node n is a neighbour of (x,y)
				if math.Abs(float64(nextX-n.X)) <= 1 && math.Abs(float64(nextY-n.Y)) <= 1 {
					neighboursSum += n.Value
				}
			}

			if neighboursSum > input {
				return neighboursSum
			}

			next := node{nextX, nextY, neighboursSum}
			values = append(values, next)
		}

		// Then change direction
		directionIndex++
	}
}
