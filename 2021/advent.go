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

type Power struct {
	Gamma   int64
	Epsilon int64
}

func InitPower(lines []string) Power {
	// Assumes all lines are the same width.
	width := len(lines[0])

	// Calculate the most and least frequent value per column.
	zeroesPerColumn := make([]int, width)
	onesPerColumn := make([]int, width)
	for _, line := range lines {
		for i, val := range line {
			switch val {
			case '0':
				zeroesPerColumn[i] += 1
			case '1':
				onesPerColumn[i] += 1
			}
		}
	}

	// Convert the per-column frequencies to decimal values.
	// E.g. [1, 0, 1, 1] -> 11.
	gamma := ""
	epsilon := ""
	for i := 0; i < width; i++ {
		if zeroesPerColumn[i] > onesPerColumn[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return Power{Gamma: g, Epsilon: e}
}

func (p Power) Consumption() int64 {
	return p.Gamma * p.Epsilon
}
