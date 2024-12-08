package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type location struct {
	row int
	col int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	antennas := make(map[rune][]location)
	maxRow := 0
	maxCol := 0

	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		col := 0
		for _, freq := range scanner.Text() {
			if freq != '.' {
				antennas[freq] = append(antennas[freq], location{row: row, col: col})
			}
			maxCol = col
			col++
		}
		maxRow = row
		row++
	}

	antinodes1 := make(map[location]struct{})

	for _, locs := range antennas {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j <= len(locs)-1; j++ {
				a := locs[i]
				b := locs[j]
				rowDiff := a.row - b.row
				colDiff := a.col - b.col
				antinodes1[location{row: a.row + rowDiff, col: a.col + colDiff}] = struct{}{}
				antinodes1[location{row: b.row - rowDiff, col: b.col - colDiff}] = struct{}{}
			}
		}
	}

	total1 := 0
	for node := range antinodes1 {
		if node.row < 0 || node.row > maxRow || node.col < 0 || node.col > maxCol {
			continue
		}
		total1++
	}
	// Answer: 367.
	fmt.Println("Solution 1:", total1)

	antinodes2 := make(map[location]struct{})

	for _, locs := range antennas {
		if len(locs) < 2 {
			continue
		}
		for _, loc := range locs {
			antinodes2[loc] = struct{}{}
		}
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j <= len(locs)-1; j++ {
				a := locs[i]
				b := locs[j]
				rowDiff := a.row - b.row
				colDiff := a.col - b.col
				for factor := 1; ; factor++ {
					node := location{row: a.row + rowDiff*factor, col: a.col + colDiff*factor}
					if node.row < 0 || node.row > maxRow || node.col < 0 || node.col > maxCol {
						break
					}
					antinodes2[node] = struct{}{}
				}
				for factor := 1; ; factor++ {
					node := location{row: b.row - rowDiff*factor, col: b.col - colDiff*factor}
					if node.row < 0 || node.row > maxRow || node.col < 0 || node.col > maxCol {
						break
					}
					antinodes2[node] = struct{}{}
				}
			}
		}
	}

	total2 := 0
	for node := range antinodes2 {
		if node.row < 0 || node.row > maxRow || node.col < 0 || node.col > maxCol {
			continue
		}
		total2++
	}
	// Answer: 1285.
	fmt.Println("Solution 2:", total2)
}
