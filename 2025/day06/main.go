package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	var lines [][]string
	var rawLines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.Fields(scanner.Text()))
		rawLines = append(rawLines, scanner.Text())
	}

	total1 := 0
	for col := 0; col < len(lines[0]); col++ {
		sign := lines[len(lines)-1][col]
		total := 0
		if sign == "*" {
			total = 1
			for row := 0; row < len(lines)-1; row++ {
				num, _ := strconv.Atoi(lines[row][col])
				total *= num
			}
		} else {
			for row := 0; row < len(lines)-1; row++ {
				num, _ := strconv.Atoi(lines[row][col])
				total += num
			}
		}
		total1 += total
	}

	// Answer: 5784380717354
	fmt.Println("Solution 1:", total1)

	total2 := 0
	var nums []int
	for col := len(rawLines[0]) - 1; col >= 0; col-- {
		digits := ""
		for row := range len(rawLines) {
			char := rawLines[row][col]
			if char == ' ' {
				continue
			}
			// Only ASCII expected.
			if char >= '0' && char <= '9' {
				digits += string(char)
				continue
			}
			// Only chars expected fron now on are + or *, so add the current
			// digits to nums.
			num, _ := strconv.Atoi(digits)
			nums = append(nums, num)
			if char == '+' {
				for _, num := range nums {
					total2 += num
				}
			}
			if char == '*' {
				total := 1
				for _, num := range nums {
					total *= num
				}
				total2 += total
			}
			nums = []int{}
			digits = ""
		}
		// Hit the end of a column; convert digits, if present, to a number.
		if len(digits) > 0 {
			num, _ := strconv.Atoi(digits)
			nums = append(nums, num)
			digits = ""
		}
	}

	// Answer: 7996218225744
	fmt.Println("Solution 2:", total2)
}
