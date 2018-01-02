package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes := getBytesInput()

	var bytesToInts []int
	for _, b := range bytes {
		bytesToInts = append(bytesToInts, int(b))
	}

	suffix := []int{17, 31, 73, 47, 23}
	input := append(bytesToInts, suffix...)

	// Part one
	// lengths := getLengths(bytes)

	ring := make([]byte, 256)
	for i := range ring {
		ring[i] = byte(i)
	}

	currentPos := 0
	skipSize := 0
	for i := 0; i < 64; i++ {
		// Run one round
		for _, length := range input {
			reverse(ring, currentPos, length)
			// Increase current pos
			currentPos += length + skipSize
			skipSize++
		}
	}

	// ring now holds the sparse hash
	denseHash := densifyHash(ring)
	fmt.Println(denseHash)

	hex := toHexString(denseHash)
	fmt.Println(hex)

	// Result of part one, checksum after one round (40132)
	// checksum := int(ring[0]) * int(ring[1])
	// fmt.Println("Checksum: ", checksum)
}

func getBytesInput() []byte {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return bytes
}

func getLengths(bytes []byte) []int {
	var lengths []int
	for _, str := range strings.Split(string(bytes), ",") {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		lengths = append(lengths, i)
	}
	return lengths
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
