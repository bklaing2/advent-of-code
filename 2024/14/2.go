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
  maxRowSums, maxColSums := 0, 0

  for i := 0; i < 1000000; i++ {
    rowSums, colSums := make([]int, W), make([]int, H)

    for _, bot := range robots {
      x, y := bot.P.X + (bot.V.X * i), bot.P.Y + (bot.V.Y * i)
      x, y = x % W, y % H

      rowSums[x]++
      colSums[y]++
    }

    maxRowSum, maxColSum := 0, 0
    for _, sum := range rowSums {
      if sum > maxRowSum {
        maxRowSum = sum
      }
    }
    for _, sum := range colSums {
      if sum > maxColSum {
        maxColSum = sum
      }
    }

    if maxRowSums < maxRowSum && maxColSums < maxColSum {
      maxRowSums = maxRowSum
      maxColSums = maxColSum

      fmt.Println(i)
    }
  }
}
