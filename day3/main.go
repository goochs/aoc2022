package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func splitRuck(line string) (x, y string) {
	lineLength := len(line)
	firstCompartment := line[:lineLength/2]
	secondCompartment := line[lineLength/2:]
	return firstCompartment, secondCompartment
	}

func sharedLetters(a, b string) rune {
	chars := []rune(a)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if strings.Contains(b, char) {
			return chars[i]
			}
		}
	return 0
	}

func RuneToInteger(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	} 
	return 0
}

func sumArray(values[] int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalPriority := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstCompartment, secondCompartment := splitRuck(scannedLine)
		letter := sharedLetters(firstCompartment, secondCompartment)
		priority := RuneToInteger(letter)
		totalPriority = append(totalPriority, priority)

		// fmt.Println("The full line is: ", scannedLine)
		// fmt.Println("The first compartment is: ", firstCompartment)
		// fmt.Println("The second compartment is: ", secondCompartment)
		// fmt.Println("The shared letter is: ", string(letter))
		// fmt.Println("The priority is: ", priority)
		// break
		}
	fmt.Println("The total priority is: ", sumArray(totalPriority))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		}
	}