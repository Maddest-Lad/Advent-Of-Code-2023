package main

import (
	"fmt"
	"internal/utils"
	"regexp"
	"strconv"
	"strings"
)

type diceGame struct {
	gameNumber int
	rounds     []gameRound
}

type gameRound struct {
	red, green, blue int
}

func main() {
	// Reuse Scanner Developer in Part 1
	scanner := utils.GetScannerFromArgs()

	// Vars
	var games []diceGame

	// Build Structs From Input
	for scanner.Scan() {
		var line string = scanner.Text()
		games = append(games, parseGame(line))
	}

	// Solve Puzzle
	var sum int

	for _, game := range games {
		if isValidGame(game) {
			sum += game.gameNumber
		}
	}

	fmt.Println(sum)
}

// The Elf would first like to know which games would have been possible if the bag contained only
// 12 red cubes, 13 green cubes, and 14 blue cubes
func isValidGame(game diceGame) bool {
	for _, round := range game.rounds {
		if round.red > 12 || round.green > 13 || round.blue > 14 {
			return false
		}
	}
	return true
}

func parseGame(line string) diceGame {
	// Splits Line Into 2 Chunks [Game, Rounds]
	parts := strings.Split(line, ":")

	// Extract Game Number
	gameNumber, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))

	// Parse Rounds
	rounds := parseRound(parts[1])

	return diceGame{gameNumber: gameNumber, rounds: rounds}
}

func parseRound(line string) []gameRound {
	parts := strings.Split(line, ";")
	rounds := make([]gameRound, len(parts)) // Created Zeroed Array of Our Rounds

	// Regex For Extracting # of Each Color `(Any Digit)<Single Whitespace>(Red|Green|Blue)`
	regex := regexp.MustCompile(`(\d+)\s(red|green|blue)`)

	for i, part := range parts {
		// Get All Matches
		matches := regex.FindAllStringSubmatch(part, -1)
		round := gameRound{red: 0, green: 0, blue: 0} // Create Empty Struct

		for _, match := range matches {

			count, _ := strconv.Atoi(match[1]) // Number Segment of Match

			switch match[2] {
			case "red":
				round.red = count
			case "green":
				round.green = count
			case "blue":
				round.blue = count
			}
		}
		rounds[i] = round
	}
	return rounds
}
