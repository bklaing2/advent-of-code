package main

import (
  "fmt"
  "os"
  "bufio"
)


var PATTERN = [4][2]int{
    {-1, -1},
    {1, 1},
    {1, -1},
    {-1, 1},
}

const MS = int('M') + int('S')


func main() {
  file, err := os.Open("4/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var rows [][]rune

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    row := []rune(scanner.Text())
    rows = append(rows, row)
  }

  xmasCount := 0

  for i := range rows {
    for j, c := range rows[i] {
      if c != 'A' {
        continue
      }

      var xmas [4]int

      for l := 0; l < 4; l++ {
        x := j + PATTERN[l][0]
        y := i + PATTERN[l][1]

        if x < 0 || x >= len(rows[i]) || y < 0 || y >= len(rows) {
          continue
        }

        xmas[l] = int(rows[y][x])
      }

      if xmas[0] + xmas[1] == MS && xmas[2] + xmas[3] == MS {
        xmasCount++
      }
    }
  }

  fmt.Println(xmasCount)
}

// 1 1
//  0
// 1 1
