package main

import (
  "strings"
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
  fileName := "input.txt"
  if len(os.Args) > 1 {
    fileName = os.Args[1]
  }
  file, err := os.Open(fileName)

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }


  var plots [][]string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    row := strings.Split(scanner.Text(), "")
    plots = append(plots, row)
  }


  checked := make([]bool, len(plots) * len(plots[0]))
  fromIndex := func(i int) (int, int) {
    return i % len(plots[0]), i / len(plots[0])
  }

  // Logic
  sum := 0
  for i, c := range checked {
    if c {
      continue
    }

    x, y := fromIndex(i)
    perimeter, area := checkNeighbors(x, y, plots, &checked)
    sum += perimeter * area
  }

  fmt.Println(sum)
}


func checkNeighbors(x, y int, plots [][]string, checked *[]bool) (int, int) {
  plot := plots[y][x]
  perimeter, area := 0, 1

  (*checked)[y * len(plots[0]) + x] = true

  for _, offset := range OFFSETS {
    i, j := x + offset[0], y + offset[1]

    if i < 0 || j < 0 || i >= len(plots[0]) || j >= len(plots) {
      perimeter++
      continue
    }

    if plots[j][i] != plot {
      perimeter++
      continue
    }

    if (*checked)[j * len(plots[0]) + i] {
      continue
    }


    p, a := checkNeighbors(i, j, plots, checked)
    perimeter += p
    area += a
  }

  return perimeter, area
}
