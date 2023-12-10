package main

import (
	"fmt"
	"internal/utils"
	"strconv"
)

type race struct {
	time int
	dist int
}

func main() {

	// Vars
	races := make([]race, 0)

	// Read Puzzle Input To Our Struct
	puzzleMap := make(map[string][]int)
	scanner := utils.GetScannerFromArgs()

	for scanner.Scan() {
		text, numbers := utils.ParseAocInput(scanner.Text())
		puzzleMap[text] = numbers
	}

	// Convert Input Into Our Objects
	for i := range puzzleMap["Time"] {
		time, distance := puzzleMap["Time"][i], puzzleMap["Distance"][i]
		races = append(races, race{time: time, dist: distance})
	}

	product, _ := partOne(races)
	_, sum := partTwo(races)

	fmt.Println("Part 1: The Product Of Every Win Condition is", product)
	fmt.Println("Part 2: There Are", sum, "Ways to Win")
}

func partOne(races []race) (product, sum int) {
	product = 0
	sum = 0

	for _, race := range races {
		count := 0
		for j := 1; j < race.time; j++ {
			//fmt.Println("Holding Button For", j, "Seconds, The Boat Will Travel", j*(race.time-j), "Milimeters")
			if j*(race.time-j) > race.dist {
				count++
				sum++
			}
		}
		if product == 0 {
			product = count
		} else {
			product *= count
		}
	}

	return
}

func partTwo(races []race) (product, sum int) {
	timeStr, distStr := "", ""

	for _, race := range races {
		timeStr += fmt.Sprint(race.time)
		distStr += fmt.Sprint(race.dist)
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)

	partTwoRace := race{time: time, dist: dist}

	return partOne([]race{partTwoRace})
}
