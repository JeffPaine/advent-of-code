package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maxJoltage(bank string, targetLen int) int {
	// A monotonically decreasing stack.
	var nums []int
	for i, digit := range bank {
		// Assume input runes are always valid numbers.
		num, _ := strconv.Atoi(string(digit))
		for len(nums) > 0 && nums[len(nums)-1] < num && targetLen-len(nums) <= len(bank)-1-i {
			nums = nums[:len(nums)-1]
		}
		nums = append(nums, num)
	}
	var sb strings.Builder
	for _, digit := range nums[:targetLen] {
		sb.WriteString(strconv.Itoa(digit))
	}
	ans, _ := strconv.Atoi(sb.String())
	return ans
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	var banks []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		banks = append(banks, scanner.Text())
	}

	total1 := 0
	for _, bank := range banks {
		total1 += maxJoltage(bank, 2)
	}

	// Answer: 17229
	fmt.Println("Solution 1:", total1)

	total2 := 0
	for _, bank := range banks {
		total2 += maxJoltage(bank, 12)
	}

	// Answer: 170520923035051
	fmt.Println("Solution 2:", total2)
}
