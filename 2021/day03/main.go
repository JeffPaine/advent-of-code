package main

import (
	"bufio"
	"log"
	"os"

	"github.com/JeffPaine/advent-of-code/2021/advent"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	report := advent.NewReport(lines)
	log.Println("Solution 1:", report.Consumption())
	log.Println("Solution 2:", report.LifeSupportRating())
}
