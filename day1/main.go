package main

import (
	"fmt"
	"os"
	"bufio"
)
var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNums(line string) int {
	var first, last int
	found := false

	for i := 0; i < len(line); i++ {
		digit := -1

		if line[i] >= '0' && line[i] <= '9' {
			digit = int(line[i] - '0')
		} else {
			for word, val := range wordToDigit {
				if i+len(word) <= len(line) && line[i:i+len(word)] == word {
					digit = val
				}
			}
		}

		if digit != -1 {
			if !found {
				first = digit
				found = true
			}
			last = digit // keep updating last
		}
	}

	// If no digits found (shouldn't happen per puzzle), return 0
	return 10*first + last
}

func main() {
	var fname string
	fmt.Print("Enter file: ")
	fmt.Scanln(&fname)
	fmt.Printf("Reading file ... %s\n", fname)
	file, err := os.Open(fname)
	if err != nil {
		fmt.Print("Failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		value := getNums(line)
		sum += value
	}
	fmt.Printf("The sum is ... %d", sum)
}
