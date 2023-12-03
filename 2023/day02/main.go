package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type game struct {
	id int

	// Max counts per color.
	red   int
	green int
	blue  int
}

func (g game) power() int {
	return g.red * g.green * g.blue
}

func parseGame(s string) (game, error) {
	// Example string:
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	var g game

	pieces := strings.Split(s, ":")
	_, err := fmt.Sscanf(pieces[0], "Game %d", &g.id)
	if err != nil {
		return game{}, err
	}
	sets := strings.Split(pieces[1], ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			var color string
			var count int
			_, err := fmt.Sscanf(cube, "%d %s", &count, &color)
			if err != nil {
				return game{}, fmt.Errorf("could not parse %q into (count, color)", cube)
			}
			switch color {
			case "red":
				if count > g.red {
					g.red = count
				}
			case "green":
				if count > g.green {
					g.green = count
				}
			case "blue":
				if count > g.blue {
					g.blue = count
				}
			}
		}
	}
	return g, nil
}

func possible(base, g game) bool {
	if base.red >= g.red && base.green >= g.green && base.blue >= g.blue {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var games []game
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game, err := parseGame(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		games = append(games, game)
	}

	total := 0
	base := game{id: -1, red: 12, green: 13, blue: 14}
	for _, game := range games {
		if possible(base, game) {
			total += game.id
		}
	}
	fmt.Println("Solution 1:", total)

	total = 0
	for _, game := range games {
		total += game.power()
	}
	fmt.Println("Solution 2:", total)
}
