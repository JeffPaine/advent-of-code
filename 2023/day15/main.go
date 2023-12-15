package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func hash(s string) int {
	curr := 0

	for _, char := range s {
		curr += int(char)
		curr *= 17
		curr = curr % 256
	}

	return curr
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var steps []string
	for scanner.Scan() {
		steps = append(steps, strings.Split(scanner.Text(), ",")...)
	}

	t1 := 0
	for _, step := range steps {
		t1 += hash(step)
	}
	// Answer: 514281.
	fmt.Println("Solution 1:", t1)

	// Part 2.
	m := make(map[int][]lens)
	for i := 0; i < 256; i++ {
		m[i] = []lens{}
	}

	for _, step := range steps {
		if strings.Contains(step, "-") {
			lensLabel := step[:len(step)-1]
			boxIdx := hash(lensLabel)
			box := m[boxIdx]
			lensIdx := -1
			for i := 0; i < len(box); i++ {
				if box[i].label == lensLabel {
					lensIdx = i
					break
				}
			}
			if lensIdx != -1 {
				if lensIdx == len(box)-1 {
					m[boxIdx] = box[:lensIdx]
				} else {
					m[boxIdx] = append(box[:lensIdx], box[lensIdx+1:]...)
				}
			}
		}
		if strings.Contains(step, "=") {
			data := strings.Split(step, "=")
			lensLabel := data[0]
			boxIdx := hash(lensLabel)
			focalLength, err := strconv.Atoi(data[1])
			if err != nil {
				panic(fmt.Sprintf("can't convert to int: %v", data[1]))
			}
			box := m[boxIdx]
			exists := false
			for i := 0; i < len(box); i++ {
				if box[i].label == lensLabel {
					exists = true
					box[i].focalLength = focalLength
					break
				}
			}
			if !exists {
				box = append(box, lens{label: lensLabel, focalLength: focalLength})
				m[boxIdx] = box
			}
		}
	}

	t2 := 0
	for boxIdx, lenses := range m {
		for lensIdx, l := range lenses {
			t2 += (1 + boxIdx) * (lensIdx + 1) * l.focalLength
		}
	}
	// Answer: 244199.
	fmt.Println("Solution 2:", t2)
}
