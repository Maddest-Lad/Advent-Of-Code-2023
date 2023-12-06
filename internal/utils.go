package utils

import (
	"bufio"
	"math"
	"os"
)

func GetScannerFromArgs() *bufio.Scanner {

	// Read The Input File
	filePath := os.Args[1]
	readFile, _ := os.Open(filePath)

	// Create Scanner to Read File Line-By-Line
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
