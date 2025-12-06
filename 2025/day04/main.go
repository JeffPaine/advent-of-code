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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	var lines [][]rune

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line []rune
		for _, r := range scanner.Text() {
			line = append(line, r)
		}
		lines = append(lines, line)
	}

	// Pad the outer edge of lines so I don't have to do awkward
	// bounds checking.
	for i := range lines {
		lines[i] = append([]rune{'x'}, lines[i]...)
		lines[i] = append(lines[i], 'x')
	}
	var padding []rune
	for range len(lines[0]) {
		padding = append(padding, 'x')
	}
	lines = append([][]rune{padding}, lines...)
	lines = append(lines, padding)

	// Part 1.
	total1 := 0
	for rowIdx, line := range lines {
		for colIdx := range line {
			if lines[rowIdx][colIdx] != '@' {
				continue
			}
			total := 0
			offset := []int{-1, 0, 1}
			for _, rowOffset := range offset {
				for _, colOffset := range offset {
					if rowOffset == 0 && colOffset == 0 {
						continue
					}
					if lines[rowIdx+rowOffset][colIdx+colOffset] == '@' {
						total++
					}
				}
			}
			if total < 4 {
				total1++
			}
		}
	}

	// Answer: 1351
	fmt.Println("Solution 1:", total1)

	// Part 2.
	total2 := 0
	for {
		var toRemove []point
		for rowIdx, line := range lines {
			for colIdx := range line {
				if lines[rowIdx][colIdx] != '@' {
					continue
				}
				total := 0
				offset := []int{-1, 0, 1}
				for _, rowOffset := range offset {
					for _, colOffset := range offset {
						if rowOffset == 0 && colOffset == 0 {
							continue
						}
						if lines[rowIdx+rowOffset][colIdx+colOffset] == '@' {
							total++
						}
					}
				}
				if total < 4 {
					toRemove = append(toRemove, point{row: rowIdx, col: colIdx})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, p := range toRemove {
			lines[p.row][p.col] = '.'
			total2++
		}
	}

	// Answer: 8345
	fmt.Println("Solution 2:", total2)
}
