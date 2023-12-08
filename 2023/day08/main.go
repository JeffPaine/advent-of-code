package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type node struct {
	name  string
	left  *node
	right *node
}

func parse(r io.Reader) (string, map[string]*node) {
	var strs []string
	var steps string
	m := make(map[string]*node)
	first := true
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if first {
			steps = scanner.Text()
			first = false
			continue
		}
		if scanner.Text() == "" {
			continue
		}
		strs = append(strs, scanner.Text())
	}

	// Make node objects.
	// Example line: AAA = (BBB, BBB)
	for _, s := range strs {
		var n node
		fields := strings.Fields(s)
		n.name = fields[0]
		m[n.name] = &n
	}

	// Link nodes together.
	for _, s := range strs {
		fields := strings.Fields(s)
		n := m[fields[0]]
		n.left = m[strings.TrimRight(strings.TrimPrefix(fields[2], "("), ",")]
		n.right = m[strings.TrimRight(fields[3], ")")]
	}

	return steps, m
}

// partOne solves part one via brute force.
func partOne(steps string, m map[string]*node) int {
	total := 0
	curr := m["AAA"]
loop:
	for {
		for _, step := range steps {
			total++
			if step == 'R' {
				curr = curr.right
			} else {
				curr = curr.left
			}
			if curr.name == "ZZZ" {
				break loop
			}
		}
	}
	return total
}

// PartTwo solves part two via the Least Common Multiple.
func partTwo(steps string, m map[string]*node) int {
	// Find the start nodes.
	var nodes []*node
	for _, n := range m {
		if strings.HasSuffix(n.name, "A") {
			nodes = append(nodes, n)
		}
	}
	var counts []int
	for _, n := range nodes {
		total := 0
	loop:
		for {
			for _, step := range steps {
				total++
				if step == 'R' {
					n = n.right
				} else {
					n = n.left
				}
				if strings.HasSuffix(n.name, "Z") {
					counts = append(counts, total)
					break loop
				}
			}
		}
	}
	return lcm(counts[0], counts[1], counts[2:]...)
}

// gcd finds the Greatest Common Divisor (GCD) between two integers. The GCD is the largest number
// that can divide both integers without a remainder.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the Least Common Multiple (LCM) for a group of integers. The LCM of a group of
// numbers is the smallest number that is a multiple of all of them.
func lcm(a, b int, ints ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(ints); i++ {
		result = lcm(result, ints[i])
	}

	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	steps, m := parse(f)

	fmt.Println("Solution 1:", partOne(steps, m))
	fmt.Println("Solution 2:", partTwo(steps, m))
}
