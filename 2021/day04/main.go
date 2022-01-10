package main

import (
	"bufio"
	"log"
	"os"

	"github.com/JeffPaine/advent-of-code/2021/advent"
)

func solution1(nums []int, boards []advent.Board) int {
	for _, num := range nums {
		for _, board := range boards {
			board.MarkSpot(num)
			if board.HasBingo() {
				return board.Score(num)
			}
		}
	}
	return -1
}

func solution2(nums []int, boards []advent.Board) int {
	lastNum := 0
	lastBoard := advent.Board{}
	for _, num := range nums {
		for i := 0; i < len(boards); i++ {
			// As soon as a board has bingo, mark it as the most
			// recent winner and then stop checking it on future
			// iterations. After iterating like this over all
			// boards, print the score of the most recent winner.
			boards[i].MarkSpot(num)
			if boards[i].HasBingo() {
				lastNum = num
				lastBoard = boards[i]
				// Delete board from the list of boards checked.
				boards = append(boards[:i], boards[i+1:]...)
			}
		}
	}
	return lastBoard.Score(lastNum)
}

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
	nums, boards1 := advent.ParseBingoInput(input)
	// Slices in Go are passed by reference, so we can't use s1 = s2 as
	// changes to one will affect the other, we have to make a copy into a
	// slice of the same length.
	boards2 := make([]advent.Board, len(boards1))
	copy(boards2, boards1)

	log.Println("Solution1:", solution1(nums, boards1))
	log.Println("Solution2:", solution2(nums, boards2))
}
