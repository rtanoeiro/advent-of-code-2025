package day04

import (
	"aoc-2025/input"
	"fmt"
	"log"
	"strings"
	"time"
)

var numPart1 = 2
var numPart2 = 12

type Roll struct {
	X int
	Y int
}

func Resolve() string {

	startTime := time.Now()
	inputData, errInput := input.ReadInputFile("day04")
	if errInput != nil {
		log.Fatal("Failed to read input file", errInput)
	}
	log.Printf("Starting solution for Part 1...")
	result1 := FindPaperRolls(inputData, false)
	log.Printf("\n")
	log.Printf("Starting solution for Part 2...")
	result2 := FindPaperRolls(inputData, true)
	log.Printf("\n")
	elapsedTime := time.Since(startTime)
	log.Printf("Results recieved %.6f seconds", elapsedTime.Seconds())
	return fmt.Sprintf(
		"The number of Paper Rolls for part 1 is %d.The result for Part2 is %d",
		result1,
		result2,
	)
}

// This function goes over the "paper roll" grid and check all 8 nearby places.
// If it finds there are less than 4 paper rolls around it, it adds to the sum.
// rebuild = false is used in part 1, so the forklift doesn't go again over the grid.
// rebuild == true is used in part 2, where we need to scan the grid again from the start
func FindPaperRolls(inputData []string, rebuild bool) int {

	updatedGrid := inputData
	lines := len(inputData)
	columns := len(inputData[0])
	// We split into two sums because during a rebuild, we start over again with a new grid.
	// Then currentLoopSum goes back to 0
	currentRunSum := 0
	finalSum := 0

	log.Printf("Showing Initial Grid")
	for i := 0; i < lines; i++ {
		log.Printf("%s", inputData[i][:])
	}

	// We'll use this variable to save all instances were we found a paper roll.
	// It will hold the coordinates for the paper roll that;s removed, and we'll go over them and replace with "x"
	var RebuildRollPost []Roll
	// Create rebuild functi
	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			// We only look if there's something when the current space is a paper roll
			if updatedGrid[i][j:j+1] == "@" {
				canRemove := checkSurroundings(i, j, updatedGrid)

				if canRemove {
					currentRunSum++
					RebuildRollPost = append(RebuildRollPost, Roll{X: i, Y: j})
				}
			}
		}

		if i == lines-1 && rebuild && currentRunSum > 0 {
			// We go back to 0 so we can restart and find more paper rolls
			finalSum += currentRunSum
			i = -1 // It will increase during the start of next iteration
			currentRunSum = 0
			updatedGrid = rebuildGrid(RebuildRollPost, updatedGrid)
			time.Sleep(time.Duration(100 * time.Millisecond))
		}
	}

	// When there's no rebuild of the grid, the currentRunSum contains the final sum.
	// When there's a rebuild, we sum that before resetting it
	if !rebuild {
		finalSum += currentRunSum
	}

	return finalSum
}

func checkSurroundings(i, j int, grid []string) bool {
	topLeft := getValueInGrid(i-1, j-1, grid)
	centerleft := getValueInGrid(i, j-1, grid)
	bottomLeft := getValueInGrid(i+1, j-1, grid)

	topCenter := getValueInGrid(i-1, j, grid)
	bottomCenter := getValueInGrid(i+1, j, grid)

	topRight := getValueInGrid(i-1, j+1, grid)
	centerRight := getValueInGrid(i, j+1, grid)
	bottomRight := getValueInGrid(i+1, j+1, grid)

	sumPapers := topLeft + centerleft + bottomLeft + topCenter + bottomCenter + topRight + centerRight + bottomRight
	if sumPapers < 4 {
		// log.Printf("Position %d,%d has a total of %d paper rolls", i, j, sumPapers)
		return true
	}
	return false
}

func getValueInGrid(i, j int, grid []string) int {
	if j < 0 {
		// log.Printf("Unable to look further left, skipping it")
		return 0
	}

	if i < 0 {
		// log.Printf("Unable to look further up, skipping it")
		return 0
	}

	// j + 1 means we went over the board to the right.
	// j - 1 means we went over the board on the left
	// i + 1 means we went over board on the bottom
	// i - 1 means we went over the board on the top
	if j+1 > len(grid[0]) {
		// log.Printf("Unable to look further right, skipping it")
		return 0
	}

	if i+1 > len(grid) {
		// log.Printf("Unable to look further down, skipping it")
		return 0
	}

	if grid[i][j:j+1] == "@" {
		return 1
	}

	return 0
}

// We have a list of rolls that will need to be removed. Each Roll has an X (line) and Y (column)
// This function goes on each Roll that needs removal, fetches the string in the current line.
// Rebuilds it based on which column need removal. After all lines are reuild, it returns the updated grid for a new forklift scan.
func rebuildGrid(rebuild []Roll, grid []string) []string {

	for i := 0; i < len(rebuild); i++ {
		currentLine := grid[rebuild[i].X]

		// Strings are immutable, so we create a new Builder in order to write the new string
		var builder strings.Builder
		builder.Grow(len(currentLine))

		// Current line represents the current X position of a Roll that needs rebuild
		for index := range currentLine {
			// If the current index in the old string is equal to the Y position (column),
			// It means a roll has to be removed, so we swap it with an "x" instead of writing whatever was in the old string
			if index == rebuild[i].Y {
				builder.WriteString(".")
				continue
			}
			// In case it doesn't match, it means there's nothing to rebuild there. So we just write whatever the old string had
			builder.Write([]byte{currentLine[index]})
		}

		// After we rebuild the current line with the "x", we need to assign that new string to the grid again, so the next iteraction
		// picks up that change
		grid[rebuild[i].X] = builder.String()
	}

	log.Printf("Showing Modified Grid")
	for i := 0; i < len(grid); i++ {
		log.Printf("%s", grid[i][:])
	}
	log.Printf("\n\n")

	return grid

}
