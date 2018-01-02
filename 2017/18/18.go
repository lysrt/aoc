package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type code string

const (
	snd code = "snd"
	set      = "set"
	add      = "add"
	mul      = "mul"
	mod      = "mod"
	rcv      = "rcv"
	jgz      = "jgz"
)

type instruction struct {
	code code
	x, y string
}

type runner struct {
	p         int
	queue     []int
	pos       int
	registers map[string]int
	send      chan int
	rcv       chan int
	sentCount int
}

func NewRunner(p int) *runner {
	r := &runner{
		p:         p,
		queue:     make([]int, 0),
		registers: make(map[string]int),
		send:      make(chan int, 1024),
	}
	r.registers["p"] = p
	return r
}

func (r *runner) run(wg *sync.WaitGroup) {
	for {
		offset := 1
		i := instructions[r.pos]
		switch i.code {
		case snd:
			r.send <- get(r.registers, i.x)
			r.sentCount++
		case set:
			r.registers[i.x] = get(r.registers, i.y)
		case add:
			r.registers[i.x] += get(r.registers, i.y)
		case mul:
			r.registers[i.x] *= get(r.registers, i.y)
		case mod:
			r.registers[i.x] %= get(r.registers, i.y)
		case rcv:
			select {
			case v := <-r.rcv:
				r.registers[i.x] = v
			case <-time.After(50 * time.Millisecond):
				wg.Done()
				return
			}
		case jgz:
			if get(r.registers, i.x) > 0 {
				offset = get(r.registers, i.y)
			}
		}
		r.pos += offset
	}
}

var instructions []instruction

func main() {
	instructions = getInstructions()
	partOne()

	r1 := NewRunner(0)
	r2 := NewRunner(1)

	r1.rcv = r2.send
	r2.rcv = r1.send

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go r1.run(wg)
	go r2.run(wg)

	wg.Wait()

	fmt.Println(r2.sentCount)
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
		}
		switch code {
		case snd:
		case rcv:
		default:
			i.y = parts[2]
		}

		is = append(is, i)
	}
	return is
}

func partOne() {
	registers := make(map[string]int)
	played := 0
	done := false
	pos := 0
	for !done {
		i := instructions[pos]
		newPlayed, offset, done := runOne(registers, i, played)
		played = newPlayed
		if done {
			break
		}
		pos += offset
	}
}

func runOne(registers map[string]int, i instruction, played int) (int, int, bool) {
	offset := 1
	terminate := false

	switch i.code {
	case snd:
		val := get(registers, i.x)
		played = val
	case set:
		val := get(registers, i.y)
		registers[i.x] = val
	case add:
		val := get(registers, i.y)
		registers[i.x] += val
	case mul:
		val := get(registers, i.y)
		registers[i.x] *= val
	case mod:
		val := get(registers, i.y)
		registers[i.x] %= val
	case rcv:
		val := get(registers, i.x)
		if val > 0 {
			fmt.Println("Recovered: ", played)
			terminate = true
		}
	case jgz:
		valX := get(registers, i.x)
		if valX > 0 {
			offset = get(registers, i.y)
		}
	default:
		panic("Unknown code")
	}

	return played, offset, terminate
}

func get(registers map[string]int, y string) int {
	val, err := strconv.Atoi(y)
	if err != nil {
		val = registers[y]
	}
	return val
}
