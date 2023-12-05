package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type chart struct {
	dest   int
	source int
	length int
}

func newChart(nums []string) chart {
	// nums: [50 98 2]
	// Pos 0: desination range start
	// Pos 1: source range start
	// Pos 2: range length
	var c chart
	// TODO: handle errors.
	val, _ := strconv.Atoi(nums[0])
	c.dest = val
	val, _ = strconv.Atoi(nums[1])
	c.source = val
	val, _ = strconv.Atoi(nums[2])
	c.length = val
	return c
}

type bundle struct {
	charts []chart
}

func (b bundle) lookup(val int) int {
	for _, c := range b.charts {
		if val >= c.source && val <= c.source+c.length {
			// val is in this chart.
			offset := c.source - c.dest
			return val - offset
		}
	}
	// val is not in our charts, so it must be 1:1.
	return val
}

func parse(r io.Reader) ([]int, []bundle) {
	scanner := bufio.NewScanner(r)
	first := true
	var seeds []int
	// The input has 7 maps in it.
	bundles := make([]bundle, 7)
	bundleNum := -1

	for scanner.Scan() {
		if first {
			nums := strings.Fields(strings.Split(scanner.Text(), ":")[1])
			for _, num := range nums {
				// TODO: handle errors.
				val, _ := strconv.Atoi(num)
				seeds = append(seeds, val)
			}
			first = false
			continue
		}
		if scanner.Text() == "" {
			continue
		}
		if strings.HasSuffix(scanner.Text(), ":") {
			bundleNum++
			continue
		}
		// Now processing a line. Example: 50 98 2
		c := newChart(strings.Fields(scanner.Text()))
		bundles[bundleNum].charts = append(bundles[bundleNum].charts, c)
	}
	return seeds, bundles
}

func solutionOne(seeds []int, bundles []bundle) int {
	var locations []int
	for _, seed := range seeds {
		out := seed
		for _, b := range bundles {
			out = b.lookup(out)
		}
		// fmt.Printf("seed: %d, out: %d\n", seed, out)
		locations = append(locations, out)
	}

	return slices.Min(locations)
}

func solutionTwo(seeds []int, bundles []bundle) int {
	// Seeds are now a range.
	minLocation := math.MaxInt
	base := 0
	for idx, seed := range seeds {
		if idx%2 == 0 {
			base = seed
			continue
		}
		for i := 0; i <= seed+1; i++ {
			location := base + i
			for _, b := range bundles {
				location = b.lookup(location)
			}
			if location < minLocation {
				minLocation = location
			}
		}
	}

	return minLocation
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	seeds, bundles := parse(f)
	fmt.Println("Solution 1", solutionOne(seeds, bundles))
	fmt.Println("Solution 2", solutionTwo(seeds, bundles))
}
