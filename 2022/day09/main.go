package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int

const (
	undefined direction = iota
	up
	down
	left
	right
)

func (d direction) String() string {
	return fmt.Sprintf("%v", []string{"undefined", "up", "down", "left", "right"}[d])
}

var directions = map[string]direction{"": undefined, "U": up, "D": down, "L": left, "R": right}

type move struct {
	dir   direction
	count int
}

func (m move) String() string {
	return fmt.Sprintf("{%v %v}", m.dir, m.count)
}

type point struct {
	x int
	y int
}

func equals(p1, p2 point) bool {
	if p1.x == p2.x && p1.y == p2.y {
		return true
	}
	return false
}

type grid struct {
	width      int
	height     int
	points     [][]point
	start      point
	head       point
	tail       point
	tailPoints map[point]struct{}
}

func (g grid) String() string {
	var out string
	// Iterate from the top row down.
	for y := g.height - 1; y >= 0; y-- {
		for _, p := range g.points[y] {
			if (equals(p, g.head) && equals(p, g.tail)) || (equals(p, g.head) && !equals(p, g.tail)) {
				out += "H"
			} else if equals(p, g.tail) {
				out += "T"
			} else if equals(p, g.start) {
				out += "s"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func newGrid(moves []move) grid {
	// Initialize dimensions.
	maxWidth := 0
	maxHeight := 0
	for _, m := range moves {
		if m.dir == left || m.dir == right {
			if m.count > maxWidth {
				maxWidth = m.count
			}
		}
		if m.dir == up || m.dir == down {
			if m.count > maxHeight {
				maxHeight = m.count
			}
		}
	}

	// We assume that the width or height is equal to one more than the longest move vertically or horizontally.
	width := maxWidth + 1
	height := maxHeight + 1

	// Initialize points.
	points := [][]point{}
	for y := 0; y < height; y++ {
		row := []point{}
		for x := 0; x < width; x++ {
			row = append(row, point{x: x, y: y})
		}
		points = append(points, row)
	}

	// Initialize the point (0,0) as the start point.
	start := point{x: 0, y: 0}
	head := point{x: 0, y: 0}
	tail := point{x: 0, y: 0}

	tailPoints := make(map[point]struct{})

	return grid{width: width, height: height, points: points, start: start, head: head, tail: tail, tailPoints: tailPoints}
}

func (g *grid) move(m move) {
	// fmt.Println(m)
	for i := 0; i < m.count; i++ {
		// Move head.
		switch m.dir {
		case right:
			g.head.x++
		case left:
			g.head.x--
		case up:
			g.head.y++
		case down:
			g.head.y--
		}

		// Move tail.
		// Rules for tail moves:
		// * tail must always be adjacent to head
		// * if tail is adjacent to head after the move: tail does not move
		// * if tail is not adjacent to head after the move:
		//   * if head moved right: tail moves right one and to the same y coord
		//   * if head moved left: tail moves left one and to the same y coord
		//   * if head moved up: tail moves up one and to the same x coord
		//   * if head moved down: tail moves down one and to the same x coord
		if adjacent(g.head, g.tail) {
			continue
		}
		switch m.dir {
		case right:
			g.tail.x++
			g.tail.y = g.head.y
		case left:
			g.tail.x--
			g.tail.y = g.head.y
		case up:
			g.tail.y++
			g.tail.x = g.head.x
		case down:
			g.tail.y--
			g.tail.x = g.head.x
		}
		g.tailPoints[g.tail] = struct{}{}
		// fmt.Println(g)
	}
}

func adjacent(p1, p2 point) bool {
	xDiff := p1.x - p2.x
	yDiff := p1.y - p2.y
	if (xDiff >= -1 && xDiff <= 1) && (yDiff >= -1 && yDiff <= 1) {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	moves := []move{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m := move{}
		var dir string
		fmt.Sscanf(scanner.Text(), "%s %d", &dir, &m.count)
		m.dir = directions[dir]
		moves = append(moves, m)
	}

	g := newGrid(moves)

	for _, m := range moves {
		g.move(m)
	}
	fmt.Println("Solution 1:", len(g.tailPoints))
}
