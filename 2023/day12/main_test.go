package main

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		s      string
		groups []int
		want   int
	}{
		{
			"???.###", []int{1, 1, 3}, 1,
		},
		{
			".??..??...?##.", []int{1, 1, 3}, 4,
		},
		{
			"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1,
		},
		{
			"????.#...#...", []int{4, 1, 1}, 1,
		},
		{
			"????.######..#####.", []int{1, 6, 5}, 4,
		},
		{
			"?###????????", []int{3, 2, 1}, 10,
		},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			cache := make(map[string]int)
			var cacheMu sync.RWMutex
			if got := count(test.s, test.groups, cache, &cacheMu); got != test.want {
				t.Errorf("count(%v, %v) = %v, want: %v", test.s, test.groups, got, test.want)
			}
		})
	}
}
