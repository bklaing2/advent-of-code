package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "bufio"
)


func main() {
  file, err := os.Open("2/input.txt")
  defer file.Close()

  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  safeCount := 0

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    reportStrings := strings.Fields(scanner.Text())

    var report []int

    for _, string := range reportStrings {
      int, _ := strconv.Atoi(string)
      report = append(report, int)
    }

    if isSafe(report) {
      safeCount++
    }
  }

  fmt.Println(safeCount)
}



func isSafe(report []int) bool {
  directionCount := 0

  for i := 0; i < len(report) - 1; i++ {
    if report[i + 1] > report[i] {
      directionCount++
    } else if report[i + 1] < report[i] {
      directionCount--
      return false
    }

    dist := dist(report[i], report[i + 1])

    if dist < 1 || dist > 3 {
      return false
    }
  }

  return true
}




func dist(a, b int) int {
  diff := b - a
  return abs(diff)
}



func abs(a int) int {
  if a >= 0 {
    return a
  }
  return -a
}
