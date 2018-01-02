package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type code string

const (
	set code = "set"
	sub      = "sub"
	mul      = "mul"
	jnz      = "jnz"
)

type instruction struct {
	code code
	x, y string
}

type runner struct {
	pos       int
	registers map[string]int
	muls      int
}

func NewRunner() *runner {
	r := &runner{
		pos:       0,
		registers: make(map[string]int),
	}
	return r
}

func main() {
	instructions := getInstructions()

	// Part one
	r := NewRunner()
	for {
		i := instructions[r.pos]
		offset := r.run(i)
		r.pos += offset

		if r.pos < 0 || r.pos >= len(instructions) {
			break
		}
	}
	fmt.Println(r.muls)

	// Part two would take too long, see asm.go asm()
	/*
		r = NewRunner()
		r.registers["a"] = 1
		for {
			i := instructions[r.pos]
			offset := r.run(i)
			r.pos += offset

			if r.pos < 0 || r.pos >= len(instructions) {
				break
			}
		}
		fmt.Println(r.registers["h"])
	*/
	fmt.Println(asm())
}

func getInstructions() []instruction {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var is []instruction
	for s.Scan() {
		parts := strings.Fields(s.Text())
		code := code(parts[0])
		i := instruction{
			code: code,
			x:    parts[1],
			y:    parts[2],
		}
		is = append(is, i)
	}
	return is
}

func (r *runner) run(i instruction) int {
	offset := 1
	switch i.code {
	case set:
		val := get(r.registers, i.y)
		r.registers[i.x] = val
	case sub:
		val := get(r.registers, i.y)
		r.registers[i.x] -= val
	case mul:
		r.muls++
		val := get(r.registers, i.y)
		r.registers[i.x] *= val
	case jnz:
		valX := get(r.registers, i.x)
		if valX != 0 {
			offset = get(r.registers, i.y)
		}
	default:
		panic("Unknown code")
	}

	return offset
}

func get(registers map[string]int, y string) int {
	val, err := strconv.Atoi(y)
	if err != nil {
		val = registers[y]
	}
	return val
}
