package main

import (
	"fmt"
	"os"
	"bufio"
)

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
	var linenums []int = make([]int, 0, 20)
	for scanner.Scan() {
		var ints []int = make([]int, 2)
		line := scanner.Text()
		
		var sec bool = false
		for _, char := range line {
			if char >= '0' && char <= '9' {
				digit := int(char - '0')
				if !sec {
					ints[0] = digit
					ints[1] = digit
					sec = true
				} else {
					ints[1] = digit
				}
			}
		}
		var linenum int = 10*ints[0] + ints[1]
		linenums = append(linenums, linenum)
	}
	var sum int = 0
	for _, i := range linenums {
		sum += i
	}
	fmt.Printf("The sum is ... %d", sum)
}
