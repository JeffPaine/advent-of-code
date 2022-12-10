package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type file struct {
	size int
	name string
}

func (f file) String() string {
	return fmt.Sprintf("%v", f.name)
}

type dir struct {
	name      string
	parent    *dir
	dirs      []*dir
	files     []*file
	totalSize int
	sized     bool
}

func (d dir) String() string {
	dirs := []string{}
	for _, di := range d.dirs {
		dirs = append(dirs, di.name)
	}
	files := []string{}
	for _, f := range d.files {
		files = append(files, f.String())
	}
	return fmt.Sprintf("dir{name: %q, size: %d, dirs (%d): %v, files (%d): %v}", d.name, d.size(), len(d.dirs), dirs, len(d.files), files)
}

func (d dir) size() int {
	if d.sized {
		return d.totalSize
	}

	total := 0
	for _, f := range d.files {
		total += f.size
	}
	for _, d := range d.dirs {
		total += d.size()
	}
	d.totalSize = total
	return total
}

func dirsLessThan(d *dir, limit int) []*dir {
	dirs := []*dir{}

	if d.size() <= limit {
		dirs = append(dirs, d)
	}

	for _, cd := range d.dirs {
		dirs = append(dirs, dirsLessThan(cd, limit)...)
	}
	return dirs
}

func dirsMoreThan(d *dir, min int) []*dir {
	dirs := []*dir{}

	if d.size() >= min {
		dirs = append(dirs, d)
	}

	for _, cd := range d.dirs {
		dirs = append(dirs, dirsMoreThan(cd, min)...)
	}
	return dirs
}

func printDirs(d *dir) {
	if len(d.dirs) == 0 {
		fmt.Println(d)
	}

	for _, cd := range d.dirs {
		printDirs(cd)
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	root := &dir{name: "/"}
	var curr *dir

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, "dir ") {
			continue
		}
		if strings.HasPrefix(t, "$ ls") {
			continue
		}
		if strings.HasPrefix(t, "$ cd ..") {
			curr = curr.parent
			continue
		}
		if strings.HasPrefix(t, "$ cd /") {
			curr = root
			continue
		}
		if strings.HasPrefix(t, "$ cd") {
			name := strings.Split(t, " ")[2]
			d := &dir{name: name, parent: curr}
			curr.dirs = append(curr.dirs, d)
			curr = d
			continue
		}
		// All that's left should be file listings.
		f := &file{}
		if _, err := fmt.Sscanf(t, "%d %s", &f.size, &f.name); err != nil {
			log.Fatal(err)
		}
		curr.files = append(curr.files, f)
	}

	dirs := dirsLessThan(root, 100000)
	total := 0
	for _, d := range dirs {
		total += d.size()
	}
	fmt.Println("Solution 1:", total)

	totalDisk := 70000000
	desiredFree := 30000000
	currFree := totalDisk - root.size()
	min := desiredFree - currFree
	dirs = dirsMoreThan(root, min)
	smallest := dirs[0].size()
	for _, d := range dirs {
		if d.size() < smallest {
			smallest = d.size()
		}
	}
	fmt.Println("Solution 2:", smallest)
}
