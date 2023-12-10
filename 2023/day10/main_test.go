package main

import (
	"strings"
	"testing"
)

func TestEnclosed(t *testing.T) {
	tests := []struct {
		desc string
		s    string
		want int
	}{
		{
			"first",
			`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`,
			4,
		},
		{
			"second",
			`..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`,
			4,
		},
		{
			"third",
			`.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`,
			8,
		},
		{
			"fourth",
			`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`,
			10,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			r := strings.NewReader(test.s)
			points, width := parse(r)
			start := findStart(points)
			walk(points, width, start, start)
			got := enclosed(points, width)
			if got != test.want {
				t.Errorf("enclosed(points, width) = %v, want: %v", got, test.want)
			}
		})
	}
}
