package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCountEnergized(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

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
	beams := []beam{{x: 0, y: 0, dir: right}}
	want := 46
	if got := countEnergized(grid, beams); got != want {
		t.Errorf("countEnergize(grid, beams, max) = %v, want: %v", got, want)
	}
}

func TestCountEnergizedOtherStart(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

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
	beams := []beam{{x: 3, y: 0, dir: down}}
	want := 51
	if got := countEnergized(grid, beams); got != want {
		t.Errorf("countEnergize(grid, beams, max) = %v, want: %v", got, want)
	}
}

func TestMaxEnergized(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

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
	want := 51
	if got := maxEnergized(grid); got != want {
		t.Errorf("maxEnergized(grid) = %v, want: %v", got, want)
	}
}
