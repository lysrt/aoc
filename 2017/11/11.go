package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type position struct {
	x, y, z int
}

type offset position

var (
	n  offset = offset{0, 1, -1}
	nw        = offset{-1, 1, 0}
	ne        = offset{1, 0, -1}
	s         = offset{0, -1, 1}
	sw        = offset{-1, 0, 1}
	se        = offset{1, -1, 0}
)

func (p *position) move(o offset) {
	p.x += o.x
	p.y += o.y
	p.z += o.z
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(b), ",")

	pos := position{}
	maxDistance := 0
	for _, d := range parts {
		var dir offset
		switch d {
		case "n":
			dir = n
		case "nw":
			dir = nw
		case "ne":
			dir = ne
		case "s":
			dir = s
		case "sw":
			dir = sw
		case "se":
			dir = se
		}
		// dirs = append(dirs, dir)

		pos.move(dir)
		distance := cubeDistance(position{}, pos)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	// Part 1
	fmt.Println(cubeDistance(position{}, pos))

	// Part 2
	fmt.Println(maxDistance)
}

func cubeDistance(a, b position) int {
	return (abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}
