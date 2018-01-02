package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	lines := ReadStringLines("input")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func ReadStringLines(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var lines [][]string
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		lines = append(lines, parts)
	}
	return lines
}

func partOne(lines [][]string) int {
	validLines := 0
	for _, line := range lines {
		if hasUniqueWords(line) {
			validLines++
		}
	}
	return validLines
}

func hasUniqueWords(line []string) bool {
	m := make(map[string]bool)
	for _, word := range line {
		_, ok := m[word]
		if !ok {
			m[word] = true
		} else {
			return false
		}
	}
	return true
}

func partTwo(lines [][]string) int {
	validLines := 0
	for _, line := range lines {
		if hasNoAnagrams(line) {
			validLines++
		}
	}
	return validLines
}

func hasNoAnagrams(line []string) bool {
	// Sort letters inside each word
	for i, word := range line {
		sortedWord := sortLetters(word)
		line[i] = sortedWord
	}

	return hasUniqueWords(line)
}

type runes []rune

func (r runes) Len() int {
	return len(r)
}

func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func sortLetters(word string) string {
	r := runes(word)
	sort.Sort(r)
	return string(r)
}
