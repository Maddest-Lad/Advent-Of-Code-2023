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

			if unicode.IsDigit(grid[i][j]) {
				adjacentSymbol := false
				lastDigitIndex := j

				for k := j; k < len(grid[0]); k++ {
					neighbors := getNeighbors(grid, i, k)

					if unicode.IsDigit(grid[i][k]) {
						lastDigitIndex = k

						fmt.Println(string(neighbors), string(grid[i][k]))
						if containsSymbol(neighbors) {
							adjacentSymbol = true
						}
					} else {
						break
					}
				}

				if adjacentSymbol {
					val, _ := strconv.Atoi(string(grid[i][j : lastDigitIndex+1]))
					sum += val
				}
				fmt.Println(string(grid[i][j:lastDigitIndex+1]), adjacentSymbol)
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
		if unicode.IsSymbol(character) {
			return true
		}
	}
	return false
}
