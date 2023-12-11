package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type galaxy struct {
	id int
	x  int
	y  int
}

func distance(g1, g2 galaxy, size int, emptyRows, emptyCols []int) int {
	xa := g1.x
	xb := g2.x
	if xa > xb {
		xa, xb = xb, xa
	}
	xCrosses := 0
	for col := xa; col <= xb; col++ {
		if slices.Contains(emptyCols, col) {
			xCrosses++
		}
	}

	ya := g1.y
	yb := g2.y
	if ya > yb {
		ya, yb = yb, ya
	}
	yCrosses := 0
	for row := ya; row <= yb; row++ {
		if slices.Contains(emptyRows, row) {
			yCrosses++
		}
	}

	xDiff := xb - xa - xCrosses + (xCrosses * size)
	yDiff := yb - ya - yCrosses + (yCrosses * size)
	return xDiff + yDiff
}

func findEmptyRows(grid [][]string) []int {
	var out []int
	for idx, row := range grid {
		galaxies := 0
		for _, s := range row {
			if s == "#" {
				galaxies++
			}
		}
		if galaxies == 0 {
			out = append(out, idx)
		}
	}
	return out
}

func findEmptyCols(grid [][]string) []int {
	var out []int
	width := len(grid[0])
	for col := 0; col < width; col++ {
		galaxies := 0
		for _, row := range grid {
			if row[col] == "#" {
				galaxies++
			}
		}
		if galaxies == 0 {
			out = append(out, col)
		}
	}
	return out
}

func findGalaxies(grid [][]string) []galaxy {
	var out []galaxy
	id := 1
	for rowIdx, row := range grid {
		for colIdx, col := range row {
			if col == "#" {
				g := galaxy{id: id, y: rowIdx, x: colIdx}
				out = append(out, g)
				id++
			}
		}
	}
	return out
}

func totalDistances(grid [][]string, size int) int {
	emptyRows := findEmptyRows(grid)
	emptyCols := findEmptyCols(grid)
	galaxies := findGalaxies(grid)
	total := 0
	for i := 0; i < len(galaxies); i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			dist := distance(a, b, size, emptyRows, emptyCols)
			total += dist
		}
	}
	return total
}

func main() {
	// 	input := `...#......
	// .......#..
	// #.........
	// ..........
	// ......#...
	// .#........
	// .........#
	// ..........
	// .......#..
	// #...#.....`
	// 	r := strings.NewReader(input)
	// 	scanner := bufio.NewScanner(r)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(f)

	grid := [][]string{}

	for scanner.Scan() {
		row := []string{}
		for _, r := range scanner.Text() {
			row = append(row, string(r))
		}
		grid = append(grid, row)
	}

	// Answer: 10885634.
	fmt.Println("Solution 1:", totalDistances(grid, 2))
	// Answer: 707505470642.
	fmt.Println("Solution 2:", totalDistances(grid, 1_000_000))
}
