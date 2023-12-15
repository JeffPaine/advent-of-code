package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func reverse(s string) string {
	var out string
	for _, v := range s {
		out = string(v) + out
	}
	return out
}

func valueByColumn(rows []string) []int {
	matches := make([]int, len(rows[0]))

	for col := 1; col < len(rows[0]); col++ {
		mid := len(rows[0]) / 2
		length := 0
		if col < mid {
			length = col
		}
		if col == mid {
			length = col
		}
		if col > mid {
			length = len(rows[0]) - col
		}
		for _, row := range rows {
			leftStart := col - length
			if leftStart < 0 {
				leftStart = 0
			}
			leftEnd := leftStart + length
			rightStart := leftEnd
			rightEnd := rightStart + length
			l := row[leftStart:leftEnd]
			r := reverse(row[rightStart:rightEnd])
			if l == r {
				matches[col]++
			}
		}
	}

	var out []int
	for col, count := range matches {
		if count == len(rows) {
			out = append(out, col)
		}
	}
	return out
}

func valueFirst(rows []string) int {
	// Check for vertical (column) reflections.
	if val := valueByColumn(rows); len(val) != 0 {
		return val[0]
	}

	// Rotate 90 degrees counter-clockwise and check for vertical (column) reflections.
	rotated := make([]string, len(rows[0]))
	outRow := 0
	for i := len(rows[0]) - 1; i >= 0; i-- {
		for _, row := range rows {
			rotated[outRow] += string(row[i])
		}
		outRow++
	}
	if val := valueByColumn(rotated); len(val) != 0 {
		return val[0] * 100
	}

	return 0
}

func smudgeValue(rows []string) int {
	originalValue := valueFirst(rows)

	row1 := make([]string, len(rows))
	copy(row1, rows)
	for rowIdx := range row1 {
		for colIdx := range row1[rowIdx] {
			orig := row1[rowIdx]
			runes := []rune(orig)
			if runes[colIdx] == '#' {
				runes[colIdx] = '.'
			} else {
				runes[colIdx] = '#'
			}
			row1[rowIdx] = string(runes)
			vals := valueByColumn(row1)
			for _, val := range vals {
				if val != originalValue {
					return val
				}
			}
			row1[rowIdx] = orig
		}
	}

	rotated := make([]string, len(rows[0]))
	outRow := 0
	for i := len(rows[0]) - 1; i >= 0; i-- {
		for _, row := range rows {
			rotated[outRow] += string(row[i])
		}
		outRow++
	}
	for rowIdx := range rotated {
		for colIdx := range rotated[rowIdx] {
			orig := rotated[rowIdx]
			runes := []rune(orig)
			if runes[colIdx] == '#' {
				runes[colIdx] = '.'
			} else {
				runes[colIdx] = '#'
			}
			rotated[rowIdx] = string(runes)
			vals := valueByColumn(rotated)
			for _, val := range vals {
				if val*100 != originalValue {
					return val * 100
				}
			}
			rotated[rowIdx] = orig
		}
	}
	panic("no solution found")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var entries [][]string
	var curr []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			entries = append(entries, curr)
			curr = []string{}
			continue
		}
		curr = append(curr, scanner.Text())
	}
	entries = append(entries, curr)

	t1 := 0
	for _, entry := range entries {
		t1 += valueFirst(entry)
	}
	// Answer: 35360.
	fmt.Println("Solution 1:", t1)

	t2 := 0
	for _, entry := range entries {
		val := smudgeValue(entry)
		t2 += val
	}
	// Answer: 36755.
	fmt.Println("Solution 2:", t2)
}
