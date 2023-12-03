package main

import (
	"reflect"
	"testing"
)

func TestParseGame(t *testing.T) {
	tests := []struct {
		s       string
		want    game
		wantErr bool
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			game{id: 1, red: 4, green: 2, blue: 6},
			false,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			game{id: 2, red: 1, green: 3, blue: 4},
			false,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			game{id: 3, red: 20, green: 13, blue: 6},
			false,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			game{id: 4, red: 14, green: 3, blue: 15},
			false,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			game{id: 5, red: 6, green: 3, blue: 2},
			false,
		},
	}

	for _, test := range tests {
		got, err := parseGame(test.s)
		if test.wantErr && err != nil {
			t.Errorf("parseGame(%q) = %+v, %v; want: %+v, %v", test.s, got, err, test.want, nil)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseGame(%q) = %+v, %v; want: %+v, %v", test.s, got, err, test.want, nil)
		}
	}
}
