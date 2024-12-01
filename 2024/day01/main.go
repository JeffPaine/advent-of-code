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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	left := []int{}
	right := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "   ")
		l, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, l)
		r, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := 0; i < len(left); i++ {
		l := left[i]
		r := right[i]
		switch {
		case l > r:
			total += (l - r)
		case r > l:
			total += (r - l)
		}
	}

	// Answer: 1970720.
	fmt.Println("Solution 1:", total)

	rcount := make(map[int]int)
	for _, num := range right {
		rcount[num] += 1
	}

	total2 := 0
	for _, num := range left {
		if count, ok := rcount[num]; ok {
			total2 += num * count
		}
	}

	// Answer: 17191599.
	fmt.Println("Solution 2:", total2)
}
