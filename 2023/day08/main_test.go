package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`
	want := 6
	r := strings.NewReader(input)
	steps, m := parse(r)
	got := partOne(steps, m)
	if got != want {
		t.Errorf("partOne(steps, m) = %v, want: %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `LR
	11A = (11B, XXX)
	11B = (XXX, 11Z)
	11Z = (11B, XXX)
	22A = (22B, XXX)
	22B = (22C, 22C)
	22C = (22Z, 22Z)
	22Z = (22B, 22B)
	XXX = (XXX, XXX)`
	want := 6
	r := strings.NewReader(input)
	steps, m := parse(r)
	got := partTwo(steps, m)
	if got != want {
		t.Errorf("partTwo(steps, m) = %v, want: %v", got, want)
	}
}

func BenchmarkPartTwoShort(b *testing.B) {
	// Note: this input only has two starting nodes, so its performance is not very indicative of
	// the actually problem data, which had 6 starting nodes.
	input := `LR
	11A = (11B, XXX)
	11B = (XXX, 11Z)
	11Z = (11B, XXX)
	22A = (22B, XXX)
	22B = (22C, 22C)
	22C = (22Z, 22Z)
	22Z = (22B, 22B)
	XXX = (XXX, XXX)`
	r := strings.NewReader(input)
	steps, m := parse(r)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = partTwo(steps, m)
	}
}

func BenchmarkPartTwoLong(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	steps, m := parse(f)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = partTwo(steps, m)
	}
}
