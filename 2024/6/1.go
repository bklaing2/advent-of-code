package main

import (
  "strings"
  "fmt"
  "os"
  "bufio"
)


func main() {
  // Setup
  file, err := os.Open("6/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var obstacles, visited [][]bool
  var i, x, y int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var row, v []bool

    for j, o := range strings.Split(scanner.Text(), "") {
      if o == "^" {
        x = j
        y = i
      }

      row = append(row, o == "#")
      v = append(v, false)
    }

    obstacles = append(obstacles, row)
    visited = append(visited, v)
    i++;
  }

  // Logic
  var positionCount, facing int

  for move(&x, &y, &facing, obstacles) {
    if !visited[y][x] {
      positionCount++
      visited[y][x] = true
    }
  }

  fmt.Println(positionCount)
}



func move(x *int, y *int, facing *int, obstacles [][]bool) bool {
  prevX, prevY := *x, *y

  switch *facing {
  case 0:
    *y--
  case 1:
    *x++
  case 2:
    *y++
  case 3:
    *x--
  }

  if *x < 0 || *x >= len(obstacles[0]) || *y < 0 || *y >= len(obstacles) {
    return false
  }

  if obstacles[*y][*x] {
    *x, *y = prevX, prevY
    *facing = (*facing + 1) % 4
  }

  return true
}
