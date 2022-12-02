package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type play int

const (
	rock play = iota
	paper
	scissors
)

func (p play) String() string {
	return []string{"rock", "paper", "scissors"}[p]
}

func newPlay(s string) (play, error) {
	switch s {
	case "A":
		return rock, nil
	case "B":
		return paper, nil
	case "C":
		return scissors, nil
	case "X":
		return rock, nil
	case "Y":
		return paper, nil
	case "Z":
		return scissors, nil
	}
	return -1, fmt.Errorf("unsupported input: newPlay(%q)", s)
}

type outcome int

const (
	win outcome = iota
	lose
	draw
)

func (o outcome) String() string {
	return []string{"win", "lose", "draw"}[o]
}

type game struct {
	mine   play
	theirs play
}

func (g game) String() string {
	return fmt.Sprintf("game{%v, %v}", g.mine, g.theirs)
}

// "Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the
// round instead ends in a draw."
func (g game) outcome() outcome {
	if g.mine == g.theirs {
		return draw
	}
	if (g.mine == rock && g.theirs == scissors) || (g.mine == scissors && g.theirs == paper) || (g.mine == paper && g.theirs == rock) {
		return win
	}
	return lose
}

// "Your total score is the sum of your scores for each round. The score for a single round is the score for the shape
// you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you
// lost, 3 if the round was a draw, and 6 if you won)."
func (g game) score() int {
	total := 0

	switch g.mine {
	case rock:
		total += 1
	case paper:
		total += 2
	case scissors:
		total += 3
	}

	switch g.outcome() {
	case win:
		total += 6
	case lose:
		total += 0
	case draw:
		total += 3
	}

	return total
}

func newGame(mine, theirs string) (game, error) {
	m, err := newPlay(mine)
	if err != nil {
		return game{}, fmt.Errorf("newPlay(%q): %v", mine, err)
	}
	t, err := newPlay(theirs)
	if err != nil {
		return game{}, fmt.Errorf("newPlay(%q): %v", theirs, err)
	}
	return game{mine: m, theirs: t}, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	games := []game{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var mine string
		var theirs string
		_, err := fmt.Sscanf(scanner.Text(), "%s %s", &theirs, &mine)
		if err != nil {
			log.Fatalf("fmt.Sscanf(%q, \"%%q %%q\", &theirs, &mine): %v", scanner.Text(), err)
		}
		g, err := newGame(mine, theirs)
		if err != nil {
			log.Fatalf("newGame(%q, %q): %v", mine, theirs, err)
		}
		games = append(games, g)
	}

	total := 0
	for _, g := range games {
		total += g.score()
	}
	log.Println("Solution 1:", total)
}
