package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id         int
	neighbours []int
}

func main() {
	nodes := getNodes()

	// Part one
	// fmt.Println("Size of groupZero: ", len(getGroup(nodes, nodes[0].id)))

	// Part two
	var groups [][]int
	var visitedNodes []int

	for len(visitedNodes) < len(nodes) {
		nonVisitedID := getFirstNonVisited(visitedNodes, len(nodes))
		group := getGroup(nodes, nonVisitedID)
		visitedNodes = append(visitedNodes, group...)

		groups = append(groups, group)
	}

	fmt.Println("Size of groupZero: ", len(groups[0]))
	fmt.Println("Groups: ", len(groups))
}

func getNodes() map[int]node {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)

	nodes := make(map[int]node)
	for r.Scan() {
		parts := strings.Split(r.Text(), " <-> ")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		var neighbours []int
		parts = strings.Split(parts[1], ", ")
		for _, p := range parts {
			neighbour, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			neighbours = append(neighbours, neighbour)
		}

		node := node{id, neighbours}
		nodes[id] = node
	}
	return nodes
}

func getGroup(nodes map[int]node, rootID int) []int {
	group := []int{rootID}
	visited := make(map[int]bool)

	for len(visited) < len(group) {
		for _, id := range group {
			if _, seen := visited[id]; !seen {
			neighbourLabel:
				for _, neighbourID := range nodes[id].neighbours {
					// Add neighbourID only if it was not already in the list
					for _, discoveredID := range group {
						if discoveredID == neighbourID {
							continue neighbourLabel
						}
					}

					group = append(group, neighbourID)
				}

				visited[id] = true
			}
		}
	}

	return group
}

func getFirstNonVisited(visited []int, total int) int {
idLabel:
	for i := 0; i < total; i++ {
		for _, visitedID := range visited {
			if visitedID == i {
				continue idLabel
			}
		}
		return i
	}
	return -1
}
