package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func opValue(op string) int {
	op = strings.TrimPrefix(op, "mul(")
	op = strings.TrimSuffix(op, ")")
	vals := strings.Split(op, ",")
	a, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}
	return a * b
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	operations := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		operations = append(operations, r.FindAllString(scanner.Text(), -1)...)
	}

	total1 := 0
	for _, op := range operations {
		if op == "do()" || op == "don't()" {
			continue
		}
		total1 += opValue(op)
	}

	// Answer: 159892596.
	fmt.Println("Solution 1:", total1)

	total2 := 0
	enabled := true
	for _, op := range operations {
		if op == "do()" {
			enabled = true
			continue
		}
		if op == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		total2 += opValue(op)
	}

	// Answer: 92626942.
	fmt.Println("Solution 2:", total2)
}
