package advent

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Amount    int
}

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func CountIncreases(nums []int) int {
	increases := 0
	last := 0
	for i, num := range nums {
		if i == 0 {
			last = num
			continue
		}
		if num > last {
			increases += 1
		}
		last = num
	}
	return increases
}

func RollingSumIncreases(nums []int) int {
	increases := 0
	first := nums[0]
	second := nums[1]
	third := nums[2]
	last := first + second + third
	for i := 3; i < len(nums); i++ {
		first = second
		second = third
		third = nums[i]
		current := first + second + third
		if current > last {
			increases += 1
		}
		last = current
	}
	return increases
}

func ParseCommands(lines []string) []Command {
	commands := []Command{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		c := Command{Direction: parts[0], Amount: i}
		commands = append(commands, c)
	}
	return commands
}

func CalculatePosition(commands []Command) Position {
	p := Position{}
	for _, c := range commands {
		switch c.Direction {
		case "forward":
			p.Horizontal += c.Amount
		case "down":
			p.Depth += c.Amount
		case "up":
			p.Depth -= c.Amount
		}
	}
	return p
}

func CalculatePositionWithAim(commands []Command) Position {
	p := Position{}
	for _, c := range commands {
		switch c.Direction {
		case "forward":
			p.Horizontal += c.Amount
			p.Depth += p.Aim * c.Amount
		case "down":
			p.Aim += c.Amount
		case "up":
			p.Aim -= c.Amount
		}
	}
	return p
}

type Report struct {
	lines []string

	Gamma     int
	Epsilon   int
	O2Rating  int
	CO2Rating int
}

func NewReport(lines []string) Report {
	r := Report{lines: lines}
	r.calcGammaAndEpsilon()
	r.O2Rating = mostCommon(r.lines)
	r.CO2Rating = leastCommon(r.lines)
	return r
}

func (r *Report) calcGammaAndEpsilon() {
	// We assume all lines are the same width.
	width := len(r.lines[0])

	// Count the the number of 0s and 1s per column.
	zeroesPerColumn := perColumn('0', r.lines)
	onesPerColumn := perColumn('1', r.lines)

	// Determine the most and least common value (0 or 1) per column and
	// construct the gamma and epsilon values out of these.
	//
	// Epsilon is comprised of all the least frequent, per-column values.
	// Gamma is comprised of all the most frequent, per-column values.
	//
	// zeroesPerColumn := [5, 4, 4]
	// onesPerColumn   := [4, 5, 5]
	//
	// Results in:
	//
	// Epsilon := 0b100
	// Gamma   := 0b011
	//
	gamma := 0
	epsilon := 0
	for i := 0; i < width; i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if zeroesPerColumn[i] > onesPerColumn[i] {
			epsilon |= 1
		} else {
			gamma |= 1
		}
	}
	r.Gamma = gamma
	r.Epsilon = epsilon
}

func (r Report) Consumption() int {
	return r.Gamma * r.Epsilon
}

func (r Report) LifeSupportRating() int {
	return r.O2Rating * r.CO2Rating
}

// BinaryStringToInt converts a binary string to an int. For example: "1001" -> 9.
func BinaryStringToInt(s string) int {
	out := 0
	for _, val := range s {
		out = out << 1
		if string(val) == "1" {
			out |= 1
		}
	}
	return out
}

func filterLines(idx, val int, lines []string) []string {
	out := []string{}
	for _, line := range lines {
		if string(line[idx]) == fmt.Sprint(val) {
			out = append(out, line)
		}
	}
	return out
}

func perColumn(char rune, lines []string) []int {
	out := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, val := range line {
			if val == char {
				out[i] += 1
			}
		}
	}
	return out
}

func common(lines []string, inverse bool) int {
	width := len(lines[0])
	matches := lines
	for i := 0; i < width; i++ {
		zeroesPerColumn := perColumn('0', matches)
		onesPerColumn := perColumn('1', matches)

		val := 0
		if zeroesPerColumn[i] == onesPerColumn[i] || zeroesPerColumn[i] < onesPerColumn[i] {
			val = 1
		}
		if inverse {
			if val == 0 {
				val = 1
			} else {
				val = 0
			}
		}

		matches = filterLines(i, val, matches)

		if len(matches) == 1 {
			break
		}
	}
	return BinaryStringToInt(matches[0])

}

func mostCommon(lines []string) int {
	return common(lines, false)
}

