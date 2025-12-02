package input

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInputFile(dayNumberString string) ([]string, error) {
	currentPath, errPath := os.Getwd()
	if errPath != nil {
		log.Fatal("Failed to get current working directory")
	}
	FilePath := fmt.Sprintf("%s/%s/%s.txt", currentPath, dayNumberString, dayNumberString)
	content, err := os.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}
