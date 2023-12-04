package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type card struct {
	id      int
	winners []int
	have    []int
}

func (c card) matches() int {
	matches := 0
	winners := make(map[int]struct{})
	for _, val := range c.winners {
		winners[val] = struct{}{}
	}
	for _, val := range c.have {
		if _, ok := winners[val]; ok {
			matches += 1
		}
	}
	return matches
}

func (c card) points() int {
	return int(math.Pow(2, float64(c.matches()-1)))
}

func parseNums(s string) ([]int, error) {
	var out []int
	for _, val := range strings.Fields(s) {
		i, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, fmt.Errorf("could not convert to int: %v", val)
		}
		out = append(out, i)
	}
	return out, nil
}

func parseCard(s string) (card, error) {
	var c card
	first := strings.Split(s, ":")
	_, err := fmt.Sscanf(first[0], "Card %d", &c.id)
	if err != nil {
		return card{}, fmt.Errorf("could not parse %q into \"Card %%d\", error: %v", first[0], err)
	}
	second := strings.Split(first[1], " | ")
	winners, err := parseNums(second[0])
	if err != nil {
		return card{}, err
	}
	c.winners = winners
	have, err := parseNums(second[1])
	if err != nil {
		return card{}, err
	}
	c.have = have

	return c, err
}

func cardsTotal(cards []card) int {
	// Iterate over the cards backwards, calculating the value of each card and its descendants.
	total := 0
	// Map of card index -> its count.
	counts := make(map[int]int)
	for idx := len(cards) - 1; idx >= 0; idx-- {
		cardTotal := 1
		matches := cards[idx].matches()
		// For each of the next matches, get their cardTotal and add to the running total.
		for i := 1; i < matches+1; i++ {
			cardTotal += counts[idx+i]
		}
		counts[idx] = cardTotal
		total += cardTotal
	}
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var cards []card
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c, err := parseCard(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		cards = append(cards, c)
	}

	total := 0
	for _, card := range cards {
		total += card.points()
	}
	fmt.Println("Solution 1:", total)

	fmt.Println("Solution 2:", cardsTotal((cards)))
}
