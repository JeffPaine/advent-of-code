package main

import (
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
}
