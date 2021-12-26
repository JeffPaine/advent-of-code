package advent

import (
	"reflect"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7
	if got := CountIncreases(input); got != want {
		t.Errorf("CountIncreases(%v) = %v, want = %v", input, got, want)
	}
}

func TestRollingSumIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5
	if got := RollingSumIncreases(input); got != want {
		t.Errorf("RollingSumIncreases(%v) = %v, want %v", input, got, want)
	}
}

func TestParseCommands(t *testing.T) {
	input := []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	want := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	if got := ParseCommands(input); reflect.DeepEqual(got, want) != true {
		t.Errorf("ParseCommands(%v) = %v, want %v", input, got, want)
	}
}

func TestCalculatePosition(t *testing.T) {
	input := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	want := Position{Horizontal: 15, Depth: 10}
	if got := CalculatePosition(input); got != want {
		t.Errorf("CalculatePosition(%v) = %v, want %v", input, got, want)
	}
}

func TestCalculatePositionWithAim(t *testing.T) {
	input := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	want := Position{Horizontal: 15, Depth: 60, Aim: 10}
	if got := CalculatePositionWithAim(input); got != want {
		t.Errorf("CalculatePositionWithAim(%v) = %v, want %v", input, got, want)
	}
}

func TestNewReport(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	report := NewReport(input)
	want := 22
	if got := report.Gamma; got != want {
		t.Errorf("NewReport(%v).Gamma = %v, want %v", input, got, want)
	}
	want = 9
	if got := report.Epsilon; got != want {
		t.Errorf("NewReport(%v).Epsilon = %v, want %v", input, got, want)
	}
}
