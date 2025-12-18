package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	id      int
	s       string
	x, y, z int
}

func newBox(s string, i int) box {
	nums := strings.Split(string(s), ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	z, _ := strconv.Atoi(nums[2])
	return box{id: i, s: s, x: x, y: y, z: z}
}

// Calculate the Euclidean Distance between two box locations.
// https://en.wikipedia.org/wiki/Euclidean_distance#Higher_dimensions.
func distance(a, b box) float64 {
	xDiff := math.Pow(float64(a.x)-float64(b.x), 2)
	yDiff := math.Pow(float64(a.y)-float64(b.y), 2)
	zDiff := math.Pow(float64(a.z)-float64(b.z), 2)
	return math.Sqrt(xDiff + yDiff + zDiff)
}

type diff struct {
	a, b     box
	distance float64
}

// A Disjoint Set Union (DSU) AKA union-find AKA merge-find.
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure.
//
// This approach settled on after first working out a brute-force-ish solution
// that ended up being Kruskal's Algorithm
// (https://en.wikipedia.org/wiki/Kruskal%27s_algorithm) and then getting some
// AI feedback that pointed me to a DSU solution.
type dsu struct {
	// parents[i] tells us who the parents of entry i is. If parents[i] == i then i
	// is itself the root of the group that it's in.
	parents []int
	// sizes[i] is the sizes of the group i is in.
	sizes []int
}

func newDSU(n int) *dsu {
	d := &dsu{
		parents: make([]int, n),
		sizes:   make([]int, n),
	}
	for i := range n {
		// To start, each entry is its own parent / root, i.e. in a group of size 1.
		d.parents[i] = i
		// To start, each group has a size of 1.
		d.sizes[i] = 1
	}
	return d
}

// find finds the root for an entry (or the group it's in) by traveling up the
// chain of parents as far as required.
func (d *dsu) find(i int) int {
	if d.parents[i] != i {
		// We're not at the root (i.e. this entry's parent is some other entry).
		// Traverse up the chain until we find the root (parent[i] == i) and
		// then update all entries in the chain to point directly at the root,
		// which will save us from having to recalculate them in the future.
		d.parents[i] = d.find(d.parents[i])
	}
	// Now that d.parents[i] points at the root, return that.
	return d.parents[i]
}

// Union merges two sets together. Returns true if a merge happened, false if
// they were already merged
func (d *dsu) union(i, j int) bool {
	rootI := d.find(i)
	rootJ := d.find(j)

	// These two entries already have the same root, no need to merge them as
	// they're already merged.
	if rootI == rootJ {
		return false
	}

	// We need to merge the two groups which means we need to pick a "direction"
	// (e.g. which one into which one). Merge smaller groups into larger groups
	// just to pick a direction.
	if d.sizes[rootI] < d.sizes[rootJ] {
		// Swap roots so that rootI is always has the largest size, this way our
		// merge direction is always the same: merge j --> i.
		rootI, rootJ = rootJ, rootI
	}

	// Do the actual merging.
	d.parents[rootJ] = rootI
	d.sizes[rootI] += d.sizes[rootJ]

	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}

	var boxes []box

	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		b := newBox(scanner.Text(), line)
		boxes = append(boxes, b)
		line++
	}

	var diffs []diff

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			a := boxes[i]
			b := boxes[j]
			d := diff{a: a, b: b, distance: distance(a, b)}
			diffs = append(diffs, d)
		}
	}

	// Sort diffs shortest --> longest.
	slices.SortFunc(diffs, func(a, b diff) int {
		return cmp.Compare(a.distance, b.distance)
	})

	d := newDSU(len(boxes))
	total1 := 1
	total2 := 0
	rounds := 0
	curr := diff{}
	for _, df := range diffs {
		curr = df

		// Solution 1.
		if rounds == 1000 {
			// Multiply together the three largest sizes. BUT, only for entries
			// that are roots. Other entry counts could be for a child part of
			// another entry and we don't want to consider them.
			var sizes []int
			for i := range len(d.sizes) {
				if d.parents[i] == i {
					sizes = append(sizes, d.sizes[i])
				}
			}
			slices.Sort(sizes)
			bottomN := 3
			for i := len(sizes) - 1; i >= len(sizes)-bottomN; i-- {
				total1 *= sizes[i]
			}
		}

		d.union(curr.a.id, curr.b.id)

		// Solution 2.
		root := d.find(curr.a.id)
		if d.sizes[root] == len(boxes) {
			total2 = curr.a.x * curr.b.x
			break
		}

		rounds++
	}

	// Answer: 62186
	fmt.Println("Solution 1:", total1)

	// Answer: 8420405530
	fmt.Println("Solution 2:", total2)
}
