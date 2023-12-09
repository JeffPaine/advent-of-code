package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	s    string
	nums []int
}

func parseLine(s string) line {
	l := line{s: s}
	ints := strings.Fields(l.s)
	var nums []int
	for _, i := range ints {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic("bad num")
		}
		nums = append(nums, num)
	}
	l.nums = nums
	return l
}

func predictForward(ints []int) int {
	if allZeroes(ints) {
		return 0
	}
	var diff []int
	for i := 1; i < len(ints); i++ {
		diff = append(diff, ints[i]-ints[i-1])
	}
	return ints[len(ints)-1] + predictForward(diff)
}

func predictBackward(ints []int) int {
	if allZeroes(ints) {
		return 0
	}
	var diff []int
	for i := 1; i < len(ints); i++ {
		diff = append(diff, ints[i]-ints[i-1])
	}
	return ints[0] - predictBackward(diff)
}

func allZeroes(ints []int) bool {
	total := 0
	for _, i := range ints {
		if i == 0 {
			total++
		}
	}
	return total == len(ints)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(f)

	var lines []line
	for scanner.Scan() {
		lines = append(lines, parseLine(scanner.Text()))
	}

	t1 := 0
	for _, l := range lines {
		t1 += predictForward(l.nums)
	}
	// 1647269739.
	fmt.Println("Solution 1:", t1)

	t2 := 0
	for _, l := range lines {
		t2 += predictBackward(l.nums)
	}
	// 864.
	fmt.Println("Solution 2:", t2)
}
