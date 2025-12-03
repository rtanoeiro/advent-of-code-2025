package day02

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Resolve() string {

	inputData, errInput := input.ReadInputFile("day02")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindInvalidIDsPart1(inputData)
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindInvalidIDsPart2(inputData)
	log.Printf("\n")
	return fmt.Sprintf(
		"The sum of all invalid IDS for Part1 is %d. The result for Part2 is %d",
		result1,
		result2,
	)
}

// *** PART 01 *** //
// Since the young Elf was just doing silly patterns, you can find the invalid IDs
// by looking for any ID which is made only of some sequence of digits repeated twice.
// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.

// None of the numbers have leading zeroes; 0101 isn't an ID at all.
// (101 is a valid ID that you would ignore.)
func FindInvalidIDsPart1(inputData []string) int {
	// Each item in the list is a range
	ranges := strings.Split(inputData[0], ",")
	var InvalidIds []int
	log.Println(ranges)
	for index := range ranges {
		min := strings.Split(ranges[index], "-")[0]
		max := strings.Split(ranges[index], "-")[1]
		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)
		InvalidIds = append(InvalidIds, FindInvalidPart1(minInt, maxInt)...)
	}
	return SumInvalidPart1(InvalidIds)
}

func FindInvalidPart1(min, max int) []int {
	var InvalidIds []int
	for i := min; i <= max; i++ {
		iStr := strconv.Itoa(i)
		sizeSr := len(iStr)
		if sizeSr%2 != 0 {
			continue
		}

		// If each half of the ID string is the same, then it's invalid
		if iStr[:sizeSr/2] == iStr[sizeSr/2:] {
			InvalidIds = append(InvalidIds, i)
		}
	}
	return InvalidIds
}

func SumInvalidPart1(Ids []int) int {
	sum := 0
	for index := range Ids {
		sum += Ids[index]
	}
	return sum
}

// *** PART 01 *** //

// *** PART 02 *** //

// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice.
// So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times),
// and 1111111 (1 seven times) are all invalid IDs.
func FindInvalidIDsPart2(inputData []string) int {
	// Based on input data, we get ranges we'll loop through
	ranges := strings.Split(inputData[0], ",")
	var InvalidIds []bool
	log.Println(ranges)
	sum := 0
	for index := range ranges {
		log.Printf("Current range %s", ranges[index])
		min := strings.Split(ranges[index], "-")[0]
		max := strings.Split(ranges[index], "-")[1]
		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)
		// Within each range we hae a min and max, and we'll find invalid ids in that range
		InvalidIds = FindInvalidPart2(minInt, maxInt)
		sum += SumInvalidPart2(InvalidIds, []int{minInt, maxInt})
	}
	return sum
}

// For part2, I decided to return a index of true/false, since comparison are more dinamic,
// as in, we need to check if various subset of the text are in the text, whenever the subset is
// found throughout the whole text, I'll tag that as True meaning that number is an invalid ID.
// This is different from part 1, because for part 1 I just needed to compare each half of the string,
// but here we need to compare several pieces of the string.
func FindInvalidPart2(min, max int) []bool {
	var InvalidIds []bool
	// Looping through through the range to find ids, we need to convert to string,
	// as we're comparing pieces of a number, hence the string conversion
	for i := min; i <= max; i++ {
		iStr := strconv.Itoa(i)
		log.Printf("Starting find function on number %d", i)
		// We always start with the first digit of the sequence, since this is a recursive function
		// The subset piece will keep increasing up to half + 1 of the string size
		result := FindStrSequence(iStr[0:1], iStr)
		InvalidIds = append(InvalidIds, result)
		//	log.Printf("Finished finding invalid ID for %d. Found: %v", i, InvalidIds)
	}
	return InvalidIds
}

// Finds if a sequence of text appears in the text repeatedly
func FindStrSequence(subset, text string) bool {
	// By default we start with valid = false, if we don't find anything across the text ever, then
	// it means it's not an invalid id
	valid := false
	sizeSubset := len(subset)
	sizeText := len(text)

	// This means we went over half of the text, so it's impossible to find a repetion, so it's false
	// This is the base case on which recursion stops
	if sizeSubset > sizeText/2 {
		//	log.Printf("Reach half of the string, returning false")
		return false
	}

	// In this case, the subset would not fill repeatdly inside the whole string, then we just
	// enter the function again with a bigger subset
	if sizeText%sizeSubset != 0 {
		//	log.Printf("Incompatible size, checking next sequence %s", text[0:sizeSubset+1])
		return FindStrSequence(text[0:sizeSubset+1], text)
	}

	// When none of the conditions above are satisfied, we can *possibly* find the subset inside the text
	// Repetitions is how many times that subset can appear inside the text, this will determing how
	// many times we'll slice the whole text forward.
	repetitions := sizeText / sizeSubset
	//log.Printf("Number of repetitions considering subset %s on whole string %s: %d", subset, text, repetitions)

	for i := 1; i < repetitions; i++ {
		// If we have 121212, and we have subset = "12" we start at 1 (0 is already the current subset)
		// We go from 1 up to as many repetitions the subtext allows.
		// In this example, that would be:
		// 1. "12" == "12*12*12"
		// 2. "12" == "1212*12*"
		// Because in this case the current subset == selectedText (sliced string), then valid becomes true
		// both iteractions, if we had "12" and the string "121211"
		// In the last iteraction, we would fall into the else condition, valid = false, then the loop
		// is broken, cause we couldn't find a sequence, we go over the next slice "121"
		// Repetitions there is 2 (121 fits twice in string of size 6)
		// In the first repetition we would find false, as "121" != "121*211*"
		selectedText := text[i*sizeSubset : (i+1)*sizeSubset]
		//log.Printf("Comparing %s == %s", subset, selectedText)
		if subset == selectedText {
			valid = true
			continue
		} else {
			valid = false
		}
		if !valid {
			//	log.Printf("Did not find any subset, increasing subset to %s", text[0:sizeSubset+1])
			return FindStrSequence(text[0:sizeSubset+1], text)
		}
	}
	return valid
}

func SumInvalidPart2(Ids []bool, numbers []int) int {
	sum := 0
	for index := range Ids {
		if Ids[index] {
			log.Printf("Adding the following number to the result %d", numbers[0]+index)
			sum += numbers[0] + index
		}
	}
	return sum
}

// *** PART 02 *** //
