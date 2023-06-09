package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
)

func stringToInt(value string) int {
	output, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	} else if err == nil {
		return output
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

	temp := []int{}
	summedValues := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			temp = append(temp, stringToInt(scanner.Text()))
		} else if scanner.Text() == "" {
			summedValues = append(summedValues, sumArray(temp))
			temp = []int{}
		}
	} 

	sort.SliceStable(summedValues, func(i, j int) bool { return summedValues[i] > summedValues[j] })
	
	singleResult := summedValues[0]
	top3Result := sumArray(summedValues[:3])

	fmt.Println("The largest number of the array is: ", singleResult, "and the sum of the largest 3 is:", top3Result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}