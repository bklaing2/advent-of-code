package main

import (
  "strings"
  "fmt"
  "os"
  "bufio"
)


var OFFSETS = [][]int{
  {0, 1},
  {1, 0},
  {0, -1},
  {-1, 0},
}

var DIAGONAL_OFFSETS = [][]int{
  {1, 1},
  {1, -1},
  {-1, -1},
  {-1, 1},
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
    _, area, side := checkNeighbors(x, y, plots, &checked)
    sum += side * area
  }

  fmt.Println(sum)
}


func checkNeighbors(x, y int, plots [][]string, checked *[]bool) (int, int, int) {
  plot := plots[y][x]
  perimeter, area, corner := 0, 1, 0

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


    p, a, c := checkNeighbors(i, j, plots, checked)
    perimeter += p
    area += a
    corner += c
  }


  // Count corners
  for o := 0; o < 4; o++ {
    oa, ob := OFFSETS[o], OFFSETS[(o + 1) % 4]
    ia, ja := x + oa[0], y + oa[1]
    ib, jb := x + ob[0], y + ob[1]

    var a, b string


    if ia < 0 || ja < 0 || ia >= len(plots[0]) || ja >= len(plots) {
      a = "-1"
    } else {
      a = plots[ja][ia]
    }

    if ib < 0 || jb < 0 || ib >= len(plots[0]) || jb >= len(plots) {
      b = "-1"
    } else {
      b = plots[jb][ib]
    }

    if a != plot && b != plot {
      corner++
      continue
    }

    od := DIAGONAL_OFFSETS[o]
    id, jd := x + od[0], y + od[1]
    var d string
    if id < 0 || jd < 0 || id >= len(plots[0]) || jd >= len(plots) {
      d = "-1"
    } else {
      d = plots[jd][id]
    }

    if a == plot && b == plot && d != plot {
      corner++
      continue
    }
  }


  return perimeter, area, corner
}
