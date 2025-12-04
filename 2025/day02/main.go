package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isTarget1(num int) bool {
	s := strconv.Itoa(num)
	// An odd number of digits means there can't be an even number of repeating digits.
	if len(s)%2 == 1 {
		return false
	}
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func isTarget2(num int) bool {
	digits := strconv.Itoa(num)
	combined := digits[1:] + digits[:len(digits)-1]
	return strings.Contains(combined, digits)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	ranges := ""

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ranges = scanner.Text()
	}

	total1 := 0
	total2 := 0

	for interval := range strings.SplitSeq(ranges, ",") {
		parts := strings.Split(interval, "-")
		first, second := parts[0], parts[1]
		start, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal("error converting string to integer:", err)
		}
		end, err := strconv.Atoi(second)
		if err != nil {
			log.Fatal("error converting string to integer:", err)
		}
		for num := start; num <= end; num++ {
			if isTarget1(num) {
				total1 += num
			}
			if isTarget2(num) {
				total2 += num
			}
		}
	}

	// Answer: 19605500130
	fmt.Println("Solution 1:", total1)

	// Answer: 36862281418
	fmt.Println("Solution 2:", total2)
}
