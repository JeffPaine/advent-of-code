package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	all := []string{}

	// Forward.
	all = append(all, rows...)

	// Vertical.
	for col := range len(rows[0]) {
		vertical := ""
		for _, row := range rows {
			vertical += string(row[col])
		}
		all = append(all, vertical)
	}

	// Diagonals down and to the right along the first row.
	for col := 0; col < len(rows[0]); col++ {
		diagonal := ""
		for newCol, row := col, 0; newCol < len(rows[0]) && row < len(rows); newCol, row = newCol+1, row+1 {
			diagonal += string(rows[row][newCol])
		}
		all = append(all, diagonal)
	}

	// Diagonals down and to the left along the first row.
	for col := len(rows[0]) - 1; col >= 0; col-- {
		diagonal := ""
		for newCol, row := col, 0; newCol >= 0 && row < len(rows); newCol, row = newCol-1, row+1 {
			diagonal += string(rows[row][newCol])
		}
		all = append(all, diagonal)
	}

	// Diagonals down and to the right along the left column.
	for row := 1; row < len(rows); row++ {
		diagonal := ""
		for col, newRow := 0, row; col < len(rows[0]) && newRow < len(rows); col, newRow = col+1, newRow+1 {
			diagonal += string(rows[newRow][col])
		}
		all = append(all, diagonal)
	}

	// Diagonals down and to the left along the right column.
	for row := 1; row < len(rows); row++ {
		diagonal := ""
		for col, newRow := len(rows[0])-1, row; col >= 0 && newRow < len(rows); col, newRow = col-1, newRow+1 {
			diagonal += string(rows[newRow][col])
		}
		all = append(all, diagonal)
	}

	total1 := 0
	for _, s := range all {
		total1 += strings.Count(s, "XMAS")
		total1 += strings.Count(s, "SAMX")
	}

	// Answer: 2549.
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for row := 1; row < len(rows)-1; row++ {
		for col := 1; col < len(rows[0])-1; col++ {
			if rows[row][col] != 'A' {
				continue
			}
			a := string([]byte{rows[row-1][col-1], rows[row][col], rows[row+1][col+1]})
			b := string([]byte{rows[row-1][col+1], rows[row][col], rows[row+1][col-1]})
			subTotal := 0
			for _, s := range []string{a, b} {
				subTotal += strings.Count(s, "MAS")
				subTotal += strings.Count(s, "SAM")
			}
			if subTotal < 2 {
				continue
			}
			total2 += 1
		}
	}

	// Answer: 2003.
	fmt.Println("Solution 2:", total2)
}
