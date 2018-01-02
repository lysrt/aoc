package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type component struct {
	in, out int
}

func main() {
	components := getComponents("input")
	fmt.Println("Best", sum(strongest(components, 0)))

	components = getComponents("input")
	fmt.Println("Longest", sum(longest(components, 0)))

}

func getComponents(fileName string) []component {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var components []component
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		parts := strings.Split(s.Text(), "/")
		in, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		out, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		c := component{in, out}
		components = append(components, c)
	}
	return components
}

func strongest(components []component, last int) []int {
	highscore := 0
	var high []int

	for i, comp := range components {
		if comp.in == last || comp.out == last {
			var in int
			if comp.in == last {
				in = comp.out
			} else {
				in = comp.in
			}

			nc := make([]component, len(components))
			copy(nc, components)

			x := append([]int{comp.in, comp.out}, strongest(append(nc[:i], nc[i+1:]...), in)...)

			if sum(x) > highscore {
				highscore = sum(x)
				high = x
			}
		}
	}
	return high
}

func longest(components []component, last int) []int {
	highscore := 0
	longestPath := 0
	var high []int

	for i, comp := range components {
		if comp.in == last || comp.out == last {
			var in int
			if comp.in == last {
				in = comp.out
			} else {
				in = comp.in
			}

			nc := make([]component, len(components))
			copy(nc, components)

			x := append([]int{comp.in, comp.out}, longest(append(nc[:i], nc[i+1:]...), in)...)

			if len(x) >= longestPath {
				longestPath = len(x)

				if sum(x) > highscore {
					highscore = sum(x)
					high = x
				}
			}
		}
	}
	return high
}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
