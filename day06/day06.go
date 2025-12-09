package day06

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"strings"
	"time"
)

func Resolve() string {

	startTime := time.Now()
	inputData, errInput := input.ReadInputFile("day06")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindSum(inputData, "1")
	log.Printf("Result for part 1 is %d", result1)
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindSum(inputData, "2")
	log.Printf("The result for Part2 is %d", result2)
	log.Printf("\n")
	elapsedTime := time.Since(startTime)
	log.Printf("Results recieved %.6f seconds", elapsedTime.Seconds())
	return fmt.Sprint("All results have been found!")
}

func FindSum(inputData []string, part string) int {
	// numLists contains how many items we have in a horizontal line, each item represents the first item of the vertical list
	// We'll save each vertical list inside the numbers list
	numLists := strings.Fields(inputData[0])
	numbers := make([][]string, len(numLists))

	if part == "1" {
		for hIndex := range len(inputData) {
			lineNumbers := strings.Fields(inputData[hIndex])
			log.Printf("Items in line after split %v", lineNumbers)
			// For example, when hIndex = 0, we are in the first line of file.
			// Each vIndex is a item in this first line, therefore, each vIndex is the first
			// element for each vertical calculation, so we save that into lists
			// By the end we turned a horizontal list into vertical lists
			for vIndex := range numLists {
				numbers[vIndex] = append(numbers[vIndex], lineNumbers[vIndex])
			}

		}
	}

	if part == "2" {
		// inputData contains each row of numbers, each index is a row
	}

	log.Printf("Lists created from input: %v", numbers)
	log.Printf("First List created: %v", numbers[0])

	sum := 0
	// Now that each new list represents the vertical items of the input, we can go over each one easily
	for i := 0; i < len(numbers); i++ {
		operator := numbers[i][len(numbers[i])-1]
		log.Printf("")

		var currentSum int
		if operator == "+" {
			currentSum = 0
		} else {
			currentSum = 1
		}
		// Since each inner list contains nth items and the nth -1 is the operator
		// We loop through each one up to the last one
		for index, currentNumber := range numbers[i] {
			// In case we get to the last item in each vertical list, we just continue
			// THat's just the operante + or *
			if index == len(numbers[i])-1 {
				break
			}

			switch operator {
			case "+":
				result := currentSum + input.ParseInt(currentNumber)
				currentSum = result
			case "*":
				result := currentSum * input.ParseInt(currentNumber)
				currentSum = result
			default:
				log.Printf("Invalid operator")
			}
		}
		sum += currentSum

	}
	return sum
}

// TODO: The problem right now is we don't know if the number is aligned left or right.
// Maybe we can read the inpuData differently, by looking at whole string and creating a custom split function that only uses a SINGLE space as separator
func buildNumber(numbers []string) []string {

	newNumbers := make([]string, len(numbers))
	biggest := 0
	operator := numbers[len(numbers)-1]
	// Since the numbers are read different then what we saw them, the key to read them correctly
	// Is to find what's the biggest number in each given list (in terms of raw text size)
	for _, number := range numbers {
		if len(number) > biggest {
			biggest = len(number)
		}
	}

	// What I'll try to do is fill up with 0's the blank spaces, so 123, 45, and 6 (vertically)
	// Would look like 123, 045 and 006
	// Then we can easily build a new number
	for index, item := range numbers {
		// We don't want to modify the + or * operator
		if index == len(numbers)-1 {
			break
		}
		var builder strings.Builder
		builder.Grow(biggest)
		sizeDifference := biggest - len(item)

		for i := 0; i < sizeDifference; i++ {
			builder.WriteString("0")
		}
		builder.WriteString(item)
		newNumbers[index] = builder.String()
	}

	log.Printf("List of rebuild numbers %s", newNumbers)

	// With the new list of numbers, we can now build new numbers
	// Suppose we have 3 numbers on each vertical list. Pick most right number from
	// first line and multiply it by 10 * pow(2) + most right number of second line * 10 *pow(1)
	// +  most right number of third line * pow(0). This will give us the number on that vertical.
	// Since we filled left numbers with 0 when they don't exist, the multiplication becomes 0.
	modifiedNumbers := []string{}
	for index := biggest - 1; index >= 0; index-- {
		// We don't want to modify the + or * operator
		var builder strings.Builder
		builder.Grow(biggest)

		for numIndex, item := range newNumbers {
			if numIndex == len(newNumbers)-1 {
				break
			}

			if item[index:index+1] == "0" {
				continue
			}

			number := input.ParseInt(item[index : index+1])
			log.Printf("Number that is going to be multiplied %d", number)
			builder.WriteString(fmt.Sprintf("%d", number))
		}
		modifiedNumbers = append(modifiedNumbers, builder.String())
		log.Printf("New number read vertically %s", builder.String())
	}
	newNumbers = append(newNumbers, operator)

	log.Printf("List of numbers after reading them verticall %s", modifiedNumbers)

	return newNumbers
}
