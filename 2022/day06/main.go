package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func add(runes []rune, r rune, count int) []rune {
	out := make([]rune, len(runes))
	copy(out, runes)

	if len(runes) < count {
		out = append(out, r)
		return out
	}

	out = append(out[1:], r)
	return out
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

	runes1 := []rune{}
	sol1 := 0
	for i, r := range input {
		runes1 = add(runes1, r, 4)
		if allDifferent(runes1, 4) {
			// Return the count, which is the index + 1.
			sol1 = i + 1
			break
		}
	}
	fmt.Println("Solution 1:", sol1)

	runes2 := []rune{}
	sol2 := 0
	for i, r := range input {
		runes2 = add(runes2, r, 14)
		if allDifferent(runes2, 14) {
			// Return the count, which is the index + 1.
			sol2 = i + 1
			break
		}
	}
	fmt.Println("Solution 2:", sol2)
}
