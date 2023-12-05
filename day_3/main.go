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
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Println(string(grid[i][j]))
			var neighbors = getNeighbors(grid, i, j)
			var _ = neighbors
		}
	}
}

func getNeighbors(grid [][]rune, row int, col int) []rune {
	var neighbors []rune

	// Prevent Out of Index Access
	rowLimit := len(grid)
	colLimit := len(grid[0])

	// Using Min/Max to Avoid Boundaries
	for i := utils.MaxOf(0, row-1); i < utils.MinOf(row+1, rowLimit); i++ {
		for j := utils.MaxOf(0, col-1); j < utils.MinOf(col+1, colLimit); j++ {
			fmt.Printf("[%d][%d] \n", i, j)
			if i != row && j != col {
				// Ignore Middle
				neighbors = append(neighbors, grid[i][j])
			}
		}
	}
	fmt.Printf("grid[%d][%d] %s: %s", row, col, string(grid[row][col]), string(neighbors))
	return neighbors
}
