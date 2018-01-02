package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Move interface {
	Execute(dancers []string)
}

type Spin struct {
	rounds int
}

func (s Spin) Execute(dancers []string) {
	for r := 0; r < s.rounds; r++ {
		last := dancers[len(dancers)-1]
		for i := len(dancers) - 1; i > 0; i-- {
			dancers[i] = dancers[i-1]
		}
		dancers[0] = last
	}
}

type Exchange struct {
	pos1, pos2 int
}

func (e Exchange) Execute(dancers []string) {
	dancers[e.pos1], dancers[e.pos2] = dancers[e.pos2], dancers[e.pos1]
}

type Partner struct {
	nameA, nameB string
}

func (p Partner) Execute(dancers []string) {
	pos1 := indexOf(dancers, p.nameA)
	pos2 := indexOf(dancers, p.nameB)

	if pos1 < 0 || pos2 < 0 {
		panic("Index not found")
	}

	dancers[pos1], dancers[pos2] = dancers[pos2], dancers[pos1]
}

func indexOf(s []string, o string) int {
	for i, e := range s {
		if e == o {
			return i
		}
	}
	return -1
}

func main() {
	moves := getMoves()

	var dancers []string
	for i := 'a'; i <= 'p'; i++ {
		dancers = append(dancers, string(i))
	}

	var seen []string

	for d := 0; d < 1000000000; d++ {
		// Looking for cycles
		if in(seen, strings.Join(dancers, "")) {
			fmt.Println(seen[1000000000%d])
			break
		}
		seen = append(seen, strings.Join(dancers, ""))

		for _, m := range moves {
			m.Execute(dancers)
		}
		if d == 0 {
			fmt.Println(strings.Join(dancers, "")) // First dance
		}
	}
}

func getMoves() []Move {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(b), ",")

	var moves []Move
	for _, p := range parts {
		move := p[0]
		switch move {
		case 'x':
			dancers := strings.Split(p[1:], "/")
			i1, err := strconv.Atoi(dancers[0])
			if err != nil {
				panic(err)
			}
			i2, err := strconv.Atoi(dancers[1])
			if err != nil {
				panic(err)
			}
			moves = append(moves, Exchange{i1, i2})
		case 'p':
			dancers := strings.Split(p[1:], "/")
			moves = append(moves, Partner{dancers[0], dancers[1]})
		case 's':
			i, err := strconv.Atoi(p[1:])
			if err != nil {
				panic(err)
			}
			moves = append(moves, Spin{i})
		}
	}
	return moves
}

func in(s []string, o string) bool {
	for _, e := range s {
		if e == o {
			return true
		}
	}
	return false
}
