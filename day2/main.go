package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

func sumArray(values[] int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func assignScoreFirstRound(firstIn string, secondIn string) int {
	elfValue := 0
	playerValue := 0
	winOrLoss := 0
	finalScore := 0

	switch firstIn {
		case "A": elfValue = 1
		case "B": elfValue = 2
		case "C": elfValue = 3
	}
	switch secondIn {
		case "X": playerValue = 1
		case "Y": playerValue = 2
		case "Z": playerValue = 3
	}
	 
	if elfValue == playerValue {
		winOrLoss = 3
	} else if elfValue == 1 {
		if playerValue == 2 {
			winOrLoss = 6
		} else if playerValue == 3 {
			winOrLoss = 0
		} 
	} else if elfValue == 2 {
		if playerValue == 1 {
			winOrLoss = 0
		} else if playerValue == 3 {
			winOrLoss = 6
		}
	} else if elfValue == 3 {
		if playerValue == 1 {
			winOrLoss = 6
		} else if playerValue == 2 {
			winOrLoss = 0
		}
	}

	finalScore = playerValue + winOrLoss
	return finalScore
}

func assignScoreSecondRound(firstIn string, secondIn string) int {
	elfValue := 0
	playerValue := 0
	winOrLoss := 0

	switch firstIn {
		case "A": elfValue = 1
		case "B": elfValue = 2
		case "C": elfValue = 3
	}
	switch secondIn {
		case "X": playerValue = 1
		case "Y": playerValue = 2
		case "Z": playerValue = 3
	}
	 
	if playerValue == 2 { //if player choses to draw, player will get 3 points plus whatever elf choses
		winOrLoss = 3 + elfValue
	} else if playerValue == 1 { //if player choses to lose, value should be whatever would lose to elf plus zero
		if elfValue == 1 {
			winOrLoss = 3
		} else if elfValue == 2 {
			winOrLoss = 1
		} else if elfValue == 3 {
			winOrLoss = 2
		} 
	} else if playerValue == 3 { //if player choses to win, value should be whatever would win to elf plus six
		if elfValue == 1 {
			winOrLoss = 6 + 2
		} else if elfValue == 2 {
			winOrLoss = 6 + 3
		} else if elfValue == 3 {
			winOrLoss = 6 + 1
		}
	}
	return winOrLoss
}

func main() {
	totalScore := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstChar := string(scannedLine[0])
		lastChar := string(scannedLine[len(scannedLine)-1])

		// totalScore = append(totalScore, assignScoreFirstRound(firstChar, lastChar))
		// fmt.Println("The first letter is ", firstChar, "and the last letter is ", lastChar, "and the score is ", assignScoreFirstRound(firstChar, lastChar))
		
		totalScore = append(totalScore, assignScoreSecondRound(firstChar, lastChar))
	} 

	fmt.Println("The total score is: ", sumArray(totalScore))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}