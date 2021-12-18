package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Example entry:
//
// #123 @ 3,2: 5x4
//
// "A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle
// 3 inches from the left edge, 2 inches from the top edge, 5 inches wide, and
// 4 inches tall."
type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

type row struct {
	spots []int
}

type fabric struct {
	rows []row
}

func makeClaim(s string) claim {
	c := claim{}
	split := strings.Split(s, " ")

	// split[0] = "#123"
	id, err := strconv.Atoi(strings.TrimPrefix(split[0], "#"))
	if err != nil {
		log.Fatal("Error converting string to int: ", err)
	}
	c.id = id

	// split[2] = "3,2:"
	split[2] = strings.TrimSuffix(split[2], ":")
	dimensions := strings.Split(split[2], ",")
	left, err := strconv.Atoi(dimensions[0])
	if err != nil {
		log.Fatal("Error converting string to int: ", err)
	}
	c.left = left

	top, err := strconv.Atoi(dimensions[1])
	if err != nil {
		log.Fatal("Error converting string to int: ", err)
	}
	c.top = top

	// split[3] = "5x4"
	nums := strings.Split(split[3], "x")
	width, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal("Error converting string to int: ", err)
	}
	c.width = width

	height, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal("Error converting string to int: ", err)
	}
	c.height = height

	return c
}

func main() {
	entries := make([]string, 0)
	claims := make([]claim, 0)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanning: ", err)
	}

	for _, entry := range entries {
		claims = append(claims, makeClaim(entry))
	}

	// Determine max dimmensions.
	maxWidth := 0
	maxHeight := 0
	for _, c := range claims {
		width := c.left + c.width
		height := c.top + c.height
		if width > maxWidth {
			maxWidth = width
		}
		if height > maxHeight {
			maxHeight = height
		}
	}

	// Create a fabric object with the number of rows and spots we need
	// based on max width and height.
	fab := fabric{rows: make([]row, maxHeight)}
	for i, _ := range fab.rows {
		fab.rows[i].spots = make([]int, maxWidth)
	}

	// Populate spots with info from claims.
	for _, c := range claims {
		for rowIdx, _ := range fab.rows {
			// Skip c.top number of rows.
			if rowIdx < c.top {
				continue
			}
			// Skip the bottom rows, too, if we're there.
			if rowIdx >= c.top+c.height {
				continue
			}
			// Now we're at a row we want to modify.
			for spotIdx, _ := range fab.rows[rowIdx].spots {
				// Skip c.left number of spots.
				if spotIdx < c.left {
					continue
				}
				// Skip the end of the row, if we're there.
				if spotIdx >= c.left+c.width {
					continue
				}
				// Now we're at the start of the spots we want to increment.
				// For the next c.width spots, increment them by one.
				fab.rows[rowIdx].spots[spotIdx] += 1
			}
		}

	}

	twoOrMore := 0
	// Count the number of spots that have two or more claims.
	for _, r := range fab.rows {
		for _, spot := range r.spots {
			if spot >= 2 {
				twoOrMore += 1
			}
		}

	}
	fmt.Println("Two or more claims: ", twoOrMore)
}
