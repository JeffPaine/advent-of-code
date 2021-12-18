package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var total int = 0
	seen := make(map[int]bool)

	for {

		file, err := os.Open("../input.txt")
		if err != nil {
			log.Fatal("Couldn't read file: ", err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal("Error converting string: ", err)
			}
			total += i
			if _, ok := seen[total]; ok {
				fmt.Println("Repeated value: ", total)
				return
			}
			seen[total] = true
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Error with scanning: ", err)
		}
		file.Close()

	}
}
