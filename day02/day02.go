package day02

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type ItemID struct {
	IsInvalid   bool
	Repetitions int
}

func Resolve() string {

	startTime := time.Now()
	inputData, errInput := input.ReadInputFile("day02")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindInvalidIDs(inputData, "1")
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindInvalidIDs(inputData, "2")
	log.Printf("\n")
	elapsedTime := time.Since(startTime)
	log.Printf("Results recieved %.4f seconds", elapsedTime.Seconds())
	return fmt.Sprintf(
		"The sum of all invalid IDS for Part1 is %d. The result for Part2 is %d",
		result1,
		result2,
	)
}

// Now, an ID is invalid if it is made only of some sequence of digits repeated JUST  twice for part 1
// Or AT LEAST twice for part 2
// So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times)
// would all be allowed on Part 2, but only 12341234 in Part 1
func FindInvalidIDs(inputData []string, part string) int {
	// Based on input data, we get ItemID we'll loop through
	AllItemIDs := strings.Split(inputData[0], ",")
	log.Println(AllItemIDs)
	sum := 0
	for index := range AllItemIDs {
		min := strings.Split(AllItemIDs[index], "-")[0]
		max := strings.Split(AllItemIDs[index], "-")[1]
		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)
		// Within each range we hae a min and max, and we'll find invalid ids in that range
		// Each ItemID in the list has the attribute Valid or not, and their number of repetitions
		ItemIDs := FindInvalid(minInt, maxInt, part)
		sum += SumInvalid(ItemIDs, []int{minInt, maxInt}, part)
	}
	return sum
}

// In each in each Range it's possible to have multiple items that are invalid
// So this function return a list where each index is tied to a number in the range min -> max
// So we can easily find an ItemID number by calculating min + current index in the ItemID list
func FindInvalid(min, max int, part string) []ItemID {
	var InvalidIds []ItemID
	// Looping through through the range to find ids, we need to convert to string,
	// as we're comparing pieces of a number, hence the string conversion
	for i := min; i <= max; i++ {
		var result ItemID
		iStr := strconv.Itoa(i)
		// log.Printf("Starting find function on number %d", i)
		// We always start with the first digit of the sequence, since this is a recursive function
		// The subset piece will keep increasing up to half + 1 of the string size
		switch part {
		case "1":
			// When we have single digits, then there's no half of it, so we use 1
			halfDigit := int(math.Max(float64(len(iStr)/2), 1))
			result = FindStrSequence(iStr[0:halfDigit], iStr)
		case "2":
			result = FindStrSequence(iStr[0:1], iStr)
		}
		InvalidIds = append(InvalidIds, result)
	}
	return InvalidIds
}

// Finds if a sequence of text appears in the text repeatedly
func FindStrSequence(subset, text string) ItemID {
	// By default we start with valid = false, if we don't find anything across the text ever, then
	// it means it's not an invalid id
	valid := false
	sizeSubset := len(subset)
	sizeText := len(text)

	// This means we went over half of the text, so it's impossible to find a repetion, so it's false
	// This is the base case on which recursion stops
	if sizeSubset > sizeText/2 {
		//	log.Printf("Reach half of the string, returning false")
		return ItemID{
			IsInvalid:   false,
			Repetitions: 0,
		}
	}

	// In this case, the subset would not fill repeatdly inside the whole string, then we just
	// enter the function again with a bigger subset
	if sizeText%sizeSubset != 0 {
		return FindStrSequence(text[0:sizeSubset+1], text)
	}

	// When none of the conditions above are satisfied, we can *possibly* find the subset inside the text
	// Repetitions is how many times that subset can appear inside the text, this will determing how
	// many times we'll slice the whole text forward.
	repetitions := sizeText / sizeSubset

	for i := 1; i < repetitions; i++ {
		// If we have 121212, and we have subset = "12" we start at 1 (0 is already the current subset)
		// We go from 1 up to as many repetitions the subtext allows.
		// In this example, that would be:
		// 1. "12" == "12*12*12"
		// 2. "12" == "1212*12*"
		// Because in this case the current subset == selectedText (sliced string), then valid becomes true
		// in both iteractions.
		// If we had "12" and the string "121211"
		// In the last iteraction, we would fall into the else condition, valid = false, then the loop
		// is broken, cause we couldn't find a sequence, we go over the next slice "121"
		// Repetitions there is 2 (121 fits twice in string of size 6)
		// In the first repetition we would find false, as "121" != "121*211*"
		selectedText := text[i*sizeSubset : (i+1)*sizeSubset]
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
	return ItemID{
		IsInvalid:   valid,
		Repetitions: repetitions,
	}
}

func SumInvalid(IDs []ItemID, IDRange []int, part string) int {
	sum := 0
	for index := range IDs {
		switch part {
		case "1":
			if IDs[index].IsInvalid && IDs[index].Repetitions == 2 {
				sum += IDRange[0] + index
			}

		case "2":
			if IDs[index].IsInvalid && IDs[index].Repetitions >= 2 {
				sum += IDRange[0] + index
			}
		}
	}
	return sum
}
