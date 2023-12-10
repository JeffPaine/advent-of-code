package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
	t    tile
	loop bool
}

type tile rune

const (
	// | is a vertical pipe connecting north and south.
	vertical tile = iota
	// - is a horizontal pipe connecting east and west.
	horizontal
	// L is a 90-degree bend connecting north and east.
	northAndEast
	// J is a 90-degree bend connecting north and west.
	northAndWest
	// 7 is a 90-degree bend connecting south and west.
	southAndWest
	// F is a 90-degree bend connecting south and east.
	southAndEast
	// . is ground; there is no pipe in this tile.
	ground
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	start
)

func (t tile) String() string {
	return []string{"vertical", "horizontal", "northAndEast", "northAndWest", "southAndWest", "southAndEast", "ground", "start"}[t]
}

func parseTile(r rune) tile {
	switch r {
	case '|':
		return vertical
	case '-':
		return horizontal
	case 'L':
		return northAndEast
	case 'J':
		return northAndWest
	case '7':
		return southAndWest
	case 'F':
		return southAndEast
	case '.':
		return ground
	case 'S':
		return start
	default:
		panic(fmt.Sprintf("unexpected rune: %v", string(r)))
	}
}

// direction is the direction of motion when moving from tile a to b.
type direction int

const (
	left direction = iota
	right
	up
	down
)

func (d direction) String() string {
	return []string{"left", "right", "up", "down"}[d]
}

// findStart returns the index of the start tile.
func findStart(points []point) int {
	index := 0
	for i, p := range points {
		if p.t == start {
			index = i
			break
		}
	}
	return index
}

func idxFromDir(d direction, curr, width int) int {
	switch d {
	case left:
		return curr - 1
	case right:
		return curr + 1
	case up:
		return curr - width
	case down:
		return curr + width
	default:
		panic(fmt.Sprintf("unsupported direction: %v", d))
	}
}

// walk uses Depth First Search (DFS) to find all points on the loop.
func walk(points []point, width int, curr, last int) {
	// Only called on valid points on the loop.
	points[curr].loop = true

	for _, dir := range []direction{left, right, up, down} {
		idx := idxFromDir(dir, curr, width)
		// Don't search via the node we just came from.
		if idx == last {
			continue
		}
		// Out of bounds.
		if idx < 0 || idx >= len(points) {
			continue
		}
		// We've reached the end of the loop.
		if points[idx].t == start {
			return
		}
		if !isValid(points[curr], points[idx], dir) {
			continue
		}
		walk(points, width, idx, curr)
	}

}

// isValid determines if moving from a to b in direction is valid.
func isValid(a, b point, dir direction) bool {
	if a.t == ground || b.t == ground {
		return false
	}
	switch dir {
	// b < a.
	case left:
		if a.t == vertical || a.t == northAndEast || a.t == southAndEast {
			return false
		}
		if b.t == vertical || b.t == northAndWest || b.t == southAndWest {
			return false
		}
		return true
	// a > b.
	case right:
		if a.t == vertical || a.t == northAndWest || a.t == southAndWest {
			return false
		}
		if b.t == vertical || b.t == northAndEast || b.t == southAndEast {
			return false
		}
		return true
	// b
	// ^
	// a
	case up:
		if a.t == horizontal || a.t == southAndWest || a.t == southAndEast {
			return false
		}
		if b.t == horizontal || b.t == northAndEast || b.t == northAndWest {
			return false
		}
		return true
	// a
	// v
	// b
	case down:
		if a.t == horizontal || a.t == northAndEast || a.t == northAndWest {
			return false
		}
		if b.t == horizontal || b.t == southAndWest || b.t == southAndEast {
			return false
		}
		return true
	default:
		panic(fmt.Sprintf("invalid direction: %v", dir))
	}

}

// enclosed counts the number of points completely enclosed within a loop. It uses a ray tracing
// strategy (https://en.wikipedia.org/wiki/Point_in_polygon).
func enclosed(points []point, width int) int {
	total := 0
	for idx, p := range points {
		// Points on the loop aren't considered.
		if p.loop {
			continue
		}

		// Calculate the index of the end of the current row.
		fromRowStart := idx % width
		fromRowEnd := width - fromRowStart
		endIdx := idx + fromRowEnd - 1

		crosses := 0
		for i := idx; i <= endIdx-1; i++ {
			next := points[i+1]
			if next.loop {
				// Critical problem: what if you're moving along a side of the polygon (example:
				// https://stackoverflow.com/a/63436180)? https://alienryderflex.com/polygon/ had an
				// explainer of the solution:
				//   > Simply follow the rule as described concerning Figure 4. Side c generates a
				//   node, because it has one endpoint below the threshold, and its other endpoint
				//   on-or-above the threshold. Side d does not generate a node, because it has both
				//   endpoints on-or-above the threshold. And side e also does not generate a node,
				//   because it has both endpoints on-or-above the threshold.
				// Solution: when calculating intersections on a horizontal part of the loop, count
				// `|` and only one of either [`F`, `7`] (vertexes with one point below and one
				// point on the line) or [`L`, `J`] (vertexes with one point above and one point on
				// the line).
				if next.t == vertical || next.t == southAndEast || next.t == southAndWest || next.t == start {
					crosses++
				}
			}
		}
		// Odd: enclosed, even: not enclosed.
		if crosses%2 == 1 {
			total++
		}
	}
	return total
}

func parse(r io.Reader) ([]point, int) {
	var points []point
	width := 0
	first := true

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for idx, val := range scanner.Text() {
			if first {
				width = idx + 1
			}
			points = append(points, point{t: parseTile(val)})
		}
		first = false
	}
	return points, width
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	points, width := parse(f)
	start := findStart(points)
	walk(points, width, start, start)

	total := 0
	for _, p := range points {
		if p.loop {
			total++
		}
	}
	// Answer: 6768.
	fmt.Println("Solution 1:", total/2)

	// Answer: 351.
	fmt.Println("Solution 2:", enclosed(points, width))
}
