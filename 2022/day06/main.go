package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func add(runes []rune, r rune, count int) []rune {
	if len(runes) < count {
		runes = append(runes, r)
		return runes
	}

	runes = append(runes[1:], r)
	return runes
}

func allDifferent(runes []rune, count int) bool {
	// Short slices don't count.
	if len(runes) < count {
		return false
	}

	m := map[rune]int{}
	for _, r := range runes {
		m[r]++
	}

	for _, val := range m {
		if val > 1 {
			return false
		}
	}

	return true
}

// strategyOne solves via allocating a new slice and populating a new map on each iteration.
func strategyOne(s string, count int) int {
	runes := []rune{}
	answer := 0
	for i, r := range s {
		runes = add(runes, r, count)
		if allDifferent(runes, count) {
			// Return the count, which is the index + 1.
			answer = i + 1
			break
		}
	}
	return answer
}

// strategyTwo solves via a rolling window of runes and a map of counts that's update on each iteration.
func strategyTwo(s string, window int) int {
	answer := 0
	runes := []rune(s)
	m := map[rune]int{}

	// Look at a sliding window of `window` runes for no duplicates.
	for i := 0; i < len(runes); i++ {
		m[runes[i]] += 1
		// Populate the map up to `window` length.
		if i < window-1 {
			continue
		}
		// If the window contains only unique runes, return the count of runes processed so far.
		if len(m) == window {
			answer = i + 1
			break
		}
		// Decrement (or delete) the oldest rune added.
		m[runes[i-window+1]]--
		if val, ok := m[runes[i-window+1]]; ok && val == 0 {
			delete(m, runes[i-window+1])
		}
	}
	return answer
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println("Solution 1:", strategyOne(input, 4))
	fmt.Println("Solution 2:", strategyOne(input, 14))
	fmt.Println("Solution 1:", strategyTwo(input, 4))
	fmt.Println("Solution 2:", strategyTwo(input, 14))
}
