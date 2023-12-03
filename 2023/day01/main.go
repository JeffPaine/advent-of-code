package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func digitToValue(s string) int {
	var digits = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	for text, val := range digits {
		if strings.HasPrefix(s, text) {
			return val
		}
	}
	return -1
}

func firstLastNums(s string) int {
	firstNum := 0
	lastNum := 0
	first := true

	for idx, char := range s {
		digit := -1

		// Check for a string number, e.g. "one".
		if d := digitToValue(s[idx:]); d != -1 {
			digit = d
		}

		// Check for a digit, e.g. "1".
		d, err := strconv.Atoi(string(char))
		if err == nil { // No error.
			digit = d
		}

		if digit < 1 {
			continue
		}

		if first {
			firstNum = digit
			first = false
		}
		lastNum = digit
	}

	return firstNum*10 + lastNum
}

func firstLastDigit(s string) int {
	first := true
	firstNum := 0
	lastNum := 0
	for _, char := range s {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		if first {
			firstNum = digit
			first = false
		}
		lastNum = digit
	}
	return firstNum*10 + lastNum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0
	for _, line := range lines {
		total += firstLastDigit(line)
	}
	fmt.Println("Solution 1:", total)

	total = 0
	for _, line := range lines {
		total += firstLastNums(line)
	}
	fmt.Println("Solution 2:", total)
}
