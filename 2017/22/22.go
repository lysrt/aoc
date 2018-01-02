package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type status int

const (
	clean status = iota
	infected
	weak
	flagged
)

type plane map[point]status

type game struct {
	plane
	pos        point
	dir        point
	infections int
}

var up = point{0, 1}

func main() {
	in := getInput("input")
	p := newPlane(in)
	startX := len(in[0]) / 2
	startY := len(in) / 2
	game := game{p, point{startX, startY}, up, 0}

	// partOne(game)

	// Part two
	for i := 0; i < 10000000; i++ {
		game.move()
	}
	fmt.Println(game.infections)
}

func getInput(fileName string) [][]bool {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var result [][]bool
	for s.Scan() {
		var resultLine []bool
		for _, b := range s.Text() {
			resultLine = append(resultLine, b == '#')
		}
		result = append(result, resultLine)
	}
	return result
}

func newPlane(in [][]bool) plane {
	n := len(in)
	plane := make(plane)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			p := point{x, y}
			if in[y][x] {
				plane[p] = infected
			} else {
				plane[p] = clean
			}
		}
	}
	return plane
}

func partOne(game game) {
	for i := 0; i < 10000; i++ {
		game.moveOne()
	}
	fmt.Println(game.infections)
}

func (g *game) moveOne() {

	// Turn, Infect
	s, ok := g.plane[g.pos]
	if ok && s == infected {
		// Infected: Turn right
		g.dir = point{g.dir.y, -g.dir.x}
		g.plane[g.pos] = clean
	} else {
		// Clean: Turn left
		g.dir = point{-g.dir.y, g.dir.x}
		g.plane[g.pos] = infected
		g.infections++
	}

	// Move
	g.pos.x = g.pos.x + g.dir.x
	g.pos.y = g.pos.y - g.dir.y // up is lower
}

func (g *game) move() {

	// Turn, Infect
	s, ok := g.plane[g.pos]
	if ok {
		if s == infected {
			// Infected: Turn right
			g.dir = point{g.dir.y, -g.dir.x}
			g.plane[g.pos] = flagged
		} else if s == weak {
			// Does not turn
			g.plane[g.pos] = infected
			g.infections++
		} else if s == flagged {
			// Reverse direction
			g.dir = point{g.dir.y, -g.dir.x}
			g.dir = point{g.dir.y, -g.dir.x}
			g.plane[g.pos] = clean
		} else {
			// Clean: Turn left
			g.dir = point{-g.dir.y, g.dir.x}
			g.plane[g.pos] = weak
		}
	} else {
		// Clean: Turn left
		g.dir = point{-g.dir.y, g.dir.x}
		g.plane[g.pos] = weak
	}

	// Move
	g.pos.x = g.pos.x + g.dir.x
	g.pos.y = g.pos.y - g.dir.y // up is lower
}
