package advent

import (
	"fmt"
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
	r.calcO2AndCO2()
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

func (r *Report) calcO2AndCO2() {
	// We assume all lines are the same width.
	width := len(r.lines[0])

	// To find oxygen generator rating, determine the most common value (0
	// or 1) in the current bit position, and keep only numbers with that
	// bit in that position. If 0 and 1 are equally common, keep values
	// with a 1 in the position being considered.
	//
	// To find CO2 scrubber rating, determine the least common value (0 or
	// 1) in the current bit position, and keep only numbers with that bit
	// in that position. If 0 and 1 are equally common, keep values with a
	// 0 in the position being considered.

	o2Matches := r.lines
	for i := 0; i < width; i++ {
		zeroesPerColumn := perColumn('0', o2Matches)
		onesPerColumn := perColumn('1', o2Matches)

		// Oxygen values are the most common, CO2 are the least.
		o2Val := 0
		if zeroesPerColumn[i] == onesPerColumn[i] || zeroesPerColumn[i] < onesPerColumn[i] {
			o2Val = 1
		}

		o2Matches = filterLines(i, o2Val, o2Matches)

		if len(o2Matches) == 1 {
			r.O2Rating = BinaryStringToInt(o2Matches[0])
			break
		}
	}

	co2Matches := r.lines
	for i := 0; i < width; i++ {
		zeroesPerColumn := perColumn('0', co2Matches)
		onesPerColumn := perColumn('1', co2Matches)

		// Oxygen values are the most common, CO2 are the least.
		o2Val := 0
		if zeroesPerColumn[i] == onesPerColumn[i] || zeroesPerColumn[i] < onesPerColumn[i] {
			o2Val = 1
		}
		co2Val := 0
		if o2Val == 0 {
			co2Val = 1
		}

		co2Matches = filterLines(i, co2Val, co2Matches)

		if len(co2Matches) == 1 {
			r.CO2Rating = BinaryStringToInt(co2Matches[0])
			break
		}
	}
}

func (r Report) Consumption() int {
	return r.Gamma * r.Epsilon
}

func (r Report) LifeSupportRating() int {
	return r.O2Rating * r.CO2Rating
}

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
