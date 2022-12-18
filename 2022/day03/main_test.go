package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

type line struct {
	first, second []rune
}

func lines() []line {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := []line{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		l := line{
			first:  runes[:(len(runes) / 2)],
			second: runes[(len(runes) / 2):],
		}

		lines = append(lines, l)
	}
	return lines
}

func BenchmarkLinearCommon(b *testing.B) {
	lines := lines()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, l := range lines {
			_ = linearCommon(l.first, l.second)
		}
	}
}

func BenchmarkMapCommon(b *testing.B) {
	lines := lines()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, l := range lines {
			_ = mapCommon(l.first, l.second)
		}
	}
}
