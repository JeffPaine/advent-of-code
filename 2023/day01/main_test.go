package main

import "testing"

func TestFirstLastDigit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		if got := firstLastDigit(test.s); got != test.want {
			t.Errorf("firstLastDigit(%q) = %v, want %v", test.s, got, test.want)
		}
	}
}

func TestFirstLastNums(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, test := range tests {
		if got := firstLastNums(test.s); got != test.want {
			t.Errorf("firstLastNums(%q) = %v, want %v", test.s, got, test.want)
		}
	}
}
