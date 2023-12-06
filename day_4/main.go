package main

import (
	"fmt"
	"internal/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type scratchCard struct {
	cardNumber     int
	cardNumbers    []int
	winningNumbers []int
}

func main() {
	// Vars
	var cards = make([]scratchCard, 0)

	// Read Puzzle Input Into Grid
	scanner := utils.GetScannerFromArgs()

	for scanner.Scan() {
		// Read Input Into Objects
		parseCard(scanner.Text())
		cards = append(cards, parseCard(scanner.Text()))
	}

	// Outputs
	fmt.Println("Part 1: The Cards Are Worth", partOne(cards), "Points")
	fmt.Println("Part 2: There Are", partTwo(cards), "Total Scratchcards")
}

func parseCard(cardString string) scratchCard {
	// Split Input Line Into <"Card"|Card Number|Winning Numbers|Your Numbers>
	regex := regexp.MustCompile(`(Card\s+)(\d+):(.*?)\|\s*(.*)`)
	matches := regex.FindStringSubmatch(cardString)

	// Convert String Data To Integers
	cardNumber, _ := strconv.Atoi(matches[2])
	winningNumbers, cardNumbers := parseNumbers(matches[3]), parseNumbers(matches[4])

	return scratchCard{cardNumber: cardNumber, cardNumbers: cardNumbers, winningNumbers: winningNumbers}
}

// Convert a String Such as "83 86 6 31 17 9 48 53" to an Array of Valid Integers
func parseNumbers(numberString string) []int {
	numbers := make([]int, 0)

	// Remove Spaces Around Data
	numberString = strings.TrimSpace(numberString)
	numberString = strings.Replace(numberString, "  ", " ", -1) // Edge Case For Segments Like "  6"

	// Extract Numbers
	for _, numString := range strings.Split(numberString, " ") {
		numInt, _ := strconv.Atoi(numString)
		numbers = append(numbers, numInt)
	}

	return numbers
}

func getMatchingNumbers(card scratchCard) int {
	var counter = 0
	for _, number := range card.cardNumbers {
		if slices.Contains(card.winningNumbers, number) {
			counter++
		}
	}
	return counter
}

// Part 1 Puzzle Logic
func partOne(cards []scratchCard) int {
	var sum int = 0

	for _, card := range cards {
		sum += utils.PowInt(2, getMatchingNumbers(card)-1) // Account for One Matching Number Being 2^0
	}
	return sum
}

// Part 2 Puzzle Logic
func partTwo(originalDeck []scratchCard) int {
	cards := make([]scratchCard, len(originalDeck))
	copy(cards, originalDeck)

	var counter = 0

	for len(cards) > 0 {
		// Pop the first card from the queue
		currentCard := cards[0]
		cards = cards[1:]
		counter++

		matchingNumbers := getMatchingNumbers(currentCard)

		// Add copies of subsequent cards based on matching numbers
		for i := 1; i <= matchingNumbers; i++ {
			nextCardIndex := currentCard.cardNumber - 1 + i
			if nextCardIndex < len(originalDeck) {
				cards = append(cards, originalDeck[nextCardIndex])
			}
		}
	}
	return counter
}
