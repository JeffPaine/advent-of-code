package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	defer file.Close()

	var total int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Error converting string: ", err)
		}
		total += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanning: ", err)
	}

	fmt.Println(total)
}
