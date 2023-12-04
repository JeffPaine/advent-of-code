package main

import (
	"reflect"
	"testing"
)

func TestParseCard(t *testing.T) {
	tests := []struct {
		s    string
		want card
	}{
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			card{id: 1, winners: []int{41, 48, 83, 86, 17}, have: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			card{id: 2, winners: []int{13, 32, 20, 16, 61}, have: []int{61, 30, 68, 82, 17, 32, 24, 19}},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			card{id: 3, winners: []int{1, 21, 53, 59, 44}, have: []int{69, 82, 63, 72, 16, 21, 14, 1}},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			card{id: 4, winners: []int{41, 92, 73, 84, 69}, have: []int{59, 84, 76, 51, 58, 5, 54, 83}},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			card{id: 5, winners: []int{87, 83, 26, 28, 32}, have: []int{88, 30, 70, 12, 93, 22, 82, 36}},
		},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			card{id: 6, winners: []int{31, 18, 13, 56, 72}, have: []int{74, 77, 10, 23, 35, 67, 36, 11}},
		},
	}

	for _, test := range tests {
		got, err := parseCard(test.s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseCard(%q) = %v; want: %v", test.s, got, test.want)
		}
	}
}

func TestCardPoints(t *testing.T) {
	tests := []struct {
		c    card
		want int
	}{
		{
			card{id: 1, winners: []int{41, 48, 83, 86, 17}, have: []int{83, 86, 6, 31, 17, 9, 48, 53}},
			8,
		},
		{
			card{id: 2, winners: []int{13, 32, 20, 16, 61}, have: []int{61, 30, 68, 82, 17, 32, 24, 19}},
			2,
		},
		{
			card{id: 3, winners: []int{1, 21, 53, 59, 44}, have: []int{69, 82, 63, 72, 16, 21, 14, 1}},
			2,
		},
		{
			card{id: 4, winners: []int{41, 92, 73, 84, 69}, have: []int{59, 84, 76, 51, 58, 5, 54, 83}},
			1,
		},
		{
			card{id: 5, winners: []int{87, 83, 26, 28, 32}, have: []int{88, 30, 70, 12, 93, 22, 82, 36}},
			0,
		},
		{
			card{id: 6, winners: []int{31, 18, 13, 56, 72}, have: []int{74, 77, 10, 23, 35, 67, 36, 11}},
			0,
		},
	}

	for _, test := range tests {
		if got := test.c.points(); got != test.want {
			t.Errorf("c.points() = %v; want: %v", got, test.want)
		}
	}
}

func TestCardsTotal(t *testing.T) {
	cards := []card{
		{id: 1, winners: []int{41, 48, 83, 86, 17}, have: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{id: 2, winners: []int{13, 32, 20, 16, 61}, have: []int{61, 30, 68, 82, 17, 32, 24, 19}},
		{id: 3, winners: []int{1, 21, 53, 59, 44}, have: []int{69, 82, 63, 72, 16, 21, 14, 1}},
		{id: 4, winners: []int{41, 92, 73, 84, 69}, have: []int{59, 84, 76, 51, 58, 5, 54, 83}},
		{id: 5, winners: []int{87, 83, 26, 28, 32}, have: []int{88, 30, 70, 12, 93, 22, 82, 36}},
		{id: 6, winners: []int{31, 18, 13, 56, 72}, have: []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}
	want := 30
	got := cardsTotal(cards)
	if got != want {
		t.Errorf("cardsTotal(cards) = %v, want: %v", got, want)
	}
}
