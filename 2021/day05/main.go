package main

import (
	"log"
	"os"

	"github.com/JeffPaine/advent-of-code/2021/advent"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	grid := advent.NewGrid(f)
	log.Println("Solution1:", grid.AtLeastTwo())
}
