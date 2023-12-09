package main

import "testing"

func TestPredictForward(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{0, 3, 6, 9, 12, 15},
			18,
		}, {
			[]int{1, 3, 6, 10, 15, 21},
			28,
		}, {
			[]int{10, 13, 16, 21, 30, 45},
			68,
		},
	}
	for _, test := range tests {
		if got := predictForward(test.nums); got != test.want {
			t.Errorf("predictForward(%v) = %v, want: %v", test.nums, got, test.want)
		}
	}
}

func TestPredictBackward(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{0, 3, 6, 9, 12, 15},
			-3,
		}, {
			[]int{1, 3, 6, 10, 15, 21},
			0,
		}, {
			[]int{10, 13, 16, 21, 30, 45},
			5,
		},
	}
	for _, test := range tests {
		if got := predictBackward(test.nums); got != test.want {
			t.Errorf("predictBackward(%v) = %v, want: %v", test.nums, got, test.want)
		}
	}
}
