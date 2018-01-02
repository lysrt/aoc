package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name        string
	Weight      int
	Children    []*Node
	TotalWeight int
}

func (node *Node) weigh() int {
	sum := node.Weight
	for _, c := range node.Children {
		sum += c.weigh()
	}
	node.TotalWeight = sum
	return sum
}

func newNode(name string, weight int) *Node {
	n := Node{
		Name:     name,
		Weight:   weight,
		Children: []*Node{},
	}
	return &n
}

func main() {
	nodeTable := createNodeMaps()
	root := getRoot(nodeTable)
	fmt.Println(root.Name)

	// Populate the TotalWeight attribute of each node
	root.weigh()

	fmt.Println(partTwo(root, 0))
}

func createNodeMaps() map[string]*Node {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nodeTable := make(map[string]*Node)
	nodeWithChildren := make(map[string][]string)

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		node := parseLine(line, nodeWithChildren)
		nodeTable[node.Name] = node
	}

	for nodeName, childrenNames := range nodeWithChildren {
		for _, name := range childrenNames {
			nodeTable[nodeName].Children = append(nodeTable[nodeName].Children, nodeTable[name])
		}
	}

	return nodeTable
}

func parseLine(line string, nodeWithChildren map[string][]string) *Node {
	var name string
	var weight int
	parts := strings.Split(line, "->")
	fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)

	node := newNode(strings.TrimSpace(name), weight)

	if len(parts) == 2 {
		var childrenNames []string
		names := strings.Split(parts[1], ",")
		for _, name := range names {
			childrenNames = append(childrenNames, strings.TrimSpace(name))
		}
		nodeWithChildren[node.Name] = childrenNames
	}

	return node
}

func getRoot(nodeTable map[string]*Node) *Node {
	children := make(map[string]bool)
	for _, node := range nodeTable {
		for _, child := range node.Children {
			if _, ok := children[child.Name]; !ok {
				children[child.Name] = true
			}
		}
	}

	var rootName string
	// The root is the only node not in children map
	for nodeName := range nodeTable {
		if _, ok := children[nodeName]; !ok {
			rootName = nodeName
			break
		}
	}

	return nodeTable[rootName]
}

func partTwo(root *Node, expectedWeight int) int {
	// List all children weights and their number of appearance
	childrenWeights := make(map[int]int)
	for _, c := range root.Children {
		w := c.TotalWeight
		if _, ok := childrenWeights[w]; ok {
			childrenWeights[w]++
		} else {
			childrenWeights[w] = 1
		}
	}

	// Find the child with the wrong weight
	weirdWeight, correctWeight := 0, 0
	min := len(root.Children)
	max := 0
	for currentWeight, appearance := range childrenWeights {
		if appearance < min {
			min = appearance
			weirdWeight = currentWeight
		}
		if appearance > max {
			max = appearance
			correctWeight = currentWeight
		}
	}

	var weirdChild *Node
	for _, c := range root.Children {
		if c.TotalWeight == weirdWeight {
			weirdChild = c
			break
		}
	}

	if weirdChild != nil {
		// If node children are not balanced, check the children
		return partTwo(weirdChild, correctWeight)
	} else {
		// Root has only balanced children, he's the culprit
		diff := root.TotalWeight - expectedWeight
		// Send the corrected weight back to the caller
		return root.Weight - diff
	}
}
