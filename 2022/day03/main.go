package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type bag struct {
	contents string
}

func (b bag) String() string {
	return fmt.Sprintf("contents: %v, first: %v, second: %v, common: %c", b.contents, string(b.first()), string(b.second()), b.common())
}

func (b bag) first() []rune {
	runes := []rune(b.contents)
	return runes[:(len(runes) / 2)]
}

func (b bag) second() []rune {
	runes := []rune(b.contents)
	return runes[(len(runes) / 2):]
}

func (b bag) common() rune {
	return linearCommon(b.first(), b.second())
}

func linearCommon(first, second []rune) rune {
	var common rune

loop:
	for _, fr := range first {
		for _, sr := range second {
			if fr == sr {
				common = fr
				break loop
			}
		}

	}
	return common
}

func mapCommon(first, second []rune) rune {
	var common rune
	m := map[rune]struct{}{}
	for _, r := range first {
		m[r] = struct{}{}
	}
	for _, r := range second {
		if _, ok := m[r]; ok {
			common = r
			break
		}
	}
	return common
}

func priority(r rune) int {
	lowercase := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	uppercase := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	priorities := map[rune]int{}
	curr := 1
	for _, c := range lowercase {
		priorities[c] = curr
		curr += 1
	}
	for _, c := range uppercase {
		priorities[c] = curr
		curr += 1
	}
	return priorities[r]
}

func findCommon(bags []bag) rune {
	// Map of a rune to the indexes of the bags it was found in.
	freqs := map[rune][]bool{}
	for i, b := range bags {
		for _, r := range b.contents {
			if len(freqs[r]) == 0 {
				// Assumes len(bags) == 3.
				freqs[r] = []bool{false, false, false}
			}
			freqs[r][i] = true
		}
	}

	// Find the common rune.
	var common rune
	for k, v := range freqs {
		if v[0] && v[1] && v[2] {
			common = k
		}
	}
	return common
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bags := []bag{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := bag{contents: scanner.Text()}
		bags = append(bags, b)
	}

	total := 0
	for _, b := range bags {
		total += priority(b.common())
	}
	fmt.Println("Solution 1:", total)

	total = 0
	for i := 0; i < len(bags); i += 3 {
		group := []bag{bags[i], bags[i+1], bags[i+2]}
		total += priority(findCommon(group))
	}
	fmt.Println("Solution 2:", total)
}
