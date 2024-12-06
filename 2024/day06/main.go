package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (d direction) String() string {
	return []string{"left", "right", "up", "down"}[d]
}

type location struct {
	x int
	y int
}

func (l location) String() string {
	return fmt.Sprintf("{x: %v, y: %v}", l.x, l.y)
}

type position struct {
	location
	dir direction
}

func (p position) next() location {
	loc := location{}
	switch p.dir {
	case up:
		loc = location{p.x, p.y - 1}
	case down:
		loc = location{p.x, p.y + 1}
	case left:
		loc = location{p.x - 1, p.y}
	case right:
		loc = location{p.x + 1, p.y}
	}
	return loc
}

func (p *position) turn() {
	switch p.dir {
	case up:
		p.dir = right
	case down:
		p.dir = left
	case left:
		p.dir = up
	case right:
		p.dir = down
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := [][]rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := []rune{}
		for _, b := range scanner.Text() {
			row = append(row, b)
		}
		rows = append(rows, row)
	}

	start := position{}

	// Find the guard's start position.
	for rowIdx, row := range rows {
		for colIdx, col := range row {
			if col == '^' {
				start = position{location: location{x: colIdx, y: rowIdx}, dir: up}
			}
		}
	}

	p := position{location: location{x: start.x, y: start.y}, dir: start.dir}
	visited := make(map[location]int)
	for {
		visited[location{p.x, p.y}] += 1
		next := p.next()
		if next.x < 0 || next.x >= len(rows[0]) || next.y < 0 || next.y >= len(rows) {
			// Next step is off the map.
			break
		}
		if rows[next.y][next.x] != '#' {
			// Step forward into next.
			p.x, p.y = next.x, next.y
			continue
		}
		// We've hit a "#", turn right
		p.turn()
	}

	// Answer: 5461.
	fmt.Println("Solution 1:", len(visited))

	loops := 0
	for rowIdx, row := range rows {
		for colIdx, col := range row {
			// Can't place the obstruction where the guard or an obstruction already is.
			if col == '^' || col == '#' {
				continue
			}
			// Temp copy of rows, so we can add a new obstacle to it.
			newRows := make([][]rune, len(rows))
			for i := range rows {
				newRows[i] = make([]rune, len(rows[i]))
				copy(newRows[i], rows[i])
			}
			// Add a new obstacle.
			newRows[rowIdx][colIdx] = '#'
			p := position{location: location{x: start.x, y: start.y}, dir: start.dir}
			positions := make(map[position]int)
			for {
				positions[p] += 1
				next := p.next()
				if next.x < 0 || next.x >= len(newRows[0]) || next.y < 0 || next.y >= len(newRows) {
					// Next step is off the map.
					break
				}
				if newRows[next.y][next.x] != '#' {
					// Step forward into next.
					p.x, p.y = next.x, next.y
				} else {
					// We've hit a "#", turn right
					p.turn()
				}
				if _, ok := positions[p]; ok {
					// We've already been to this location and facing this direction, this is a loop.
					loops += 1
					break
				}
				positions[p] += 1
			}
		}
	}

	// Answer: 1836.
	fmt.Println("Solution 2:", loops)
}
