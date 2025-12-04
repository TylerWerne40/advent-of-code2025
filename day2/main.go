package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

var bagContents = map[string]int{
	"red":		12,
	"green":	13,
	"blue":		14,
}

func main() {
var fname string
	fmt.Print("Enter file: ")
	fmt.Scanln(&fname)
	fmt.Printf("Reading file ... %s\n", fname)
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var gameIndex int
	var sum int
	for gameIndex = 1; scanner.Scan(); gameIndex++ {
		line := scanner.Text()
		if strings.HasPrefix(line, "Game ") {
			colon := strings.Index(line, ":")
			if colon == -1 {
				continue
			}
			line = line[colon+2:]
		}
		pulls := strings.Split(line, ";")
		// possible := true
		var smallestBag = map[string]int {
			"blue":		0,
			"red":		0,
			"green":	0,
		}
		for _, pull := range pulls {
			cubes := strings.Split(pull, ",")
			for _, cube := range cubes {
				cube := strings.TrimSpace(cube)
				parts := strings.Fields(cube)
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				color = strings.Trim(color,",;")
				if count > smallestBag[color] {
					fmt.Printf("new %s color count %d\n", color, count)
					smallestBag[color] = count
				}
			}
		}
		power := 1
		if true {
			for k := range smallestBag {
				power = power*smallestBag[k]
			}
			sum += power
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	fmt.Printf("Sum of power is ... %d", sum)
}