func leastCommon(lines []string) int {
	return common(lines, true)
}

type Spot struct {
	num    int
	marked bool
}

type Board struct {
	rows [][]Spot
}

func (b *Board) MarkSpot(num int) {
	for i := range b.rows {
		for j := range b.rows[i] {
			if b.rows[i][j].num == num {
				b.rows[i][j].marked = true
			}
		}
	}
}

func (b Board) HasBingo() bool {
	size := len(b.rows[0])
	countPerRow := make([]int, size)
	countPerColumn := make([]int, size)
	for ir, row := range b.rows {
		for ic, s := range row {
			if s.marked {
				countPerRow[ir]++
				countPerColumn[ic]++
			}
		}
	}
	for _, count := range countPerRow {
		if count == size {
			return true
		}
	}
	for _, count := range countPerColumn {
		if count == size {
			return true
		}
	}
	return false
}

func (b Board) Score(lastCalled int) int {
	unmarked := 0
	for _, row := range b.rows {
		for _, s := range row {
			if !s.marked {
				unmarked += s.num
			}
		}
	}
	return unmarked * lastCalled
}

func ParseBingoInput(lines []string) ([]int, []Board) {
	moves := []int{}
	for _, val := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		moves = append(moves, num)
	}

	boards := []Board{}
	currBoard := Board{}

	for i := 1; i < len(lines); i++ {
		// End of current board (if any).
		if lines[i] == "" {
			if len(currBoard.rows) > 0 {
				boards = append(boards, currBoard)
				currBoard = Board{}
			}
			continue
		} else {
			currBoard.rows = append(currBoard.rows, parseBoardLine(lines[i]))
		}
	}
	if len(currBoard.rows) > 0 {
		boards = append(boards, currBoard)
	}
	return moves, boards
}

func parseBoardLine(line string) []Spot {
	spots := []Spot{}
	strs := strings.Fields(line)
	for _, val := range strs {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		spots = append(spots, Spot{num: num, marked: false})
	}
	return spots
}

type line struct {
	x1, y1, x2, y2 int
}

func maxVals(lines []line) (int, int) {
	maxX := 0
	maxY := 0
	for _, l := range lines {
		if l.x1 > maxX {
			maxX = l.x1
		}
		if l.x2 > maxX {
			maxX = l.x2
		}
		if l.y1 > maxY {
			maxY = l.y1
		}
		if l.y2 > maxY {
			maxY = l.y2
		}
	}
	return maxX, maxY
}

type Grid struct {
	rows [][]int
}

func NewGrid(r io.Reader) Grid {
	// Parse the input.
	// Example line: 0,9 -> 5,9
	var lines []line
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var l line
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &l.x1, &l.y1, &l.x2, &l.y2)
		if err != nil {
			log.Fatalln(err)
		}
		lines = append(lines, l)
	}

	// Initialize the empty rows.
	maxX, maxY := maxVals(lines)
	var rows [][]int
	for i := 0; i <= maxY; i++ {
		rows = append(rows, make([]int, maxX+1))
	}

	// Fill out rows with data from the parsed lines.
	for _, line := range lines {
		switch {
		// Line is horizontal.
		case line.y1 == line.y2:
			var row int = line.y1
			var start int = line.x1
			var finish int = line.x2
			if line.x2 < line.x1 {
				start, finish = finish, start
			}
			for i := start; i <= finish; i++ {
				rows[row][i] += 1
			}
		// Line is vertical.
		case line.x1 == line.x2:
			var column int = line.x1
			var start int = line.y1
			var finish int = line.y2
			if line.y2 < line.y1 {
				start, finish = finish, start
			}
			for i := start; i <= finish; i++ {
				rows[i][column] += 1
			}
		default:
			// For now, we simply ignore diagonal lines.

		}
	}

	return Grid{rows: rows}

}

// AtLeastTwo counts the number of points that have 2 or more overlapping lines.
func (g Grid) AtLeastTwo() int {
	sum := 0
	for _, row := range g.rows {
		for _, col := range row {
			if col > 1 {
				sum++
			}
		}
	}
	return sum
}

func (g Grid) String() string {
	var out strings.Builder
	fmt.Fprintf(&out, "rows (rows: %v x columns: %v):\n", len(g.rows), len(g.rows[0]))
	for _, row := range g.rows {
		fmt.Fprintf(&out, "  %v\n", row)
	}
	return out.String()
}
