package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type elf struct {
	items []int
}

func (e elf) calories() int {
	total := 0
	for _, item := range e.items {
		total += item
	}
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Solution 1: find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
	elves := []elf{}
	items := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// End of a single elf's list of items.
		if scanner.Text() == "" {
			e := elf{}
			// Slices in Go are passed by reference, so we need to copy contents rather than use s1 = s2 style assignment
			// otherwise the slices would refer to the same object.
			e.items = make([]int, len(items))
			copy(e.items, items)
			elves = append(elves, e)
			items = []int{}
			continue
		}

		var item int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &item)
		if err != nil {
			log.Fatalf("fmt.Sscanf(%q, \"%%d\", %v) = %v", scanner.Text(), item, err)
		}
		items = append(items, item)
	}

	max := 0
	for _, e := range elves {
		if e.calories() > max {
			max = e.calories()
		}
	}
	log.Println("Solution 1:", max)

	// Solution 2: find the total Calories carried by the top three Elves carrying the most Calories.
	cals := []int{}
	for _, c := range elves {
		cals = append(cals, c.calories())
	}
	sort.Ints(cals)
	log.Println("Solution 2:", cals[len(cals)-1]+cals[len(cals)-2]+cals[len(cals)-3])
}
