package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		g1        galaxy
		g2        galaxy
		size      int
		emptyRows []int
		emptyCols []int
		want      int
	}{
		{
			galaxy{id: 1, x: 3, y: 0},
			galaxy{id: 7, x: 7, y: 8},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			15,
		},
		{
			galaxy{id: 3, x: 0, y: 2},
			galaxy{id: 6, x: 9, y: 6},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			17,
		},
		{
			galaxy{id: 8, x: 0, y: 9},
			galaxy{id: 9, x: 4, y: 9},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			5,
		},
		{
			galaxy{id: 5, x: 1, y: 5},
			galaxy{id: 9, x: 4, y: 9},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			9,
		},
		{
			galaxy{id: 3, x: 0, y: 2},
			galaxy{id: 6, x: 9, y: 6},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			17,
		},
		{
			galaxy{id: 1, x: 3, y: 0},
			galaxy{id: 2, x: 7, y: 1},
			2,
			[]int{3, 7},
			[]int{2, 5, 8},
			6,
		},
		{
			galaxy{id: 1, x: 3, y: 0},
			galaxy{id: 2, x: 7, y: 1},
			10,
			[]int{3, 7},
			[]int{2, 5, 8},
			14,
		},
	}
	for _, test := range tests {
		if got := distance(test.g1, test.g2, test.size, test.emptyRows, test.emptyCols); got != test.want {
			t.Errorf("distance(%v, %v, %v, %v, %v) = %v, want: %v", test.g1, test.g2, test.size, test.emptyRows, test.emptyCols, got, test.want)
		}
	}
}

func TestTotalDistances(t *testing.T) {
	tests := []struct {
		size int
		want int
	}{
		{
			2,
			374,
		},
		{
			10,
			1030,
		},
		{
			100,
			8410,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("size_%v", test.size), func(t *testing.T) {
			input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
			r := strings.NewReader(input)
			scanner := bufio.NewScanner(r)
			grid := [][]string{}

			for scanner.Scan() {
				row := []string{}
				for _, r := range scanner.Text() {
					row = append(row, string(r))
				}
				grid = append(grid, row)
			}
			if got := totalDistances(grid, test.size); got != test.want {
				t.Errorf("totalDistances(grid, %v) = %v, want: %v", test.size, got, test.want)
			}
		})
	}
}
