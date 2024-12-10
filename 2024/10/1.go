package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "bufio"
)


var OFFSETS = [][]int{
  {0, 1},
  {0, -1},
  {1, 0},
  {-1, 0},
}


func main() {
  // Setup
  file, err := os.Open("10/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var heights [][]int

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    row := strings.Split(scanner.Text(), "")

    ints := make([]int, len(row))
    for i, s := range row {
      num, _ := strconv.Atoi(s)
      ints[i] = num
    }

    heights = append(heights, ints)
  }


  // Logic
  sum := 0

  for y := 0; y < len(heights); y++ {
    for x := 0; x < len(heights[y]); x++ {
      if heights[y][x] != 0 {
        continue
      }

      visited := make([]bool, len(heights) * len(heights[0]))
      sum += tracePaths(x, y, heights, visited)
    }
  }


  fmt.Println(sum)
}


func tracePaths(x, y int, heights [][]int, visited []bool) int {
  height := heights[y][x]

  if height == 9 {
    i := y * len(heights[0]) + x
    seen := visited[i]
    visited[i] = true
    if seen {
      return 0
    }
    return 1
  }

  sum := 0

  for _, offset := range OFFSETS {
    i, j := x + offset[0], y + offset[1]

    if i < 0 || j < 0 || i >= len(heights[0]) || j >= len(heights) {
      continue
    }

    if heights[j][i] != height + 1 {
      continue
    }

    sum += tracePaths(i, j, heights, visited)
  }

  return sum
}
