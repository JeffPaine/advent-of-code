package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func input() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	out := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		out = scanner.Text()
	}
	return out
}

func BenchmarkStrategyOne(b *testing.B) {
	var tests = []struct {
		count int
	}{
		{4},
		{14},
	}

	for _, test := range tests {
		b.Run(fmt.Sprintf("count_%d", test.count), func(b *testing.B) {
			input := input()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strategyOne(input, test.count)
			}
		})
	}

}

func BenchmarkStrategyTwo(b *testing.B) {
	var tests = []struct {
		count int
	}{
		{4},
		{14},
	}

	for _, test := range tests {
		b.Run(fmt.Sprintf("count_%d", test.count), func(b *testing.B) {
			input := input()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = strategyTwo(input, test.count)
			}
		})
	}

}
