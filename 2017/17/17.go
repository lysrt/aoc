package main

import "fmt"

func main() {

	// Part one
	buffer, pos := run(2017)
	fmt.Println(buffer[pos+1])

	// Part two
	// Don't use a buffer, just look at when pos is equal to 1
	pos = 0
	valAfterZero := 0
	for round := 1; round <= 50000000; round++ {
		pos = (pos+316)%round + 1 // len(buffer) is replaced by round
		if pos == 1 {
			valAfterZero = round
		}
	}
	fmt.Println(valAfterZero)

}

func run(rounds int) ([]int, int) {
	buffer := []int{0}
	pos := 0
	for round := 1; round <= rounds; round++ {
		new := (pos+316)%len(buffer) + 1

		// Insert the new value after pos
		buffer = append(buffer[:new], append([]int{round}, buffer[new:]...)...)

		pos = new
	}

	return buffer, pos
}
