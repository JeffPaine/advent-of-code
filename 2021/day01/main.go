package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/JeffPaine/advent-of-code/2021/advent"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	nums := []int{}
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, val)
	}
	log.Println("Solution 1:", advent.CountIncreases(nums))
	log.Println("Solution 2:", advent.RollingSumIncreases(nums))
}
