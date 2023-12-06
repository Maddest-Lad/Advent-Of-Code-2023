package main

import (
	"fmt"
	"internal/utils"
	"strconv"
	"unicode"
)

func main() {
	// Create Representation of Intput Matrix
	var grid [][]rune
	var sum int = 0

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
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// First Digit of The Number
			if unicode.IsDigit(grid[i][j]) {
				lastDigitIndex := j
				allNeighbors := make([]rune, 0)

				// Iterate Through the Entire Number
				for k := j; k < len(grid[0]) && unicode.IsDigit(grid[i][k]); k++ {
					lastDigitIndex = k
					allNeighbors = append(allNeighbors, getNeighbors(grid, i, k)...) // Variadic Operator (Unpacking / * in Python)
				}

				//println(string(grid[i][j:lastDigitIndex+1]), string(allNeighbors), containsSymbol(allNeighbors))

				if containsSymbol(allNeighbors) {
					val, _ := strconv.Atoi(string(grid[i][j : lastDigitIndex+1]))
					sum += val
				}

				// Skip Remainder of Valid Number
				j = lastDigitIndex
			}
		}
	}
	fmt.Println("The sum is", sum)
}

func getNeighbors(grid [][]rune, row int, col int) []rune {
	neighbors := make([]rune, 0)

	numRows := len(grid)
	numCols := len(grid[0])

	// Define relative positions of the neighbors
	positions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top Left, Top, Top Right
		{0, -1}, {0, 1}, // Left, <Skipping Middle>,Right
		{1, -1}, {1, 0}, {1, 1}, // Bottom Left, Bottom, Bottom Right
	}

	for _, pos := range positions {
		newRow := row + pos[0]
		newCol := col + pos[1]

		// Check if new row and column indices are within bounds
		if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
			neighbors = append(neighbors, grid[newRow][newCol])
		}
	}

	return neighbors
}

func containsSymbol(slice []rune) bool {
	for _, character := range slice {
		if !unicode.IsDigit(character) && character != '.' {
			return true
		}
	}
	return false
}
