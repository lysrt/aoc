package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rule struct {
	size    int
	in, out string
}

func main() {
	rules := getRules()

	s := NewSquare(".#./..#/###")

	fmt.Println(s)
	// Part one, grow 5 times
	for i := 0; i < 18; i++ {
		fmt.Println(i + 1)
		s = s.grow(rules)
	}
	// fmt.Println("Result:\n", strings.Replace(s.String(), "/", "\n ", -1))

	fmt.Println("Count: ", s.count())
}

func getRules() []rule {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var rules []rule
	for s.Scan() {
		parts := strings.Split(s.Text(), " => ")
		s := getSize(parts[0])
		r := rule{
			size: s,
			in:   parts[0],
			out:  parts[1],
		}
		rules = append(rules, r)
	}

	return rules
}

func getSize(in string) int {
	return strings.Count(in, "/") + 1
}

func (in square) grow(rules []rule) square {
	totalSize := len(in)

	squaresPerLine := 0
	dividedSize := 0
	if totalSize%2 == 0 {
		squaresPerLine = totalSize / 2
		dividedSize = 2
	} else if totalSize%3 == 0 {
		squaresPerLine = totalSize / 3
		dividedSize = 3
	} else {
		panic("Wrong totalSize")
	}

	// fmt.Println("Dividing into", squaresPerLine, "squares of size", dividedSize)
	squares := make([][]square, squaresPerLine)
	for i := 0; i < squaresPerLine; i++ {
		squares[i] = make([]square, squaresPerLine)
		for j := 0; j < squaresPerLine; j++ {
			sq := in.sub(i, j, dividedSize)
			squares[i][j] = sq
		}
	}

	// For all divided squares: put them into []string
	for i := 0; i < squaresPerLine; i++ {
		for j := 0; j < squaresPerLine; j++ {
			// get square:
			square := squares[i][j]
			newSquare := square.match(rules)
			squares[i][j] = newSquare
		}
	}

	result := assemble(squares)
	return result
}

func (s square) sub(i, j, n int) square {
	// 0, 0, 3 on 3x3 returns s

	//   (1,0) ==> take 2,0 3,0 and 2,1 3,1
	//.. ..
	//.. ..
	//
	//.. ..
	//.. ..
	result := emptySquare(n)
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			result[k][l] = s[i*n+k][j*n+l]
		}
	}
	return result
}

func assemble(squares [][]square) square {
	// fmt.Println("Assembling", len(squares)*len(squares), "squares of size", len(squares[0][0]))
	if len(squares) == 1 {
		return squares[0][0]
	}

	//... ...
	//... ...
	//
	//... ...
	//... ...
	n := len(squares) * len(squares[0][0])
	subN := len(squares[0][0])
	result := emptySquare(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = squares[i/subN][j/subN][i%subN][j%subN]
		}
	}
	return result
}

func (s square) match(rules []rule) square {
	for _, r := range rules {
		if s.matchOne(r) {
			return NewSquare(r.out)
		}
	}
	panic("no match")
}

func (s square) matchOne(r rule) bool {
	if r.in == s.String() {
		return true
	} else if s.rotate().String() == r.in {
		return true
	} else if s.rotate().rotate().String() == r.in {
		return true
	} else if s.rotate().rotate().rotate().String() == r.in {
		return true
	} else if s.flip().String() == r.in {
		return true
	} else if s.flip().rotate().String() == r.in {
		return true
	} else if s.flip().rotate().rotate().String() == r.in {
		return true
	} else if s.flip().rotate().rotate().rotate().String() == r.in {
		return true
	} else {
		return false
	}
}

type square [][]byte

func NewSquare(str string) square {
	totalSize := getSize(str)
	s := make(square, totalSize)
	parts := strings.Split(str, "/")
	for i, p := range parts {
		line := make([]byte, totalSize)
		for j, b := range p {
			line[j] = byte(b)
		}
		s[i] = line
	}
	return s
}

func emptySquare(n int) square {
	result := make(square, n)
	for i := range result {
		result[i] = make([]byte, n)
	}
	return result
}

func (s square) String() string {
	var lines []string
	for _, l := range s {
		line := ""
		for _, b := range l {
			line += string(b)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "/")
}

func (s square) count() int {
	count := 0
	for _, b := range s.String() {
		if b == '#' {
			count++
		}
	}
	return count
}

func (s square) rotate() square {
	// #$   ;#   .;   $.
	// ;.   .$   $#   #;
	result := emptySquare(len(s))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			result[len(s)-j-1][i] = s[i][j]
		}
	}
	return result
}

func (s square) flip() square {
	var result square
	for _, l := range s {
		// reverse l
		reversed := make([]byte, len(l))
		for i, b := range l {
			reversed[len(l)-i-1] = b
		}
		result = append(result, reversed)
	}
	return result
}
