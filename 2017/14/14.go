package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "wenycdww"

	var table []string
	for i := 0; i < 128; i++ {
		hash := knotHash(fmt.Sprintf("%s-%d", input, i))
		bin := hexToBin(hash)

		table = append(table, bin)
	}

	var coords []coord
	for y, r := range table {
		for x, c := range r {
			if c == '1' {
				coords = append(coords, coord{x, y})
			}
		}
	}

	// Part one
	fmt.Println(len(coords))

	// Part two
	groups := countGroups(coords)
	fmt.Println(groups)
}

func knotHash(input string) string {
	bytes := []byte(input)
	var ints []int
	for _, b := range bytes {
		ints = append(ints, int(b))
	}

	suffix := []int{17, 31, 73, 47, 23}
	suffixedInts := append(ints, suffix...)

	ring := make([]byte, 256)
	for i := range ring {
		ring[i] = byte(i)
	}

	currentPos := 0
	skipSize := 0

	for i := 0; i < 64; i++ {
		// Run one round
		for _, length := range suffixedInts {
			reverse(ring, currentPos, length)
			// Increase current pos
			currentPos += length + skipSize
			skipSize++
		}
	}

	// ring now holds the sparse hash
	denseHash := densifyHash(ring)
	hex := toHexString(denseHash)

	return hex
}

func reverse(ring []byte, from, size int) {
	for i := from; i < from+size/2; i++ {
		j := 2*from + size - i - 1
		ring[i%len(ring)], ring[j%len(ring)] = ring[j%len(ring)], ring[i%len(ring)]
	}
}

func densifyHash(ring []byte) []byte {
	var result []byte
	for block := 0; block < 16; block++ {
		xor := 0
		for i := 0; i < 16; i++ {
			xor ^= int(ring[block*16+i])
		}
		result = append(result, byte(xor))
	}
	return result
}

func toHexString(hash []byte) string {
	result := ""
	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}
	return result
}

func hexToBin(hex string) string {
	result := ""
	for _, b := range hex {
		i, err := strconv.ParseInt(string(b), 16, 8)
		if err != nil {
			panic(err)
		}
		result += fmt.Sprintf("%04b", i)
	}
	return result
}

type coord struct {
	x, y int
}

func countGroups(coords []coord) int {
	var seen []coord
	groups := 0
	for _, c := range coords {
		if in(seen, c) {
			continue
		}

		groups++
		queue := []coord{c}
		for len(queue) > 0 {
			// pop queue
			currentCoord := queue[0]
			queue = append(queue[:0], queue[1:]...)

			seen = append(seen, currentCoord)
			// neighbours
			ns := []coord{
				coord{currentCoord.x + 1, currentCoord.y},
				coord{currentCoord.x - 1, currentCoord.y},
				coord{currentCoord.x, currentCoord.y + 1},
				coord{currentCoord.x, currentCoord.y - 1},
			}
			for _, n := range ns {
				if n.x >= 0 && n.x < 128 && n.y >= 0 && n.y < 128 && in(coords, n) && !in(seen, n) {
					queue = append(queue, n)
				}
			}
		}
	}

	return groups
}

func in(s []coord, c coord) bool {
	for _, current := range s {
		if c == current {
			return true
		}
	}
	return false
}
