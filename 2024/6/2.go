package main

import (
  "strings"
  "fmt"
  "os"
  "bufio"
)

var STEP = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

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
  var positionCount, loopCount, facing int
  var potentialObstacles [][2]int
  origX, origY, origF := x, y, facing

  for move(&x, &y, &facing, obstacles) {
    if !visited[y][x] {
      positionCount++
      visited[y][x] = true

      if x != origX || y != origY {
        potentialObstacles = append(potentialObstacles, [2]int{x, y})
      }
    }
  }

  loopCount = checkPositions(origX, origY, origF, obstacles, potentialObstacles)

  fmt.Println(positionCount, loopCount)
}


func move(x, y, facing *int, obstacles [][]bool) bool {
  prevX, prevY := *x, *y

  *x, *y = *x + STEP[*facing][0], *y + STEP[*facing][1]

  if *x < 0 || *x >= len(obstacles[0]) || *y < 0 || *y >= len(obstacles) {
    return false
  }

  if obstacles[*y][*x] {
    *x, *y = prevX, prevY
    *facing = (*facing + 1) % 4
  }

  return true
}


func checkPositions(x, y, facing int, obstacles [][]bool, potentialObstacles [][2]int) int {
  loopCount := 0

  for _, po := range potentialObstacles {
    o := make([][]bool, len(obstacles))
    for k := range obstacles {
      o[k] = append([]bool(nil), obstacles[k]...)
    }

    o[po[1]][po[0]] = true

    loopCount += checkIfLoops(x, y, facing, o)
  }

  return loopCount
}


func checkIfLoops(x, y, facing int, o [][]bool) int {
  sizeX, sizeY := len(o[0]), len(o)

  visited := make([]bool, 4 * sizeX * sizeY)

  for move(&x, &y, &facing, o) {
    i :=  facing * (sizeX * sizeY) + y * sizeX + x
    if visited[i] {
      return 1
    }

    visited[i] = true
  }

  return 0
}
