package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func tiltNorth(grid [][]rune) [][]rune {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 'O' {
				continue
			}
			// Shadow y for just the below loop.
			y := y
			for {
				if y <= 0 {
					break
				}
				if grid[y-1][x] == '#' || grid[y-1][x] == 'O' {
					break
				}
				// Move contents up one row.
				grid[y][x], grid[y-1][x] = grid[y-1][x], grid[y][x]
				y--
			}
		}
	}
	return grid
}

func spinCycle(grid [][]rune) [][]rune {
	// North.
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 'O' {
				continue
			}
			// Shadow y for just the below loop.
			y := y
			for {
				if y <= 0 {
					break
				}
				if grid[y-1][x] == '#' || grid[y-1][x] == 'O' {
					break
				}
				// Move contents up one row.
				grid[y][x], grid[y-1][x] = grid[y-1][x], grid[y][x]
				y--
			}
		}
	}
	// West.
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 'O' {
				continue
			}
			x := x
			for {
				if x <= 0 {
					break
				}
				if grid[y][x-1] == '#' || grid[y][x-1] == 'O' {
					break
				}
				grid[y][x], grid[y][x-1] = grid[y][x-1], grid[y][x]
				x--
			}
		}
	}
	// South.
	for y := len(grid) - 1; y >= 0; y-- {
		for x := range grid[y] {
			if grid[y][x] != 'O' {
				continue
			}
			// Shadow y for just the below loop.
			y := y
			for {
				if y >= len(grid)-1 {
					break
				}
				if grid[y+1][x] == '#' || grid[y+1][x] == 'O' {
					break
				}
				// Move contents up one row.
				grid[y][x], grid[y+1][x] = grid[y+1][x], grid[y][x]
				y++
			}
		}
	}
	// East.
	for y := range grid {
		for x := len(grid[0]) - 1; x >= 0; x-- {
			if grid[y][x] != 'O' {
				continue
			}
			x := x
			for {
				if x >= len(grid[0])-1 {
					break
				}
				if grid[y][x+1] == '#' || grid[y][x+1] == 'O' {
					break
				}
				grid[y][x], grid[y][x+1] = grid[y][x+1], grid[y][x]
				x++
			}
		}
	}
	return grid
}

func weigh(grid [][]rune) int {
	total := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				total += len(grid) - y
			}
		}
	}
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	grid1 := make([][]rune, len(grid))
	copy(grid1, grid)
	grid1 = tiltNorth(grid)
	// Answer: 110821.
	fmt.Println("Solution 1:", weigh(grid1))

	grid2 := make([][]rune, len(grid))
	copy(grid2, grid)
	// The problem calls for the state after 1 billion spins, but observing the
	// weights we can see that there's a repeating pattern.
	for i := 0; i < 1_000; i++ {
		grid = spinCycle(grid)
		// if i%100 == 0 && i != 0 {
		// 	fmt.Printf("iterations: %v, weight: %v\n", i, weigh(grid2))
		// }
	}
	// Answer: 83516.
	fmt.Println("Solution 2:", weigh(grid2))
}
