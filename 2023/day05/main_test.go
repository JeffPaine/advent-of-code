package main

import (
	"strings"
	"testing"
)

const input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestSolutionOne(t *testing.T) {
	want := 35
	seeds, bundles := parse(strings.NewReader(input))
	got := solutionOne(seeds, bundles)
	if got != want {
		t.Errorf("solutionOne(seeds, bundles) = %v, want %v", got, want)
	}
}

func TestSolutionTwo(t *testing.T) {
	want := 46
	seeds, bundles := parse(strings.NewReader(input))
	got := solutionTwo(seeds, bundles)
	if got != want {
		t.Errorf("solutionTwo(seeds, bundles) = %v, want %v", got, want)
	}
}

func BenchmarkSolutionTwo(b *testing.B) {
	seeds, bundles := parse(strings.NewReader(input))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solutionTwo(seeds, bundles)
	}
}
