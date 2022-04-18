package main

import (
	"log"
	"os"

	"github.com/JeffPaine/advent-of-code/2021/advent"
)

func main() {
	// Ignoring diagonal lines.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	grid := advent.NewGrid(f, false)
	log.Println("Solution1:", grid.AtLeastTwo())
	f.Close()

	// Including diagonal lines.
	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	grid = advent.NewGrid(f, true)
	log.Println("Solution2:", grid.AtLeastTwo())
}
