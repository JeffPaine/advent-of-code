package main

import "testing"

func TestSumPartNumbers(t *testing.T) {
	// 467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..
	s := "467..114.....*........35..633.......#...617*...........+.58...592...........755....$.*.....664.598.."
	want := 4361
	if got, _ := sumPartNumbers(s, 10); got != want {
		t.Errorf("sumPartNumbers() = %v; want %v", got, want)
	}
}

func TestSumGearRatios(t *testing.T) {
	s := "467..114.....*........35..633.......#...617*...........+.58...592...........755....$.*.....664.598.."
	want := 467835
	_, parts := sumPartNumbers(s, 10)
	if got := sumGearRatios(s, 10, parts); got != want {
		t.Errorf("sumGearRatios() = %v; want %v", got, want)
	}

}
