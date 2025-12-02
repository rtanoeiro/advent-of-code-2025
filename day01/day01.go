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
	password := findPassword(inputData)
	return fmt.Sprintf("The password for the current day is %d", password)
}

func findPassword(inputData []string) int {
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
		newLocation, addPassword := rotate(currentLocation, rotations[i], directions[i])
		currentLocation = newLocation
		password += addPassword
	}

	return password
}

func rotate(currentLocation, rotation int, direction string) (int, int) {
	oldLocation := currentLocation
	crossedZero := false
	var newLocation int
	password := 0
	// Every 100 rotations gets the gear to the same place.
	// So we just rotate whatever is left from dividing this rotations by 100
	if rotation > 100 {
		multiplier := rotation / 100
		password += multiplier
		rotation = rotation - (multiplier * 100)
	}

	switch direction {
	case "L":
		newLocation = oldLocation - rotation
		if newLocation < 0 {
			newLocation = 100 + newLocation

			// In case we start from the beginning, we don't want to add 1 to the password
			if oldLocation != 0 {
				password += 1
				crossedZero = true
			}
		}
	case "R":
		newLocation = oldLocation + rotation
		if newLocation >= 100 {
			newLocation = newLocation - 100

			// In case we start from the beginning, we don't want to add 1 to the password
			if oldLocation != 0 {
				password += 1
				crossedZero = true
			}
		}
	default:
		log.Printf("Invalid Direction")
	}

	if newLocation == 0 && oldLocation != 0 && !crossedZero {
		password += 1
	}

	return newLocation, password
}
