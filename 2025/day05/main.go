package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	isRanges := true
	var ranges []interval
	var ids []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			isRanges = false
			continue
		}
		if isRanges {
			nums := strings.Split(scanner.Text(), "-")
			start, _ := strconv.Atoi(nums[0])
			end, _ := strconv.Atoi(nums[1])
			r := interval{start: start, end: end}
			ranges = append(ranges, r)
		} else {
			id, _ := strconv.Atoi(scanner.Text())
			ids = append(ids, id)
		}
	}

	slices.SortFunc(ranges, func(a, b interval) int {
		return cmp.Compare(a.start, b.start)
	})

	var combined []interval
	for _, b := range ranges {
		if len(combined) == 0 || combined[len(combined)-1].end < b.start {
			combined = append(combined, b)
			continue
		}
		// At this point, b's start is in a's range.
		combined[len(combined)-1].end = max(combined[len(combined)-1].end, b.end)
	}

	total1 := 0

	// Sort ids so we can rule out ranges from the next round as we iterate through them.
	slices.Sort(ids)
	idx := 0
	for _, id := range ids {
		for idx < len(combined) {
			if combined[idx].start <= id && id <= combined[idx].end {
				total1++
				break
			}
			idx++
		}
	}

	// Answer: 885
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for _, r := range combined {
		total2 += r.end - r.start + 1
	}

	// Answer: 348115621205535
	fmt.Println("Solution 2:", total2)
}
