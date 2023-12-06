package main

import "testing"

func TestWins(t *testing.T) {
	tests := []struct {
		r    race
		want int
	}{
		{
			race{time: 7, dist: 9},
			4,
		},
		{
			race{time: 15, dist: 40},
			8,
		},
		{
			race{time: 30, dist: 200},
			9,
		},
	}

	for _, test := range tests {
		if got := test.r.wins(); got != test.want {
			t.Errorf("r.wins() = %v, want %v", got, test.want)
		}
	}

}

func TestTotalWins(t *testing.T) {
	races := []race{
		{time: 7, dist: 9},
		{time: 15, dist: 40},
		{time: 30, dist: 200},
	}
	want := 288
	if got := totalWins(races); got != want {
		t.Errorf("totalWins(races) = %v, want %v", got, want)
	}
}
