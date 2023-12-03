package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
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
		runes := []rune(line)

		// Loop From Left
		var left rune
		for index := range runes {
			if unicode.IsDigit(runes[index]) {
				left = runes[index]
				break
			}
		}

		// Loop From Right - A Bit Less Elegant
		var right rune
		for index := range runes {
			i := len(runes) - 1 - index // Convert Index to Be From The Right
			if unicode.IsDigit(runes[i]) {
				right = runes[i]
				break
			}
		}

		// Final Conversion to Int
		val, _ := strconv.Atoi(string(left) + string(right))
		sum += val
	}

	// Print Results
	fmt.Println(sum)
}
