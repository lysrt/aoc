package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	x, y, z int
}

func (v vec) add(a vec) vec {
	return vec{
		v.x + a.x,
		v.y + a.y,
		v.z + a.z,
	}
}

func (v vec) collide(a vec) bool {
	return v.x == a.x && v.y == a.y && v.z == a.z
}

func (v vec) dist() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

type particle struct {
	position, velocity, acceleration vec
}

func main() {
	// partOne()

	// Part two
	particles := getParticles()

	for i := 0; i < 50; i++ {
		for _, p := range particles {
			if p != nil {
				p.velocity = p.velocity.add(p.acceleration)
				p.position = p.position.add(p.velocity)
			}
		}
		deleteCollisions(particles)
	}

	valid := 0
	for _, p := range particles {
		if p != nil {
			valid++
		}
	}
	fmt.Println(valid)
}

func deleteCollisions(particles []*particle) {
	collisions := make(map[int]bool)
	for i := 0; i < len(particles); i++ {
		current := particles[i]
		if current == nil {
			continue
		}
		for j := i + 1; j < len(particles); j++ {
			other := particles[j]
			if other == nil {
				continue
			}
			if current.position.collide(other.position) {
				if _, ok := collisions[i]; !ok {
					collisions[i] = true
				}
				if _, ok := collisions[j]; !ok {
					collisions[j] = true
				}
			}
		}
	}
	removeIndexes(particles, collisions)
}

func removeIndexes(particles []*particle, indexes map[int]bool) {
	for key, _ := range indexes {
		particles[key] = nil
	}
}

func partOne() {
	particles := getParticles()

	// Part one
	for i := 0; i < 500; i++ {
		for _, p := range particles {
			p.velocity = p.velocity.add(p.acceleration)
			p.position = p.position.add(p.velocity)
		}
	}
	minDist := math.MaxInt64
	minIndex := -1
	for i, p := range particles {
		if p.position.dist() < minDist {
			minIndex = i
			minDist = p.position.dist()
		}
	}
	fmt.Println(minIndex)
}

func getParticles() []*particle {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var particles []*particle
	for s.Scan() {
		parts := strings.Split(s.Text(), ", ")
		p := particle{
			parseVec(parts[0]),
			parseVec(parts[1]),
			parseVec(parts[2]),
		}

		particles = append(particles, &p)
	}

	return particles
}

func parseVec(s string) vec {
	parts := strings.Split(s[3:len(s)-1], ",")
	components := [3]int{}
	for i, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		components[i] = val
	}

	return vec{
		components[0],
		components[1],
		components[2],
	}
}
