package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// replacer for converting the number style
	number_replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)

	var sum int

	// Read The Input File
	filePath := os.Args[1]
	readFile, _ := os.Open(filePath)

	// Create Scanner to Read File Line-By-Line
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	// For Every Line in Our Input File
	for scanner.Scan() {
		// Get the Text Line and Convert it to a Slice of Runes
		var line string = scanner.Text()

		// Replace Each Instance
		line = number_replacer.Replace(line)
		line = number_replacer.Replace(line)

		// Convert to Runes
		runes := []rune(line)

		// Extract Number Values
		var left rune = get_left(runes)
		var right rune = get_right(runes)

		// Final Conversion to Int
		val, _ := strconv.Atoi(string(left) + string(right))
		sum += val
	}

	// Print Results
	fmt.Println(sum)
}

// Loop From Left
func get_left(runes []rune) rune {
	for index := range runes {
		if unicode.IsDigit(runes[index]) {
			return runes[index]
		}
	}
	return -1
}

// Loop From Right - A Bit Less Elegant
func get_right(runes []rune) rune {
	for index := range runes {
		i := len(runes) - 1 - index // Convert Index to Be From The Right
		if unicode.IsDigit(runes[i]) {
			return runes[i]
		}
	}
	return -1
}
