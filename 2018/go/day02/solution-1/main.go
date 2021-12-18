package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countTwoAndThree(s string) (bool, bool) {
	letters := make(map[rune]int)
	two := false
	three := false

	for _, char := range s {
		letters[char] += 1
	}

	for _, v := range letters {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}
	return two, three
}

func main() {
	totalTwo := 0
	totalThree := 0

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		two, three := countTwoAndThree(scanner.Text())
		if two {
			totalTwo += 1
		}
		if three {
			totalThree += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanner: ", err)
	}

	fmt.Println("two: ", totalTwo)
	fmt.Println("three: ", totalThree)
	fmt.Println("total: ", totalTwo*totalThree)
}
