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

	input := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	nums, boards := advent.ParseBingoInput(input)
	for _, num := range nums {
		for _, board := range boards {
			board.MarkSpot(num)
			if board.HasBingo() {
				log.Println("Solution 1:", board.Score(num))
				return
			}
		}
	}
}
