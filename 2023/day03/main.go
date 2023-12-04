package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type part struct {
	start int
	end   int
	val   int
}

func isPeriod(r rune) bool {
	return r == '.'
}

func isDigit(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func isSymbol(r rune) bool {
	return !isPeriod(r) && !isDigit(r)
}

func sumPartNumbers(s string, rowLen int) (int, []part) {
	var parts []part
	total := 0
	curr := ""
	start := -1
	end := -1

	for idx, char := range s {
		if len(curr) == 0 && (isPeriod(char) || isSymbol(char)) {
			continue
		}
		if isDigit(char) {
			if start < 0 {
				start = idx
			}
			curr += string(char)
			// Keep going, unless we're at the end of the string.
			if idx != len(s)-1 {
				continue
			}
		}
		// We've just passed the end of the current number.
		end = idx - 1
		var toCheck []int
		// Check above.
		for i := start - rowLen - 1; i < end-rowLen+2; i++ {
			toCheck = append(toCheck, i)
		}
		// Check beside.
		toCheck = append(toCheck, start-1)
		toCheck = append(toCheck, end+1)
		// Check below
		for i := start + rowLen - 1; i < end+rowLen+2; i++ {
			toCheck = append(toCheck, i)
		}
		isPart := false
		for _, i := range toCheck {
			if i < 0 || i >= len(s) {
				continue
			}
			if isSymbol(rune(s[i])) {
				isPart = true
				break
			}
		}
		if isPart {
			val, _ := strconv.Atoi(curr)
			total += val
			p := part{start: start, end: end, val: val}
			parts = append(parts, p)
		}
		curr = ""
		start = -1
		end = -1
	}

	return total, parts
}

func sumGearRatios(s string, rowLen int, parts []part) int {
	total := 0
	for idx, char := range s {
		if char != '*' {
			continue
		}
		var nearby []part
		for _, part := range parts {
			// Above.
			aboveLeft := idx - rowLen - 1
			aboveRight := idx - rowLen + 1
			if part.end == aboveLeft || part.start == aboveRight || (idx-rowLen >= part.start && idx-rowLen <= part.end) {
				nearby = append(nearby, part)
			}
			// Beside.
			left := idx - 1
			right := idx + 1
			if part.end == left || part.start == right || (idx >= part.start && idx <= part.end) {
				nearby = append(nearby, part)
			}
			// Below.
			belowLeft := idx + rowLen - 1
			belowRight := idx + rowLen + 1
			if part.end == belowLeft || part.start == belowRight || (idx+rowLen >= part.start && idx+rowLen <= part.end) {
				nearby = append(nearby, part)
			}
		}
		if len(nearby) == 2 {
			total += nearby[0].val * nearby[1].val
		}
	}
	return total
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var s string
	var rowLen int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if rowLen == 0 {
			rowLen = len(scanner.Text())
		}
		s += scanner.Text()
	}

	total, parts := sumPartNumbers(s, rowLen)
	fmt.Println("Solution 1:", total)
	fmt.Println("Solution 2:", sumGearRatios(s, rowLen, parts))
}
