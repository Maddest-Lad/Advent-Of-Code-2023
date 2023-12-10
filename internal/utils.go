package utils

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetScannerFromArgs() *bufio.Scanner {

	// Read The Input File
	filePath := os.Args[1]
	readFile, _ := os.Open(filePath)

	// Create Scanner to Read File Line-By-Line
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func ParseAocInput(line string) (string, []int) {
	// Attempt to Match `<String>: []int` into a reasonable input
	regex := regexp.MustCompile(`^([A-Za-z0-9]+):\s*((?:\d+\s*)+)$`) // https://regex101.com/
	matches := regex.FindStringSubmatch(line)

	return matches[1], ParseNumbers(matches[2])
}

// Convert a String Such as "83 86 6 31 17 9 48 53" to an Array of Valid Integers
func ParseNumbers(numberString string) []int {
	numbers := make([]int, 0)

	// Remove Spaces Around Data
	numberString = strings.TrimSpace(numberString)
	numberString = strings.Join(strings.Fields(numberString), " ")

	// Extract Numbers
	for _, numString := range strings.Split(numberString, " ") {
		numInt, _ := strconv.Atoi(numString)
		numbers = append(numbers, numInt)
	}

	return numbers
}
