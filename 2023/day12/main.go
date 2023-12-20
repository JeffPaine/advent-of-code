package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parse(s string) (string, []int) {
	parts := strings.Split(s, " ")

	var out []int
	for _, char := range strings.Split(parts[1], ",") {
		val, err := strconv.Atoi(char)
		if err != nil {
			panic(fmt.Sprintf("cannot convert char %v to int", char))
		}
		out = append(out, val)
	}
	return parts[0], out
}

func unfold(s string, nums []int) (string, []int) {
	outS := strings.Join([]string{s, s, s, s, s}, "?")
	var outNums []int
	count := 5
	for i := 0; i < count; i++ {
		outNums = append(outNums, nums...)
	}

	return outS, outNums
}

// count counts the possible solutions to s. It uses the strategy outlined in
// https://www.youtube.com/watch?v=g3Ms5e7Jdqo.
func count(s string, nums []int, cache map[string]int, cacheMu *sync.RWMutex) int {
	key := s
	for _, num := range nums {
		key += fmt.Sprintf("-%v", num)
	}
	if cache != nil {
		cacheMu.RLock()
		val, ok := cache[key]
		cacheMu.RUnlock()
		if ok {
			return val
		}
	}

	if s == "" {
		if len(nums) == 0 {
			// An empty string with empty nums has a single valid solution (itself).
			return 1
		} else {
			// An empty string with non-empty nums (e.g. the expectation that there's still broken
			// springs) has no valid solutions.
			return 0
		}
	}

	if len(nums) == 0 {
		if strings.Contains(s, "#") {
			// We have no more nums to account for more broken springs, but there's still at least
			// one broken spring left.
			return 0
		} else {
			// No more nums and no more broken springs, this is a valid match.
			return 1
		}
	}

	total := 0
	if s[0] == '.' {
		// Single working springs are irrelevant for counting broken springs, ignore it.
		total += count(s[1:], nums, cache, cacheMu)
	}
	if s[0] == '#' {
		if len(s) < nums[0] {
			// No valid solution exists if we expect more broken springs left than springs.
			return total
		}

		// This is the start of a block, how do we know if this block is valid (countable)?
		//
		// 1. The expected block of broken springs is not longer than the remaining springs.
		isNotTooLong := nums[0] <= len(s)
		// 2. The expected block of broken springs doesn't contain any operational springs.
		hasNoOperational := !strings.Contains(s[:nums[0]], ".")
		// 3. The expected block of broken springs has a valid end: either it coincides with the end
		// of the string or the first char after this block is not a broken spring (which would
		// imply a longer than expected block).
		hasValidEnd := (nums[0] == len(s) || s[nums[0]] != '#')
		if isNotTooLong && hasNoOperational && hasValidEnd {
			var newS string
			if len(s) > nums[0] {
				// +1 below so we skip over the char after the current block as it's either a "." in
				// which case we ignore it, or it's a "?" and we don't want it to get optionally turned
				// into a "#".
				newS = s[nums[0]+1:]
			}
			var newNums []int
			if len(nums) > 1 {
				newNums = nums[1:]
			}
			total += count(newS, newNums, cache, cacheMu)
		}
	}
	if s[0] == '?' {
		// Try making it a ".".
		next := ""
		if len(s) > 1 {
			next = s[1:]
		}
		total += count(next, nums, cache, cacheMu)

		if len(s) < nums[0] {
			// No valid solution exists if we expect more broken springs left than springs.
			return total
		}

		// Try making it a "#".
		isNotTooLong := nums[0] <= len(s)
		hasNoOperational := !strings.Contains(s[:nums[0]], ".")
		hasValidEnd := (nums[0] == len(s) || s[nums[0]] != '#')
		if isNotTooLong && hasNoOperational && hasValidEnd {
			var newS string
			if len(s) > nums[0] {
				newS = s[nums[0]+1:]
			}
			var newNums []int
			if len(nums) > 1 {
				newNums = nums[1:]
			}
			total += count(newS, newNums, cache, cacheMu)
		}
	}
	cacheMu.Lock()
	defer cacheMu.Unlock()
	cache[key] = total
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Global cache to minimize repeat lookup costs (this ended up making a
	// massive difference).
	cache := make(map[string]int)
	var cacheMu sync.RWMutex

	t1 := 0
	for _, line := range lines {
		s, nums := parse(line)
		t1 += count(s, nums, cache, &cacheMu)
	}
	// Answer: 7670.
	fmt.Println("Solution 1:", t1)

	t2 := 0
	for _, line := range lines {
		s, nums := parse(line)
		s, nums = unfold(s, nums)
		t2 += count(s, nums, cache, &cacheMu)
	}
	// Parallelize the work. Ended up not using this as it made essentially no
	// difference in the performance, whereas the cache made *all* the difference.
	//
	// Spawn workers.
	// var wg sync.WaitGroup
	// numCPUs := 4
	// work := make(chan string, numCPUs)
	// for i := 0; i <= cap(work); i++ {
	// 	go func(w chan string, wg *sync.WaitGroup) {
	// 		for line := range w {
	// 			s, nums := parse(line)
	// 			s, nums = unfold(s, nums)
	// 			t2 += count(s, nums, cache, &cacheMu)
	// 			wg.Done()
	// 		}
	// 	}(work, &wg)
	// }
	// // Send work down the work channel.
	// for _, line := range lines {
	// 	wg.Add(1)
	// 	work <- line
	// }
	// wg.Wait()
	// close(work)

	// Answer: 157383940585037.
	fmt.Println("Solution 2:", t2)
}
