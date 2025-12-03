package day03

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

var numPart1 = 2
var numPart2 = 12

func Resolve() string {

	startTime := time.Now()
	inputData, errInput := input.ReadInputFile("day03")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindMaxJoltage(inputData, numPart1)
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindMaxJoltage(inputData, numPart2)
	log.Printf("\n")
	elapsedTime := time.Since(startTime)
	log.Printf("Results recieved %.6f seconds", elapsedTime.Seconds())
	return fmt.Sprintf(
		"The maximum joltaje for for the current bank is %d .The result for Part2 is %d",
		result1,
		result2,
	)
}

func FindMaxJoltage(inputData []string, numDigits int) int {
	var maxJoltages [][]int
	for _, bank := range inputData {
		maxfound := findJoltage(bank, numDigits)
		maxJoltages = append(maxJoltages, maxfound)
	}
	sum := 0
	for _, joltage := range maxJoltages {
		sum += calculateMaxJoltageFromDigits(joltage, numDigits)
	}

	return sum
}

func calculateMaxJoltageFromDigits(digits []int, numDigits int) int {
	sum := 0
	for i := 0; i < numDigits; i++ {
		toAdd := digits[i] * int(math.Pow(10, float64(numDigits-i-1)))
		sum += toAdd
	}
	return sum
}

// For numDigits = 2 and text = 123456789, len = 9
// 1st number we go up to 8, index 7  = len(text) - numDigits + 0
// 2nd number we go up to 9 index 8  = len(text) - numDigits + 1

// For numDigits = 3 and text = 123456789, len = 9
// 1st number we go up to 7, index 6 = len(text) - numDigits + 0
// 2nd number we go up to 8 index 7 = len(text) - numDigits + 1
// 3rd number we go up to 9, index 8 = len(text) - numDigits + 2
func findJoltage(bank string, numDigits int) []int {
	// The goal is to find the max number in the list. Once that number is found
	// we go to all other elements to it's right, always finding the max,
	// until we find the biggest number with the number of digits we want
	var digits []int
	nextIndex := 0
	for i := 0; i < numDigits; i++ {
		subText := bank[nextIndex : len(bank)-numDigits+i+1] //+1 because the right side is not inclusive
		// In order to save some processing, once we find a digit, the biggest number MUST be after it, so we start from the next index after it
		digit, advanceIndex := findMax(subText)
		// Every time we find a number, the next higher index advances
		nextIndex += advanceIndex
		digits = append(digits, digit)
	}

	return digits
}

func findMax(text string) (int, int) {
	max, _ := strconv.Atoi(text[0:1]) // We assign the max value to the first number available
	nextIndex := 1                    // Because we assign the maximum to the first number, we assign the next Index to the very next index

	// In case we already got to the last number, we return it, as it should be the max
	if len(text) == 1 {
		return max, nextIndex
	}

	// We loop through each item in the sequence of numbers, if they are bigger than the max (our first item)
	// Then we save the max and their index, so we can start the next search from that index + 1
	for i := 1; i < len(text); i++ {
		// We know the index 0 is assigned to the max already, so we started from 1
		joltageInt, _ := strconv.Atoi(text[i : i+1])
		if joltageInt > max {
			max = joltageInt
			nextIndex = i + 1
		}
	}
	return max, nextIndex
}
