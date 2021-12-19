package advent

import "testing"

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
		t.Errorf("RollingSumIncreases(%v) = %v, want = %v", input, got, want)
	}
}
