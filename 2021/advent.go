package advent

import (
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

	Gamma   int
	Epsilon int
}

func NewReport(lines []string) Report {
	r := Report{lines: lines}
	r.populateFields()
	return r
}

func (r *Report) populateFields() {
	// We assume all lines are the same width.
	width := len(r.lines[0])

	// Count the the number of 0s and 1s per column.
	zeroesPerColumn := make([]int, width)
	onesPerColumn := make([]int, width)
	for _, line := range r.lines {
		for i, val := range line {
			switch val {
			case '0':
				zeroesPerColumn[i] += 1
			case '1':
				onesPerColumn[i] += 1
			}
		}
	}

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
