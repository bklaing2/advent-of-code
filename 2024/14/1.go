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
  // W = 11
  // H = 7
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
  quadrants := []int{0, 0, 0, 0}

  for _, bot := range robots {
    x, y := bot.P.X + (bot.V.X * 100), bot.P.Y + (bot.V.Y * 100)
    x, y = x % W, y % H

    q := getQuadrant(x, y)
    if q == -1 {
      continue
    }
    quadrants[q]++
  }

  safetyFactor := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
  fmt.Println(safetyFactor)
}


func getQuadrant(x, y int) int {
  midX, midY := W / 2, H / 2

  if y < midY {
    if x < midX {
      return 0
    } else if x > midX {
      return 1
    }
  } else if y > midY {
    if x < midX {
      return 2
    } else if x > midX {
      return 3
    }
  }

  return -1
}
