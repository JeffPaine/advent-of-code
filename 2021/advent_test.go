package advent

import (
	"reflect"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7
	if got := CountIncreases(input); got != want {
		t.Errorf("CountIncreases(%v) = %v, want = %v", input, got, want)
	}
}

func TestRollingSumIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5
	if got := RollingSumIncreases(input); got != want {
		t.Errorf("RollingSumIncreases(%v) = %v, want %v", input, got, want)
	}
}

func TestParseCommands(t *testing.T) {
	input := []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	want := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	if got := ParseCommands(input); reflect.DeepEqual(got, want) != true {
		t.Errorf("ParseCommands(%v) = %v, want %v", input, got, want)
	}
}

func TestCalculatePosition(t *testing.T) {
	input := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	want := Position{Horizontal: 15, Depth: 10}
	if got := CalculatePosition(input); got != want {
		t.Errorf("CalculatePosition(%v) = %v, want %v", input, got, want)
	}
}

func TestCalculatePositionWithAim(t *testing.T) {
	input := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	want := Position{Horizontal: 15, Depth: 60, Aim: 10}
	if got := CalculatePositionWithAim(input); got != want {
		t.Errorf("CalculatePositionWithAim(%v) = %v, want %v", input, got, want)
	}
}

func TestNewReport(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	report := NewReport(input)
	want := 22
	if got := report.Gamma; got != want {
		t.Errorf("NewReport(%v).Gamma = %v, want %v", input, got, want)
	}
	want = 9
	if got := report.Epsilon; got != want {
		t.Errorf("NewReport(%v).Epsilon = %v, want %v", input, got, want)
	}
}

func TestLifeSupportRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	report := NewReport(input)
	want := 230
	if got := report.LifeSupportRating(); got != want {
		t.Errorf("LifeSupportRating() = %v, want %v", got, want)
	}
}

func TestBinaryStringToInt(t *testing.T) {
	input := "1011"
	want := 11
	if got := BinaryStringToInt(input); got != want {
		t.Errorf("BinaryStringToInt(%v) = %v, want %v", input, got, want)
	}
}

func TestParseBoardLine(t *testing.T) {
	input := "19  8  7 25 23"
	want := []Spot{{19, false}, {8, false}, {7, false}, {25, false}, {23, false}}
	if got := parseBoardLine(input); reflect.DeepEqual(got, want) != true {
		t.Errorf("parseBoardLine(%v) = %v, want %v", input, got, want)
	}
}

func TestParseBingoInput(t *testing.T) {
	input := []string{"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		// "",
		// " 3 15  0  2 22",
		// " 9 18 13 17  5",
		// "19  8  7 25 23",
		// "20 11 10 24  4",
		// "14 21 16 12  6",
		// "",
		// "14 21 17 24  4",
		// "10 16 15  9 19",
		// "18  8 23 26 20",
		// "22 11 13  6  5",
		// " 2  0 12  3  7",
	}
	movesWant := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	if got, _ := ParseBingoInput(input); reflect.DeepEqual(got, movesWant) != true {
		t.Errorf("ParseBingoInput(%v) = %v, want %v", input, got, movesWant)
	}
	boardWant := []Board{{
		rows: [][]Spot{
			{
				{num: 22, marked: false},
				{num: 13, marked: false},
				{num: 17, marked: false},
				{num: 11, marked: false},
				{num: 0, marked: false},
			},
			{
				{num: 8, marked: false},
				{num: 2, marked: false},
				{num: 23, marked: false},
				{num: 4, marked: false},
				{num: 24, marked: false},
			},
			{
				{num: 21, marked: false},
				{num: 9, marked: false},
				{num: 14, marked: false},
				{num: 16, marked: false},
				{num: 7, marked: false},
			},
			{
				{num: 6, marked: false},
				{num: 10, marked: false},
				{num: 3, marked: false},
				{num: 18, marked: false},
				{num: 5, marked: false},
			},
			{
				{num: 1, marked: false},
				{num: 12, marked: false},
				{num: 20, marked: false},
				{num: 15, marked: false},
				{num: 19, marked: false},
			},
		},
	}}
	if _, got := ParseBingoInput(input); reflect.DeepEqual(got, boardWant) != true {
		t.Errorf("ParseBingoInput(%v) = %v, want %v", input, got, boardWant)
	}
}

