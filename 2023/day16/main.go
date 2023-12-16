package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type symbol rune

const (
	// .
	emptySpace symbol = iota
	// /
	mirrorRight
	// \
	mirrorLeft
	// |
	vertical
	// -
	horizontal
)

func parseSymbol(r rune) symbol {
	switch r {
	case '.':
		return emptySpace
	case '/':
		return mirrorRight
	case '\\':
		return mirrorLeft
	case '|':
		return vertical
	case '-':
		return horizontal
	default:
		panic(fmt.Sprintf("unsupported symbol: %v", r))
	}
}

type tile struct {
	symbol    symbol
	energized bool
}

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

type beam struct {
	x   int
	y   int
	dir direction
}

func (b beam) String() string {
	return fmt.Sprintf("{x: %v, y: %v, dir: %v}", b.x, b.y, b.dir)
}

func (b *beam) moveToNext() {
	switch b.dir {
	case right:
		b.x++
	case left:
		b.x--
	// Up visually is actually decreasing the y value.
	case up:
		b.y--
	// Down visually is actually increasing the y value.
	case down:
		b.y++
	}
}

type visit struct {
	x   int
	y   int
	dir direction
}

func countEnergized(grid [][]tile, beams []beam) int {
	visited := make(map[visit]struct{})
	for {
		moved := 0
		for i := range beams {
			// Skip locations + directions we've already done.
			v := visit{x: beams[i].x, y: beams[i].y, dir: beams[i].dir}
			if _, ok := visited[v]; ok {
				continue
			}
			// Record visit.
			visited[v] = struct{}{}

			// Skip beams that are off-grid.
			if beams[i].x < 0 || beams[i].x >= len(grid[0]) || beams[i].y < 0 || beams[i].y >= len(grid) {
				continue
			}

			grid[beams[i].y][beams[i].x].energized = true

			sym := grid[beams[i].y][beams[i].x].symbol
			switch sym {
			// /
			case mirrorRight:
				switch beams[i].dir {
				case right:
					beams[i].dir = up
				case left:
					beams[i].dir = down
				case up:
					beams[i].dir = right
				case down:
					beams[i].dir = left
				}
				// \
			case mirrorLeft:
				switch beams[i].dir {
				case right:
					beams[i].dir = down
				case left:
					beams[i].dir = up
				case up:
					beams[i].dir = left
				case down:
					beams[i].dir = right
				}
				// |
			case vertical:
				switch beams[i].dir {
				case right, left:
					beams[i].dir = up

					x := beams[i].x
					y := beams[i].y
					dir := down
					b := beam{x: x, y: y, dir: dir}
					beams = append(beams, b)
				}
				// -
			case horizontal:
				switch beams[i].dir {
				case up, down:
					beams[i].dir = left

					x := beams[i].x
					y := beams[i].y
					dir := right
					b := beam{x: x, y: y, dir: dir}
					beams = append(beams, b)
				}
			case emptySpace:
				// Do nothing.
			default:
				panic(fmt.Sprintf("unsupported symbol: %v", sym))
			}
			beams[i].moveToNext()
			moved++
		}
		if moved == 0 {
			break
		}
	}
	return sumEnergized(grid)
}

func sumEnergized(grid [][]tile) int {
	total := 0
	for _, row := range grid {
		for _, t := range row {
			if t.energized {
				total++
			}
		}
	}
	return total
}

func maxEnergized(grid [][]tile) int {
	starts := []beam{}
	width := len(grid[0])
	height := len(grid)
	// Sides.
	for i := 0; i < height; i++ {
		starts = append(starts, beam{x: 0, y: i, dir: right})
		starts = append(starts, beam{x: width - 1, y: i, dir: left})
	}
	// Top and bottom.
	for i := 0; i < width; i++ {
		starts = append(starts, beam{x: i, y: 0, dir: down})
		starts = append(starts, beam{x: i, y: height - 1, dir: up})
	}

	max := 0
	for _, s := range starts {
		var tempGrid [][]tile
		for _, row := range grid {
			tempRow := make([]tile, len(row))
			copy(tempRow, row)
			tempGrid = append(tempGrid, tempRow)
		}
		count := countEnergized(tempGrid, []beam{s})
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var grid [][]tile
	for scanner.Scan() {
		var row []tile
		for _, r := range scanner.Text() {
			sym := parseSymbol(r)
			t := tile{symbol: sym}
			row = append(row, t)
		}
		grid = append(grid, row)
	}

	var grid1 [][]tile
	for _, row := range grid {
		tempRow := make([]tile, len(row))
		copy(tempRow, row)
		grid1 = append(grid1, tempRow)
	}
	beams := []beam{{x: 0, y: 0, dir: right}}
	// Answer: 6740.
	fmt.Println("Solution 1:", countEnergized(grid1, beams))

	var grid2 [][]tile
	for _, row := range grid {
		tempRow := make([]tile, len(row))
		copy(tempRow, row)
		grid2 = append(grid2, tempRow)
	}
	max := maxEnergized(grid2)
	// Answer: 7041.
	fmt.Println("Solution 2:", max)
}
