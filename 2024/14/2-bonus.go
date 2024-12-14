package main

import (
  "strconv"
  "fmt"
  "os"
  "bufio"
  "regexp"
)

const (
  W = 101
  H = 103
  SECONDS = 7051
)

type Point struct {
  X, Y int
}

type Robot struct {
  P, V Point
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

  r, _ := regexp.Compile(`-?\d+`)

  var robots []Robot
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    var bot Robot

    line := scanner.Text()

    nums := r.FindAllString(line, -1)

    bot.P.X, _ = strconv.Atoi(nums[0])
    bot.P.Y, _ = strconv.Atoi(nums[1])
    bot.V.X, _ = strconv.Atoi(nums[2])
    bot.V.Y, _ = strconv.Atoi(nums[3])

    if bot.V.X < 0 {
      bot.V.X += W
    }
    if bot.V.Y < 0 {
      bot.V.Y += H
    }

    robots = append(robots, bot)
  }


  // Logic
  positions := make(map[Point]bool)

  for _, bot := range robots {
    x, y := bot.P.X + (bot.V.X * SECONDS), bot.P.Y + (bot.V.Y * SECONDS)
    x, y = x % W, y % H

    positions[Point{x, y}] = true
  }

  for y := 0; y < H; y++ {
    for x := 0; x < W; x++ {
      if positions[Point{x, y}] {
        fmt.Print("#")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}
