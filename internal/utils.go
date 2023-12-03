package utils

import (
	"bufio"
	"os"
)

func Get_scanner_from_args() *bufio.Scanner {

	// Read The Input File
	filePath := os.Args[1]
	readFile, _ := os.Open(filePath)

	// Create Scanner to Read File Line-By-Line
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	return scanner
}
