package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	total int
	nums  []int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	entries := []entry{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		e := entry{}
		vals := strings.Split(scanner.Text(), ":")
		total, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal(err)
		}
		e.total = total
		nums := []int{}
		numStrs := strings.Fields(vals[1])
		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		e.nums = nums

		entries = append(entries, e)
	}

	cache := make(map[string]int)

	total1 := 0
	for _, e := range entries {
		// Cache keys generated for this entry.
		keys := []string{}
		for idx, num := range e.nums {
			if idx == 0 {
				key := fmt.Sprint(num)
				cache[key] = num
				keys = []string{key}
				continue
			}
			keysThisLoop := []string{}
			for _, key := range keys {
				plusKey := fmt.Sprintf("%v+%v", key, num)
				plusVal := cache[key] + num
				multKey := fmt.Sprintf("%v*%v", key, num)
				multVal := cache[key] * num
				cache[plusKey] = plusVal
				cache[multKey] = multVal
				keysThisLoop = append(keysThisLoop, plusKey)
				keysThisLoop = append(keysThisLoop, multKey)
				if idx < len(e.nums)-1 {
					continue
				}
				if plusVal == e.total || multVal == e.total {
					total1 += e.total
					break
				}
			}
			keys = keysThisLoop
		}
	}
	// Answer: 303876485655.
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for _, e := range entries {
		// Cache keys generated for this entry.
		keys := []string{}
		for idx, num := range e.nums {
			if idx == 0 {
				key := fmt.Sprint(num)
				cache[key] = num
				keys = []string{key}
				continue
			}
			keysThisLoop := []string{}
			for _, key := range keys {
				plusKey := fmt.Sprintf("%v+%v", key, num)
				plusVal := cache[key] + num
				multKey := fmt.Sprintf("%v*%v", key, num)
				multVal := cache[key] * num
				concatKey := fmt.Sprintf("%v||%v", key, num)
				concatVal, err := strconv.Atoi(fmt.Sprintf("%v%v", cache[key], num))
				if err != nil {
					log.Fatal(err)
				}
				cache[plusKey] = plusVal
				cache[multKey] = multVal
				cache[concatKey] = concatVal
				keysThisLoop = append(keysThisLoop, plusKey)
				keysThisLoop = append(keysThisLoop, multKey)
				keysThisLoop = append(keysThisLoop, concatKey)
				if idx < len(e.nums)-1 {
					continue
				}
				if plusVal == e.total || multVal == e.total || concatVal == e.total {
					total2 += e.total
					break
				}
			}
			keys = keysThisLoop
		}
	}
	// Answer: 146111650210682.
	fmt.Println("Solution 2:", total2)
}
