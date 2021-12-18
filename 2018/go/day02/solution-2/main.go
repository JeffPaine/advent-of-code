package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func justOneCharDiff(s1 string, s2 string) bool {
	diffs := 0
	s1Runes := []rune(s1)
	s2Runes := []rune(s2)

	for i := 0; i < len(s1Runes); i++ {
		if s1Runes[i] != s2Runes[i] {
			diffs += 1
		}
	}
	return diffs == 1
}

func main() {

	entries := make([]string, 0)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanner: ", err)
	}

	for i, _ := range entries {
		s1 := entries[i]
		for _, s2 := range entries[i+1:] {
			if justOneCharDiff(s1, s2) {
				fmt.Println("Found two strings that only differ by one character:")
				fmt.Println(s1)
				fmt.Println(s2)

				fmt.Println("All matching characters:")
				s1Runes := []rune(s1)
				s2Runes := []rune(s2)
				for j := 0; j < len(s1Runes); j++ {
					if s1Runes[j] == s2Runes[j] {
						fmt.Printf("%c", s1Runes[j])
					}
				}
				fmt.Println()
			}
		}
	}

}
