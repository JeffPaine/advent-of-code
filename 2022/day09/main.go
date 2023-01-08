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

// touching determines if two points are "touching" per the problem's definition. Diagonally adjacent or even
// overlapping also count as "touching".
func touching(p1, p2 point) bool {
	xDiff := p1.x - p2.x
	yDiff := p1.y - p2.y
	if (xDiff >= -1 && xDiff <= 1) && (yDiff >= -1 && yDiff <= 1) {
		return true
	}
	return false
}

type knot struct {
	point
	parent  *knot
	child   *knot
	visited map[point]struct{}
}

func (k *knot) move(d direction) {
	// Move the knot if it's the head.
	if k.parent == nil {
		switch d {
		case right:
			k.x++
		case left:
			k.x--
		case up:
			k.y++
		case down:
			k.y--
		}
		if k.child != nil {
			k.child.move(d)
		}
		return
	}

	// We don't need to move knots that are already touching.
	if touching(k.point, k.parent.point) {
		if k.child != nil {
			k.child.move(d)
		}
		return
	}

	// The knots are not touching, move them.
	if k.x == k.parent.x || k.y == k.parent.y {
		// Points are in the same row or column.
		// "If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move
		// one step in that direction so it remains close enough"
		k.y += (k.parent.y - k.y) / 2
		k.x += (k.parent.x - k.x) / 2
	} else {
		// Points aren't in the same row or column and one delta is more than 1, move diagonally towards the parent.
		// "Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always
		// moves one step diagonally to keep up"
		if k.parent.x-k.x < 0 {
			k.x--
		} else {
			k.x++
		}
		if k.parent.y-k.y < 0 {
			k.y--
		} else {
			k.y++
		}
	}
	k.visited[k.point] = struct{}{}

	if k.child != nil {
		k.child.move(d)
	}
}

func newKnot() *knot {
	k := &knot{}
	k.visited = map[point]struct{}{}
	k.visited[k.point] = struct{}{}
	return k
}

func newKnots(count int) (*knot, *knot) {
	head := newKnot()
	curr := head
	for i := 0; i < count-1; i++ {
		k := newKnot()
		k.parent = curr
		curr.child = k
		curr = k
	}
	return head, curr
}

func main() {
	// f, err := os.Open("example.txt")
	// f, err := os.Open("example2.txt")
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

	head, tail := newKnots(2)
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			head.move(m.dir)
		}
	}
	fmt.Println("Solution 1:", len(tail.visited))

	head2, tail2 := newKnots(10)
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			head2.move(m.dir)
		}
	}
	fmt.Println("Solution 2:", len(tail2.visited))
}
