package main

import (
  "fmt"
  "os"
  "bufio"
)


var PATTERN = [8][2]int{
    {1, 0},
    {1, 1},
    {0, 1},
    {-1, 1},
    {-1, 0},
    {-1, -1},
    {0, -1},
    {1, -1},
  }

var XMAS = []rune{'X', 'M', 'A', 'S'}


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
      if c != XMAS[0] {
        continue
      }

      var xmas [4][8]bool
      for k := 0; k < 8; k++ {
        xmas[0][k] = true
      }

      for k := 1; k <= 3; k++ {
        for l := 0; l < 8; l++ {
          if !xmas[k-1][l] {
            continue
          }

          x := j + PATTERN[l][0] * k
          y := i + PATTERN[l][1] * k

          if x < 0 || x >= len(rows[i]) || y < 0 || y >= len(rows) {
            continue
          }

          xmas[k][l] = rows[y][x] == XMAS[k]
        }
      }


      for k := 0; k < 8; k++ {
        if xmas[3][k] {
          xmasCount++
        }
      }
    }
  }

  fmt.Println(xmasCount)
}

// 3  3  3
//  2 2 2
//   111
// 3210123
//   111
//  2 2 2
// 3  3  3
