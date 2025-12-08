package input

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func ParseInt(number string) int {
	numberInt, _ := strconv.Atoi(number)
	return numberInt
}

func ParseInt64(number string) int64 {
	numberInt, _ := strconv.ParseInt(number, 10, 64)
	return numberInt
}
