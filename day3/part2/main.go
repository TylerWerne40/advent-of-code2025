package main

import (
	"bufio"
	"fmt"
  "math"
	"os"
)

type Coordinate struct {
  y int
  x int
}

func check_symbol(ch byte) bool {
  return ch == '*'
}

func check_byte(ch byte) bool {
  return ch != '.' && ('0' <= ch && ch <= '9') 
}

func find_adjacent(grid []string, coord Coordinate) []byte{
  ret_byte := make([]byte, 0)
  ch := grid[coord.y][coord.x]
  ret_byte = append(ret_byte, ch)
  for i := coord.x - 1; i >= 0; i-- {
    ch = grid[coord.y][i]
    if check_byte(ch) {
      ret_byte = append([]byte{ch}, ret_byte...)
    } else {
      break
    }
  }
  for i:= coord.x + 1; i<len(grid); i++ {
    ch = grid[coord.y][i]
    if check_byte(ch) {  
      ret_byte = append(ret_byte, ch)
    } else {
      break
    }
  }
  return ret_byte
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
  lines := make([]string, 0)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  rows := len(lines)
  cols := len(lines[0])
  directions := []struct{ dr, dc int }{
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1},           {0, 1},
    {1, -1},  {1, 0},  {1, 1},
  }
  matches := make(map[Coordinate]byte)
  
  gearNumbers := make(map[Coordinate]map[Coordinate]int)
  for r := 0; r < rows; r++ {
    for c := 0; c < cols; c++ {
      sym := lines[r][c]
      if !check_symbol(sym) {
        continue
      }
      gearCoord := Coordinate{y: r, x: c}
      for _, d := range directions {
        nr, nc := r + d.dr, c + d.dc
        if nr >= 0 && nr < rows && nc >= 0 && nc < cols { // Bounds check
          ch := lines[nr][nc]
          if !check_byte(ch) {
            continue
          }
          if gearNumbers[gearCoord] == nil {
            gearNumbers[gearCoord] = make(map[Coordinate]int)
          }
          // matches[Coordinate{nr, nc}] = ch
          row := lines[nr]
          left := nc
          for left > 0 && row[left-1] >= '0' && row[left-1] <= '9' {
            left--
          }
          right := nc
          for right+1 < len(row) && row[right+1] >= '0' && row[right+1] <= '9' {
            right++
          }
          numKey := Coordinate{y: nr, x: left}
          if _, exists := gearNumbers[gearCoord][numKey]; exists {
            continue
          }
          gearNum := find_adjacent(lines, numKey)
          num := 0
          for i, v := range gearNum {
            num += int(math.Pow10(len(gearNum) - 1 - i)) * int(v - '0')
          }
          gearNumbers[gearCoord][numKey] = num
        }
      }
      if yo, ok := gearNumbers[gearCoord]; ok {
        if len(yo) < 2 && len(yo) > 0{
          delete(gearNumbers, gearCoord)
        }
      }
    }
  }
  // have matches, now find nums adjacent (left-right) to matches
  nums := make([][]byte, 0) // remember to make inner bytes
  for coord := range matches {
    byte_arr := find_adjacent(lines, coord) // inner bytes made by func
    nums = append(nums, byte_arr)
  }
  sum := 0
  for _, numsMap := range gearNumbers {
    if len(numsMap) == 2 {
      product := 1
      for _, num := range numsMap {
        product *= num
      }
      sum += product
    }
  }
  fmt.Printf("The sum is %d\n", sum)
}
