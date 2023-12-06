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

	// Vars
	var sum int = 0     // Part 1
	var gearSum int = 0 // Part 2

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

	// Track Numbers Adjacent to Gears
	gearMap := make(map[[2]int][]int, 0)

	// Logic to Scan For Chars
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// First Digit of The Number
			if unicode.IsDigit(grid[i][j]) {
				lastDigitIndex := j
				allNeighbors := make([]rune, 0)
				neighboringGears := make([][2]int, 0)

				// Iterate Through the Entire Number
				for k := j; k < len(grid[0]) && unicode.IsDigit(grid[i][k]); k++ {
					lastDigitIndex = k
					neighbors, gears := getNeighbors(grid, i, k)

					// Variadic Operator (Unpacking / * in Python)
					allNeighbors = append(allNeighbors, neighbors...)
					neighboringGears = append(neighboringGears, gears...)
				}

				if containsSymbol(allNeighbors) {
					val, _ := strconv.Atoi(string(grid[i][j : lastDigitIndex+1]))
					sum += val

					// Part 2: Add Gears To Gear Map
					for _, gear := range neighboringGears {
						if !contains(gearMap[gear], val) {
							gearMap[gear] = append(gearMap[gear], val)
						}
					}
				}

				// Skip Remainder of Valid Number
				j = lastDigitIndex
			}
		}
	}

	// Part 2 Final Logic
	for _, numbers := range gearMap {
		// Check if there are exactly two numbers
		if len(numbers) == 2 {
			// Multiply the two numbers and add the product to the sum
			gearSum += numbers[0] * numbers[1]
		}
	}

	// Final Output
	fmt.Println("Part 1: The sum is", sum)
	fmt.Println("Part 2: The sum of all gear ratios is ", gearSum)
}

func getNeighbors(grid [][]rune, row int, col int) ([]rune, [][2]int) {
	neighbors := make([]rune, 0)
	gearLocations := make([][2]int, 0)

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

			// Check For Gears
			if grid[newRow][newCol] == '*' {
				gearLocations = append(gearLocations, [2]int{newRow, newCol})
			}
		}
	}

	return neighbors, gearLocations
}

func containsSymbol(slice []rune) bool {
	for _, character := range slice {
		if !unicode.IsDigit(character) && character != '.' {
			return true
		}
	}
	return false
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
