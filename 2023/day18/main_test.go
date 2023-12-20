package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCalculateVolume(t *testing.T) {
	input := `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	var steps []step
	for scanner.Scan() {
		steps = append(steps, parseStep(scanner.Text()))
	}
	points := makePoints(steps)
	want := 62
	if got := calculateVolume(steps, points); got != want {
		t.Errorf("calculateVolume(steps, points)= %v, want: %v", got, want)
	}
}

func TestCalculateVolumeWithHex(t *testing.T) {
	input := `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	var steps []step
	for scanner.Scan() {
		steps = append(steps, parseHex(scanner.Text()))
	}
	points := makePoints(steps)
	want := 952408144115
	if got := calculateVolume(steps, points); got != want {
		t.Errorf("calculateVolume(steps, points)= %v, want: %v", got, want)
	}
}
