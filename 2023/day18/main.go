package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	right direction = iota
	left
	up
	down
)

func (d direction) String() string {
	return []string{"right", "left", "up", "down"}[d]
}

type step struct {
	dir      direction
	quantity int
}

func parseStep(s string) step {
	// Example: R 6 (#70c710)
	fields := strings.Split(s, " ")
	var in step
	switch fields[0] {
	case "R":
		in.dir = right
	case "L":
		in.dir = left
	case "U":
		in.dir = up
	case "D":
		in.dir = down
	}
	quantity, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(fmt.Sprintf("could not parse %q to integer: %v", fields[1], err))
	}
	in.quantity = quantity
	return in
}

func parseHex(s string) step {
	// Example: #70c710 = R 461937
	// Each hexadecimal code is six hexadecimal digits long. The first five
	// hexadecimal digits encode the distance in meters as a five-digit
	// hexadecimal number.
	//
	// The last hexadecimal digit encodes the direction to dig: 0 means R, 1
	// means D, 2 means L, and 3 means U.
	fields := strings.Split(s, " ")
	hex := strings.TrimPrefix(fields[2], "(#")
	hex = strings.TrimSuffix(hex, ")")

	i, err := strconv.ParseInt(hex[:5], 16, 64)
	if err != nil {
		panic(err)
	}
	st := step{quantity: int(i)}
	switch hex[5] {
	case '0':
		st.dir = right
	case '1':
		st.dir = down
	case '2':
		st.dir = left
	case '3':
		st.dir = up
	default:
		panic("unexpected direction")
	}
	return st
}

type point struct {
	x int
	y int
}

func makePoints(steps []step) []point {
	var points []point

	x := 0
	y := 0
	points = append(points, point{x: x, y: y})
	for idx, st := range steps {
		switch st.dir {
		case right:
			x += st.quantity
		case left:
			x -= st.quantity
		case up:
			y += st.quantity
		case down:
			y -= st.quantity
		}
		if idx == len(steps)-1 {
			continue
		}
		points = append(points, point{x: x, y: y})
	}

	return points
}

// calculateArea calculates the area of a polygon given its points using
// https://en.wikipedia.org/wiki/Shoelace_formula.
func calculateArea(points []point) int {
	var total float64
	prev := len(points) - 1
	for i := 0; i < len(points); i++ {
		length := float64((points[prev].x + points[i].x) * (points[prev].y - points[i].y))
		total += length
		prev = i
	}
	return int(math.Abs(total / 2.0))
}

func calculateLength(steps []step) int {
	total := 0
	for _, s := range steps {
		total += s.quantity
	}
	return total
}

func calculateVolume(steps []step, points []point) int {
	// Using https://en.wikipedia.org/wiki/Pick's_theorem.
	return calculateArea(points) + 1 + calculateLength(steps)/2
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var steps1 []step
	for _, l := range lines {
		steps1 = append(steps1, parseStep(l))
	}
	points1 := makePoints(steps1)
	// Answer: 46359.
	fmt.Println("Solution 1:", calculateVolume(steps1, points1))

	var steps2 []step
	for _, l := range lines {
		steps2 = append(steps2, parseHex(l))
	}
	points := makePoints(steps2)
	// Answer: 59574883048274.
	fmt.Println("Solution 2:", calculateVolume(steps2, points))
}
