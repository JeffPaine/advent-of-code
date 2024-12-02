package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func diffsOK(levels []int) bool {
	diffs := []int{}
	for i := 1; i < len(levels); i++ {
		diffs = append(diffs, levels[i]-levels[i-1])
	}
	positiveOK := true
	for _, diff := range diffs {
		if diff < 1 || diff > 3 {
			positiveOK = false
		}
	}
	negativeOK := true
	for _, diff := range diffs {
		if diff > -1 || diff < -3 {
			negativeOK = false
		}
	}
	return positiveOK || negativeOK
}

func orderOK(levels []int) bool {
	increasingOK := true
	for _, level := range levels {
		if level < 1 {
			increasingOK = false
		}
	}
	decreasingOK := true
	for _, level := range levels {
		if level > -1 {
			decreasingOK = false
		}
	}
	return increasingOK || decreasingOK
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levels := []int{}
		entries := strings.Split(scanner.Text(), " ")
		for _, entry := range entries {
			num, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}

	total1 := 0
	for _, levels := range reports {
		if diffsOK(levels) && orderOK(levels) {
			total1 += 1
		}
	}

	// Answer: 236.
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for _, levels := range reports {
		// First, try normal checks.
		if diffsOK(levels) && orderOK(levels) {
			total2 += 1
			continue
		}

		// Next, try removing a single level and then doing normal checks.
		for i := 0; i < len(levels); i++ {
			newLevels := []int{}
			newLevels = append(newLevels, levels[:i]...)
			newLevels = append(newLevels, levels[i+1:]...)
			if orderOK(newLevels) && diffsOK(newLevels) {
				total2 += 1
				// Only count one, working permutation.
				break
			}
		}
	}

	// Answer: 308.
	fmt.Println("Solution 2:", total2)
}
