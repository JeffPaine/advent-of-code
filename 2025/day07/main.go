package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	row int
	col int
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func processGrid(start point, m map[point]int, splitters map[point]struct{}, grid [][]rune) int {
	if count, ok := m[start]; ok {
		return count
	}

	total := 0
	p := start

	for {
		if p.row >= len(grid) {
			// Hit the end of the grid.
			total = 1
			break
		}
		if grid[p.row][p.col] != '^' {
			// Optional visual for debugging.
			// grid[p.row][p.col] = '|'
			p.row += 1
			continue
		}

		// We're now at a splitter.
		splitters[p] = struct{}{}

		left := 0
		right := 0

		// Left of splitter.
		if p.col-1 >= 0 {
			left = processGrid(point{row: p.row, col: p.col - 1}, m, splitters, grid)
		}
		// Right of splitter.
		if p.col+1 <= len(grid[0]) {
			right = processGrid(point{row: p.row, col: p.col + 1}, m, splitters, grid)
		}
		total = left + right
		break
	}

	m[start] = total

	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	var grid [][]rune

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line []rune
		for _, r := range scanner.Text() {
			line = append(line, r)
		}
		grid = append(grid, line)
	}

	var root point

	// Find the start.
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				root = point{row: row, col: col}
			}
		}
	}

	splitters := make(map[point]struct{})
	timelines := processGrid(root, make(map[point]int), splitters, grid)

	// Answer: 1651
	fmt.Println("Solution 1:", len(splitters))

	// Answer: 108924003331749
	fmt.Println("Solution 2:", timelines)
}
