package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type tree struct {
	height   int
	maxLeft  int
	maxRight int
	maxUp    int
	maxDown  int
}

func (t tree) String() string {
	return fmt.Sprintf("{[%2d %2d %2d %2d] h:%d v:%v}", t.maxLeft, t.maxRight, t.maxUp, t.maxDown, t.height, t.visible())
}

func (t tree) visible() bool {
	// Outer edge trees are considered visible.
	if t.maxLeft == -1 || t.maxRight == -1 || t.maxUp == -1 || t.maxDown == -1 {
		return true
	}

	// A tree is visible if all the tree in any direction (left, right, up, down) are shorter than it.
	if t.height > t.maxLeft || t.height > t.maxRight || t.height > t.maxUp || t.height > t.maxDown {
		return true
	}

	return false
}

func addMaxes(grid [][]*tree) {
	// We've already populated maxLeft and maxUp during parsing, now we do the rest.

	// Find maxRight.
	// First: row 0, last column -> first column.
	// Second: row 1, last column -> first column.
	// Etc.
	for _, row := range grid {
		maxRight := -1
		for i := len(row) - 1; i >= 0; i-- {
			row[i].maxRight = maxRight
			if row[i].height > maxRight {
				maxRight = row[i].height
			}
		}
	}

	// Find maxDown.
	// First: col 0, last row -> first row.
	// Second: col 1, last row -> first row.
	// Etc.
	for colI := 0; colI < len(grid[0]); colI++ {
		maxDown := -1
		// Iterate over the rows in reverse.
		for rowI := len(grid) - 1; rowI >= 0; rowI-- {
			grid[rowI][colI].maxDown = maxDown
			if grid[rowI][colI].height > maxDown {
				maxDown = grid[rowI][colI].height
			}
		}
	}
}

// scenicScore calculates the scenic score.
// TODO: surely there's a more efficient way to do this... but I'm running out of time.
func scenicScore(tr *tree, rowI int, colI int, grid [][]*tree) int {

	left := 0
	trees := grid[rowI][:colI]
	for i := len(trees) - 1; i >= 0; i-- {
		left++
		if trees[i].height >= tr.height {
			break
		}
	}

	right := 0
	for _, t := range grid[rowI][colI+1:] {
		right++
		if t.height >= tr.height {
			break
		}
	}

	up := 0
	for i := rowI - 1; i >= 0; i-- {
		up++
		if grid[i][colI].height >= tr.height {
			break
		}
	}

	down := 0
	for _, row := range grid[rowI+1:] {
		down++
		if row[colI].height >= tr.height {
			break
		}
	}

	return left * right * up * down
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := [][]*tree{}
	maxUp := []int{}
	initialized := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Initialize maxUp on only the first pass.
		if !initialized {
			maxUp = make([]int, len([]rune(scanner.Text())))
			for i := range maxUp {
				maxUp[i] = -1
			}
			initialized = true
		}

		// Parse the row into trees.
		row := []*tree{}
		maxLeft := -1
		for colI, val := range scanner.Text() {
			height, err := strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, &tree{height: height, maxLeft: maxLeft, maxUp: maxUp[colI]})
			if height > maxLeft {
				maxLeft = height
			}
			if height > maxUp[colI] {
				maxUp[colI] = height
			}
		}
		grid = append(grid, row)
	}

	addMaxes(grid)
	total := 0
	for _, row := range grid {
		// fmt.Println()
		for _, t := range row {
			// fmt.Println(*t)
			if t.visible() {
				total++
			}
		}
	}
	fmt.Println("Solution 1:", total)

	max := 0
	for rowI, row := range grid {
		for colI, t := range row {
			score := scenicScore(t, rowI, colI, grid)
			if score > max {
				max = score
			}
		}
	}
	fmt.Println("Solution 2:", max)
}
