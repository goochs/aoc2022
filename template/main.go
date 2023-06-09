package main

import (
	"bufio"
	"os"
	"log"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	} 

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}