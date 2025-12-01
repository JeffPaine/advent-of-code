package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type node struct {
	val  int
	next *node
	prev *node
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	zero := &node{val: 0}
	curr := zero
	for i := 1; i <= 100; i++ {
		next := &node{}
		if i == 100 {
			next = zero
		} else {
			next = &node{val: i}
		}
		curr.next = next
		next.prev = curr
		curr = next
	}

	// "Spin" the dial to 50.
	curr = zero
	for range 50 {
		curr = curr.next
	}

	total1 := 0
	total2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		sign := runes[0]
		num, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			log.Fatal("error converting to integer:", err)
		}
		for range num {
			if sign == 'L' {
				curr = curr.prev
			} else {
				curr = curr.next
			}
			if curr.val == 0 {
				total2++
			}
		}
		if curr.val == 0 {
			total1++
		}
	}

	// Answer: 1158
	fmt.Println("Solution 1:", total1)

	// Answer: 6860
	fmt.Println("Solution 2:", total2)
}
