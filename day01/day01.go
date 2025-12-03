package day01

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"strconv"
)

// You could follow the instructions, but your recent required official North Pole secret entrance security training seminar
// taught you that the safe is actually a decoy.
// The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

var StartingPoint = 50

func Resolve() string {

	inputData, errInput := input.ReadInputFile("day01")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	password1 := findPassword(inputData, "1")
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	password2 := findPassword(inputData, "2")
	log.Printf("\n")

	return fmt.Sprintf("Result part 1: %d.\n Result Part 2: %d", password1, password2)
}

func findPassword(inputData []string, part string) int {
	var password = 0
	var directions []string
	var rotations []int

	for i := 0; i < len(inputData); i++ {
		directions = append(directions, inputData[i][0:1])
		amount, _ := strconv.Atoi(inputData[i][1:])
		rotations = append(rotations, amount)
	}

	var currentLocation = StartingPoint
	for i := 0; i < len(inputData); i++ {
		var newLocation int
		var addPassword int
		// Because I don't want to write two similar functions I just added if conditions to flow logic based on part 1 or 2
		if part == "1" {
			newLocation, addPassword = rotate(currentLocation, rotations[i], directions[i], "1")
		} else {
			newLocation, addPassword = rotate(currentLocation, rotations[i], directions[i], "2")
		}
		currentLocation = newLocation
		password += addPassword
	}

	return password
}

func rotate(currentLocation, rotation int, direction string, part string) (int, int) {
	oldLocation := currentLocation
	crossedZero := false
	var newLocation int
	password := 0

	// Every 100 rotations gets the gear to the same place.
	// So we just rotate whatever is left from dividing this rotations by 100
	if rotation > 100 {
		multiplier := rotation / 100
		rotation = rotation - (multiplier * 100)

		// If we're into part 2, each time the gear cross 0, then password is added by 1
		if part == "2" {
			password += multiplier
		}
	}

	switch direction {
	case "L":
		newLocation = oldLocation - rotation
		if newLocation < 0 {
			newLocation = 100 + newLocation

			// If we're into part 2, it's important to know if we crossed 0 or not, cause if we STARTED at 0, then the password cannot be added,
			// even though we went over (start) it, we didn't "crossed" it
			if part == "2" {
				// In case we start from the beginning, we don't want to add 1 to the password
				if oldLocation != 0 {
					password += 1
					crossedZero = true
				}
			}
		}
	case "R":
		newLocation = oldLocation + rotation
		if newLocation >= 100 {
			newLocation = newLocation - 100

			if part == "2" {
				// In case we start from the beginning, we don't want to add 1 to the password
				if oldLocation != 0 {
					password += 1
					crossedZero = true
				}
			}
		}
	default:
		log.Printf("Invalid Direction")
	}

	// For part 1, password is added only if it ends in 0.
	if part == "1" {
		if newLocation == 0 {
			password++
		}
		// For part 2, logic is a bit trickier. Location can be 0, but we can't have crossed 0, we have to STOP at it
	} else {
		if newLocation == 0 && !crossedZero {
			password++
		}
	}

	return newLocation, password
}
