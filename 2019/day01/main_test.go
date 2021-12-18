package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestOne(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	want := 3182375
	got := one(nums)
	if got != want {
		t.Errorf("%d != %d", got, want)
	}
}

func TestTwo(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	want := 4770725
	got := two(nums)
	if got != want {
		t.Errorf("%d != %d", got, want)
	}
}
