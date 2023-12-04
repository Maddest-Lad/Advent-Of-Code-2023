package main

import (
	"fmt"
	"internal/utils"
)

func main() {
	// Create Representation of Intput Matrix
	var grid [][]rune

	// Read Puzzle Input Into Grid
	scanner := utils.GetScannerFromArgs()

	for scanner.Scan() {
		var lineRunes []rune
		line := scanner.Text()

		// Convert line to slice of runes
		for _, char := range line {
			lineRunes = append(lineRunes, char)
		}

		// Append to grid
		grid = append(grid, lineRunes)
	}

	// Logic to Scan For Chars
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf(string(char))
		}
		fmt.Println()
	}
}