func TestHasBingoRow(t *testing.T) {
	board := Board{
		rows: [][]Spot{
			{
				{num: 22, marked: false},
				{num: 13, marked: false},
				{num: 17, marked: false},
				{num: 11, marked: false},
				{num: 0, marked: false},
			},
			{
				{num: 8, marked: false},
				{num: 2, marked: false},
				{num: 23, marked: false},
				{num: 4, marked: false},
				{num: 24, marked: false},
			},
			{
				{num: 21, marked: false},
				{num: 9, marked: false},
				{num: 14, marked: false},
				{num: 16, marked: false},
				{num: 7, marked: false},
			},
			{
				{num: 6, marked: false},
				{num: 10, marked: false},
				{num: 3, marked: false},
				{num: 18, marked: false},
				{num: 5, marked: false},
			},
			{
				{num: 1, marked: false},
				{num: 12, marked: false},
				{num: 20, marked: false},
				{num: 15, marked: false},
				{num: 19, marked: false},
			},
		},
	}
	// One short of a winning row.
	nums := []int{22, 13, 17, 11}
	for _, num := range nums {
		board.MarkSpot(num)
	}
	want := false
	if got := board.HasBingo(); got != want {
		t.Errorf("board.HasBing() = %v, want %v", got, want)
	}

	// Complete the row.
	board.MarkSpot(0)
	want = true
	if got := board.HasBingo(); got != want {
		t.Errorf("board.HasBing() = %v, want %v", got, want)
	}
}

func TestHasBingoColumn(t *testing.T) {
	board := Board{
		rows: [][]Spot{
			{
				{num: 22, marked: false},
				{num: 13, marked: false},
				{num: 17, marked: false},
				{num: 11, marked: false},
				{num: 0, marked: false},
			},
			{
				{num: 8, marked: false},
				{num: 2, marked: false},
				{num: 23, marked: false},
				{num: 4, marked: false},
				{num: 24, marked: false},
			},
			{
				{num: 21, marked: false},
				{num: 9, marked: false},
				{num: 14, marked: false},
				{num: 16, marked: false},
				{num: 7, marked: false},
			},
			{
				{num: 6, marked: false},
				{num: 10, marked: false},
				{num: 3, marked: false},
				{num: 18, marked: false},
				{num: 5, marked: false},
			},
			{
				{num: 1, marked: false},
				{num: 12, marked: false},
				{num: 20, marked: false},
				{num: 15, marked: false},
				{num: 19, marked: false},
			},
		},
	}
	// One short of a winning column.
	nums := []int{22, 8, 21, 6}
	for _, num := range nums {
		board.MarkSpot(num)
	}
	want := false
	if got := board.HasBingo(); got != want {
		t.Errorf("board.HasBing() = %v, want %v", got, want)
	}

	// Complete the column.
	board.MarkSpot(1)
	want = true
	if got := board.HasBingo(); got != want {
		t.Errorf("board.HasBing() = %v, want %v", got, want)
	}
}

func TestScore(t *testing.T) {
	board := Board{
		rows: [][]Spot{
			{
				{num: 22, marked: false},
				{num: 13, marked: false},
				{num: 17, marked: false},
				{num: 11, marked: false},
				{num: 0, marked: false},
			},
			{
				{num: 8, marked: false},
				{num: 2, marked: false},
				{num: 23, marked: false},
				{num: 4, marked: false},
				{num: 24, marked: false},
			},
			{
				{num: 21, marked: false},
				{num: 9, marked: false},
				{num: 14, marked: false},
				{num: 16, marked: false},
				{num: 7, marked: false},
			},
			{
				{num: 6, marked: false},
				{num: 10, marked: false},
				{num: 3, marked: false},
				{num: 18, marked: false},
				{num: 5, marked: false},
			},
			{
				{num: 1, marked: false},
				{num: 12, marked: false},
				{num: 20, marked: false},
				{num: 15, marked: false},
				{num: 19, marked: false},
			},
		},
	}
	nums := []int{22, 8, 21, 6, 1}
	last := 1
	for _, num := range nums {
		board.MarkSpot(num)
	}
	want := 242
	if got := board.Score(last); got != want {
		t.Errorf("board.Score(%v) = %v, want %v", last, got, want)
	}
}
