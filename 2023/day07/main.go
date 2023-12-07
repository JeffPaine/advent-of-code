package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// knd is a card hand kind, e.g. "full house", etc.
type knd int

const (
	highCard knd = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func (k knd) String() string {
	return []string{
		"highCard",
		"onePair",
		"twoPair",
		"threeOfAKind",
		"fullHouse",
		"fourOfAKind",
		"fiveOfAKind",
	}[k]
}

// crd is a type of card, e.g. "ten", "jack", etc.
type crd int

// // Part 1.
// const (
// 	two crd = iota
// 	three
// 	four
// 	five
// 	six
// 	seven
// 	eight
// 	nine
// 	ten
// 	jack
// 	queen
// 	king
// 	ace
// )

// // Part 1.
// func (c crd) String() string {
// 	return []string{
// 		"2",
// 		"3",
// 		"4",
// 		"5",
// 		"6",
// 		"7",
// 		"8",
// 		"9",
// 		"T",
// 		"J",
// 		"Q",
// 		"K",
// 		"A",
// 	}[c]
// }

// Part 1.
// func cardVal(r rune) crd {
// 	switch r {
// 	case '2':
// 		return two
// 	case '3':
// 		return three
// 	case '4':
// 		return four
// 	case '5':
// 		return five
// 	case '6':
// 		return six
// 	case '7':
// 		return seven
// 	case '8':
// 		return eight
// 	case '9':
// 		return nine
// 	case 'T':
// 		return ten
// 	case 'J':
// 		return jack
// 	case 'Q':
// 		return queen
// 	case 'K':
// 		return king
// 	case 'A':
// 		return ace
// 	default:
// 		panic("unsupported card value")
// 	}
// }

// Part 2.
const (
	joker crd = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	queen
	king
	ace
)

// Part 2.
func (c crd) String() string {
	return []string{
		"J",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"Q",
		"K",
		"A",
	}[c]
}

// Part 2.
func cardVal(r rune) crd {
	switch r {
	case '2':
		return two
	case '3':
		return three
	case '4':
		return four
	case '5':
		return five
	case '6':
		return six
	case '7':
		return seven
	case '8':
		return eight
	case '9':
		return nine
	case 'T':
		return ten
	case 'J':
		return joker
	case 'Q':
		return queen
	case 'K':
		return king
	case 'A':
		return ace
	default:
		panic("unsupported card value")
	}
}

type hand struct {
	cards []crd
	bid   int
	kind  knd
}

func (h hand) String() string {
	var cards string
	for _, c := range h.cards {
		cards += fmt.Sprintf("%v", c)
	}
	return fmt.Sprintf("{cards: %v, bid: %v, kind: %v}", cards, h.bid, h.kind)
}

func newHand(s string) hand {
	var h hand
	fields := strings.Fields(s)
	h.bid, _ = strconv.Atoi(fields[1])
	for _, val := range fields[0] {
		h.cards = append(h.cards, cardVal(val))
	}
	h.score()
	return h
}

// // Part 1.
// func (h *hand) score() {
// 	// A hand is exactly one kind.
// 	m := make(map[crd]int)
// 	for _, c := range h.cards {
// 		m[c] += 1
// 	}
// 	var counts []int
// 	for _, val := range m {
// 		counts = append(counts, val)
// 	}
// 	// Sort counts in ascending order, e.g. [1, 2, 3].
// 	sort.Ints(counts)
// 	maxCount := counts[len(counts)-1]
// 	if maxCount == 5 {
// 		h.kind = fiveOfAKind
// 		return
// 	}
// 	if maxCount == 4 {
// 		h.kind = fourOfAKind
// 		return
// 	}
// 	// Full house, where three cards have the same label, and the remaining two cards share a
// 	// different label: 23332
// 	if len(counts) == 2 && counts[0] == 2 && counts[1] == 3 {
// 		h.kind = fullHouse
// 		return
// 	}
// 	if len(counts) == 3 && maxCount == 3 {
// 		h.kind = threeOfAKind
// 		return
// 	}
// 	if counts[0] == 1 && counts[1] == 2 && counts[2] == 2 {
// 		h.kind = twoPair
// 		return
// 	}
// 	if counts[0] == 1 && counts[1] == 1 && counts[2] == 1 && counts[3] == 2 {
// 		h.kind = onePair
// 		return
// 	}
// 	h.kind = highCard
// }

func (h *hand) score() {
	m := make(map[crd]int)
	for _, c := range h.cards {
		m[c] += 1
	}

	jokerCount := m[joker]
	if jokerCount == 5 {
		h.kind = fiveOfAKind
		return
	}
	// Ignore joker counts.
	delete(m, joker)

	var counts []int
	for _, val := range m {
		counts = append(counts, val)
	}
	// Sort counts in ascending order, e.g. [1, 2, 3].
	sort.Ints(counts)
	maxCount := counts[len(counts)-1]

	// FiveOfAKind.
	if maxCount == 5 {
		h.kind = fiveOfAKind
		return
	}
	if maxCount == 4 && jokerCount == 1 {
		h.kind = fiveOfAKind
		return
	}
	if maxCount == 3 && jokerCount == 2 {
		h.kind = fiveOfAKind
		return
	}
	if maxCount == 2 && jokerCount == 3 {
		h.kind = fiveOfAKind
		return
	}
	if maxCount == 1 && jokerCount == 4 {
		h.kind = fiveOfAKind
		return
	}

	// FourOfAKind.
	if maxCount == 4 {
		h.kind = fourOfAKind
		return
	}
	if maxCount == 3 && jokerCount == 1 {
		h.kind = fourOfAKind
		return
	}
	if maxCount == 2 && jokerCount == 2 {
		h.kind = fourOfAKind
		return
	}
	if maxCount == 1 && jokerCount == 3 {
		h.kind = fourOfAKind
		return
	}

	// fullHouse.
	// Full house, where three cards have the same label, and the remaining two cards share a
	// different label: 23332
	if len(counts) == 2 && counts[0] == 2 && counts[1] == 3 {
		h.kind = fullHouse
		return
	}
	if len(counts) == 2 && counts[0] == 1 && counts[1] == 3 && jokerCount == 1 {
		h.kind = fullHouse
		return
	}
	if len(counts) == 2 && counts[0] == 2 && counts[1] == 2 && jokerCount == 1 {
		h.kind = fullHouse
		return
	}

	// threeOfAKind.
	if len(counts) == 3 && maxCount == 3 {
		h.kind = threeOfAKind
		return
	}
	if len(counts) == 3 && maxCount == 2 && jokerCount == 1 {
		h.kind = threeOfAKind
		return
	}
	if len(counts) == 3 && maxCount == 1 && jokerCount == 2 {
		h.kind = threeOfAKind
		return
	}

	// twoPair.
	if len(counts) == 3 && counts[0] == 1 && counts[1] == 2 && counts[2] == 2 {
		h.kind = twoPair
		return
	}
	if len(counts) == 2 && counts[0] == 1 && counts[1] == 2 && jokerCount == 1 {
		h.kind = twoPair
		return
	}

	// onePair.
	if len(counts) == 4 && counts[0] == 1 && counts[1] == 1 && counts[2] == 1 && counts[3] == 2 {
		h.kind = onePair
		return
	}
	if len(counts) == 3 && counts[0] == 1 && counts[1] == 1 && counts[2] == 1 && jokerCount == 1 {
		h.kind = onePair
		return
	}
	if len(counts) == 4 && jokerCount == 1 {
		h.kind = onePair
		return
	}
	h.kind = highCard
}

func sortHands(hs []hand) []hand {
	sort.Slice(hs, func(i, j int) bool {
		if hs[i].kind != hs[j].kind {
			return hs[i].kind < hs[j].kind
		}
		// Hands have the same kind, sort by individual cards.
		for k := 0; k < len(hs[i].cards); k++ {
			if hs[i].cards[k] == hs[j].cards[k] {
				continue
			}
			return hs[i].cards[k] < hs[j].cards[k]
		}
		panic("unexpected sorting condition")
	})
	return hs
}

func main() {
	// example := `32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483`
	// r := strings.NewReader(example)
	// scanner := bufio.NewScanner(r)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(f)

	var hands []hand
	for scanner.Scan() {
		h := newHand(scanner.Text())
		hands = append(hands, h)
	}
	sortHands(hands)
	total := 0
	for idx, h := range hands {
		val := (idx + 1) * h.bid
		total += val
	}
	fmt.Println("total:", total)
	// Part 1: 250946742
	// Part 2: 251824095
}
