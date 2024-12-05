package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func valid(update []int, rules map[int][]int) bool {
	result := true
	for idx, num := range update {
		ruleVals, ok := rules[num]
		if !ok {
			// No rules for the current number.
			continue
		}
		for _, val := range ruleVals {
			valIdx := slices.Index(update, val)
			if valIdx == -1 {
				continue
			}
			if idx > valIdx {
				// Update ordering breaks a given rule.
				result = false
			}
		}
	}
	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// Parse rules.
	rules := make(map[int][]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		ints := strings.Split(scanner.Text(), "|")
		a, err := strconv.Atoi(ints[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(ints[1])
		if err != nil {
			log.Fatal(err)
		}
		rules[a] = append(rules[a], b)
	}

	// Parse updates.
	updates := [][]int{}
	for scanner.Scan() {
		update := []int{}
		nums := strings.Split(scanner.Text(), ",")
		for _, text := range nums {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	total1 := 0
	for _, update := range updates {
		if !valid(update, rules) {
			continue
		}
		total1 += update[(len(update) / 2)]
	}

	// Answer: 6260.
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for _, update := range updates {
		if valid(update, rules) {
			continue
		}
		for range update {
			for idx := 0; idx < len(update); idx++ {
				ruleVals, ok := rules[update[idx]]
				if !ok {
					// No rules for the current number.
					continue
				}
				for _, ruleVal := range ruleVals {
					ruleValIdx := slices.Index(update, ruleVal)
					if ruleValIdx == -1 {
						// Rule value not in update.
						continue
					}
					if idx > ruleValIdx {
						update[idx], update[ruleValIdx] = update[ruleValIdx], update[idx]
					}
				}
			}
		}
		total2 += update[(len(update) / 2)]
	}

	// Answer: 5346.
	fmt.Println("Solution 2:", total2)
}
