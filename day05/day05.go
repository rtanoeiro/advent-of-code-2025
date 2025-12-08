package day05

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Resolve() string {

	startTime := time.Now()
	inputData, errInput := input.ReadInputFile("day05")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindFresh(inputData, "1")
	log.Printf("The number of Paper Rolls for part 1 is %d", result1)
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindFresh(inputData, "2")
	log.Printf("The result for Part2 is %d", result2)
	log.Printf("\n")
	elapsedTime := time.Since(startTime)
	log.Printf("Results recieved %.6f seconds", elapsedTime.Seconds())
	return fmt.Sprint("All results have been found!")
}

func FindFresh(inputData []string, part string) int64 {
	var ranges []string
	var ids []string
	var lineSplit int
	for index := range inputData {
		if len(inputData[index]) == 0 {
			lineSplit = index
		}
	}

	for index := range inputData {
		if index < lineSplit {
			ranges = append(ranges, inputData[index])
		}
		if index > lineSplit {
			ids = append(ids, inputData[index])
		}
	}

	switch part {
	case "1":
		fresh := int64(0)
		for _, itemId := range ids {
			itemInt, _ := strconv.Atoi(itemId)
			for _, currentRange := range ranges {
				isFresh := false
				rangeSplit := strings.Split(currentRange, "-")
				rangeMin, _ := strconv.ParseInt(rangeSplit[0], 10, 64)
				rangeMax, _ := strconv.ParseInt(rangeSplit[1], 10, 64)
				isFresh = checkRanges(rangeMin, rangeMax, int64(itemInt))

				// In case we find the item ID in ANY of the ranges, we already skip to the next itemID
				if isFresh {
					fresh++
					break
				}
			}
		}
		return fresh

	case "2":
		sortedRanges := sortRanges(ranges)
		removedOverlapping := removeOverlapping(sortedRanges)

		sum := int64(0)
		for index := range removedOverlapping {
			currentRange := strings.Split(removedOverlapping[index], "-")
			curRangeMin, _ := strconv.ParseInt(currentRange[0], 10, 64)
			curRangeMax, _ := strconv.ParseInt(currentRange[1], 10, 64)
			//log.Printf("Checking how many Ids there are in the range %d-%d", curRangeMin, curRangeMax)
			sum += curRangeMax - curRangeMin + 1
		}
		return sum
	}
	return 0
}

func checkRanges(rangeMin, rangeMax, itemID int64) bool {
	if itemID >= rangeMin && itemID <= rangeMax {
		return true
	}
	return false
}

// In order to cleverly check all ids in a range we need sort each range by their min value.
// It doesn't matter each range maximum value, only their minimum, cause we can take a look if the minimum from the next range
// is bigger than the current range maximum, if they are, the next range starts with the current minimum + 1
// I decided to use bubble sort cause it's easier to use it when we have a range, I couldn't visualize using merge sort
// recursion problem and keep track of the maximum at the same time so I could actually sorte the full range, not just the min values
func sortRanges(ranges []string) []string {
	end := len(ranges)
	for {
		var swapping = false
		for index := 1; index < end; index++ {
			currentRange := strings.Split(ranges[index], "-")
			curRangeMin, _ := strconv.ParseInt(currentRange[0], 10, 64)
			curRangeMax, _ := strconv.ParseInt(currentRange[1], 10, 64)

			previousRange := strings.Split(ranges[index-1], "-")
			prevRangeMin, _ := strconv.ParseInt(previousRange[0], 10, 64)
			prevRangeMax, _ := strconv.ParseInt(previousRange[1], 10, 64)

			if prevRangeMin > curRangeMin {
				temp := fmt.Sprintf("%d-%d", prevRangeMin, prevRangeMax)
				ranges[index-1] = fmt.Sprintf("%d-%d", curRangeMin, curRangeMax)
				ranges[index] = temp
				swapping = true

			}
		}
		end--

		if !swapping {
			break
		}
	}
	return ranges
}

func removeOverlapping(ranges []string) []string {
	for index := 0; index < len(ranges); index++ {
		if index == len(ranges)-1 {
			break
		}

		currentRange := strings.Split(ranges[index], "-")
		curRangeMin, _ := strconv.ParseInt(currentRange[0], 10, 64)
		curRangeMax, _ := strconv.ParseInt(currentRange[1], 10, 64)

		nextRange := strings.Split(ranges[index+1], "-")
		nextRangeMin, _ := strconv.ParseInt(nextRange[0], 10, 64)
		nextRangeMax, _ := strconv.ParseInt(nextRange[1], 10, 64)

		if nextRangeMin <= curRangeMax {
			if nextRangeMax <= curRangeMax {
				log.Printf("Next range %d-%d, is fully contained in current range %d-%d", curRangeMin, curRangeMax, nextRangeMin, nextRangeMax)
				updatedNextRange := fmt.Sprintf("%d-%d", curRangeMax+1, curRangeMax)
				ranges[index+1] = updatedNextRange
				continue
			}
			updatedNextRange := fmt.Sprintf("%d-%d", curRangeMax+1, nextRangeMax)
			ranges[index+1] = updatedNextRange
			continue
		}
	}
	return ranges
}
