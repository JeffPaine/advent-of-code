package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type assignment struct {
	start  int
	finish int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pairs := [][]assignment{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Lines look like "2-4,6-8".
		var a assignment
		var b assignment
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a.start, &a.finish, &b.start, &b.finish)
		if err != nil {
			log.Fatal(err)
		}
		pairs = append(pairs, []assignment{a, b})
	}

	total := 0
	for _, p := range pairs {
		a, b := p[0], p[1]
		// One completely contains the other.
		if (a.start <= b.start && a.finish >= b.finish) || (b.start <= a.start && b.finish >= a.finish) {
			total++
		}
	}
	fmt.Println("Solution 1:", total)

	total = 0
	for _, p := range pairs {
		a, b := p[0], p[1]
		// One contains a single part of the other.
		if (a.start <= b.start && a.finish >= b.start) || (b.start <= a.start && b.finish >= a.start) {
			total++
		}
	}
	fmt.Println("Solution 2:", total)
}
