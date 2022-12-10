package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	// Include filename and line in logs.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type crate struct {
	val rune
}

func (c crate) String() string {
	return fmt.Sprintf("%c", c.val)
}

type stack struct {
	crates []crate
}

func (s stack) String() string {
	return fmt.Sprintf("stack{crates: %v}", s.crates)
}

func (s *stack) push(c crate) {
	// "Top" is the last item.
	s.crates = append(s.crates, c)

	// "Top" is the first item.
	// s.crates = append([]crate{c}, s.crates...)
}

func (s *stack) pop() crate {
	if len(s.crates) == 0 {
		log.Fatalf("attempted to call pop() on empty stack: %v", s)
	}
	// "Top" is the last item.
	c := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]

	// "Top" is the first item.
	// c := s.crates[0]
	// s.crates = s.crates[1:]
	return c
}

func (s *stack) peek() crate {
	// "Top" is the last item.
	return s.crates[len(s.crates)-1]

	// "Top" is the first item.
	// return s.crates[0]
}

type step struct {
	count, from, to int
}

func apply(stacks []stack, s step) {
	for i := 0; i < s.count; i++ {
		stacks[s.to].push(stacks[s.from].pop())
	}
}

// apply9001 is like apply, but it applies multiple moves at once and maintains their original ordering
// on the destination stack.
func apply9001(stacks []stack, s step) {
	temp := stack{}
	for i := 0; i < s.count; i++ {
		temp.push(stacks[s.from].pop())
	}
	for i := 0; i < s.count; i++ {
		stacks[s.to].push(temp.pop())
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := 0
	stackLines := 8
	stackCount := 9
	input := make([][]rune, stackCount)

	steps := []step{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		// Parse the stacks assignment section, which is 8 lines long.
		if lines < stackLines {
			currStack := 0
			line := []rune(scanner.Text())
			// The number we care about is every 4 characters.
			for i := 1; i < len(line); i += 4 {
				token := line[i]
				if token != ' ' {
					input[currStack] = append(input[currStack], token)
				}
				currStack++
			}
			lines++
			continue
		}

		// Skip the stack number and following blank line.
		if lines < stackLines+2 {
			lines++
			continue
		}

		// Parse the steps section.
		s := step{}
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &s.count, &s.from, &s.to)
		if err != nil {
			log.Fatalf("fmt.Sscanf(%q, \"move %%d from %%d to %%d\", ...) error: %v", scanner.Text(), err)
		}
		// Steps are given as 1-indexed, but our internal representation is 0-indexed, so convert them.
		s.from -= 1
		s.to -= 1
		steps = append(steps, s)
	}

	// Stacks are parsed in top first, but we need to push() them on to the stack in bottom first order.
	stacks1 := make([]stack, stackCount)
	stacks2 := make([]stack, stackCount)
	for i, in := range input {
		// Reverse iterate.
		for j := len(in) - 1; j >= 0; j-- {
			// Populate two at once as these are slices of slices and deep copying them would be a pain.
			stacks1[i].push(crate{in[j]})
			stacks2[i].push(crate{in[j]})
		}
	}

	// Solution 1.
	for _, st := range steps {
		apply(stacks1, st)
	}
	sol1 := []string{}
	for _, s := range stacks1 {
		sol1 = append(sol1, s.peek().String())
	}
	fmt.Println("Solution 1:", strings.Join(sol1, ""))

	// Solution 2.
	for _, st := range steps {
		apply9001(stacks2, st)
	}
	sol2 := []string{}
	for _, s := range stacks2 {
		sol2 = append(sol2, s.peek().String())
	}
	fmt.Println("Solution 2:", strings.Join(sol2, ""))
}
