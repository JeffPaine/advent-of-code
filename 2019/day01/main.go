package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(f string) []int {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return nums
}

func calcFuel(i int) int {
	// "divide by three, round down to even integer, and subtract 2."
	return i/3 - 2
}

func one(nums []int) int {
	var total int

	for _, n := range nums {
		total += calcFuel(n)
	}

	return total
}

func two(nums []int) int {
	var total int

	for _, n := range nums {
		for currFuel := calcFuel(n); currFuel > 0; currFuel = calcFuel(currFuel) {
			total += currFuel
		}
	}
	return total
}

func main() {
	nums := readFile("input.txt")
	fmt.Println("total:", one(nums))
	fmt.Println("total:", two(nums))
}
