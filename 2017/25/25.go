package main

import "fmt"

type state struct {
	name               string
	zeroValue, zeroDir int
	oneValue, oneDir   int
	zeroNext, oneNext  string
}

func main() {
	states := map[string]state{
		"A": state{"A", 1, +1, 0, -1, "B", "F"},
		"B": state{"B", 0, +1, 0, +1, "C", "D"},
		"C": state{"C", 1, -1, 1, +1, "D", "E"},
		"D": state{"D", 0, -1, 0, -1, "E", "D"},
		"E": state{"E", 0, +1, 1, +1, "A", "C"},
		"F": state{"F", 1, -1, 1, +1, "A", "A"},
	}
	tape := []int{0}
	cursor := 0

	steps := 12994925
	state := "A"

	for i := 0; i < steps; i++ {
		value := tape[cursor]
		s := states[state]
		if value == 0 {
			tape[cursor] = s.zeroValue
			cursor += s.zeroDir
			state = s.zeroNext
		} else if value == 1 {
			tape[cursor] = s.oneValue
			cursor += s.oneDir
			state = s.oneNext
		}

		if cursor < 0 {
			tape = append([]int{0}, tape...)
			cursor++
		}
		if cursor >= len(tape) {
			tape = append(tape, 0)
		}
	}

	sum := 0
	for _, t := range tape {
		sum += t
	}
	fmt.Println("Sum:", sum)
}
