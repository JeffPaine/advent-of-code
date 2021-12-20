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
	commands := advent.ParseCommands(lines)
	p1 := advent.CalculatePosition(commands)
	log.Println("Solution 1:", p1.Horizontal*p1.Depth)
	p2 := advent.CalculatePositionWithAim(commands)
	log.Println("Solution 1:", p2.Horizontal*p2.Depth)
}
