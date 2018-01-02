package main

import (
	"bufio"
	"fmt"
	"os"
)

type block byte

const (
	space block = ' '
	h           = '-'
	v           = '|'
	plus        = '+'
)

type direction byte

const (
	up direction = iota
	down
	left
	right
)

func run(blocks [][]block, startX int) (string, int) {
	result := ""
	steps := 0

	x := startX
	y := 0
	direction := down

	done := false
	for !done {
		steps++
		current := blocks[y][x]
		switch current {
		case h:
			fallthrough
		case v:
			switch direction {
			case up:
				y--
			case down:
				y++
			case left:
				x--
			case right:
				x++
			}
		case plus:
			// Change direction before moving
			if direction == up || direction == down {
				// Look left and right and decide
				if blocks[y][x-1] == space {
					direction = right
					x++
				} else {
					direction = left
					x--
				}
			} else if direction == left || direction == right {
				// Look up and down and decide
				if blocks[y-1][x] == space {
					direction = down
					y++
				} else {
					direction = up
					y--
				}
			}
		default:
			result += string(current)
			switch direction {
			case up:
				if blocks[y-1][x] != space {
					y--
				} else {
					return result, steps
				}
			case down:
				if blocks[y+1][x] != space {
					y++
				} else {
					return result, steps
				}
			case left:
				if blocks[y][x-1] != space {
					x--
				} else {
					return result, steps
				}
			case right:
				if blocks[y][x+1] != space {
					x++
				} else {
					return result, steps
				}
			}
		}
	}
	panic("Wrong exit")
}

func main() {
	blocks := getBlocks()

	startX := 0
	for i, b := range blocks[0] {
		if b == v {
			startX = i
			break
		}
	}

	out, steps := run(blocks, startX)
	fmt.Println(out)
	fmt.Println(steps)
}

func getBlocks() [][]block {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	toBlockLine := func(bytes []byte) []block {
		line := make([]block, len(bytes))
		for i, b := range bytes {
			line[i] = block(b)
		}
		return line
	}

	var blocks [][]block
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		blocks = append(blocks, toBlockLine(s.Bytes()))
	}
	return blocks
}
