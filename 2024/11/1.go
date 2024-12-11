package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
)


const NUM_BLINKS = 25


func main() {
  // Setup
  fileName := "input.txt"
  if len(os.Args) > 1 {
    fileName = os.Args[1]
  }
  b, err := os.ReadFile("11/" + fileName)

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  s := string(b)

  var stones []int

  for _, n := range strings.Fields(s) {
    num, _ := strconv.Atoi(string(n))
    stones = append(stones, num)
  }


  // Logic
  stoneCount := 0
  for _, s := range stones {
    stoneCount += blink(s, NUM_BLINKS)
  }
  fmt.Println(stoneCount)
}


func blink(stone, n int) int {
  if n == 0 {
    return 1
  }

  if stone == 0 {
    return blink(1, n - 1)
  }

  if hasEvenNumberOfDigits(stone) {
    left, right := split(stone)
    return blink(left, n - 1) + blink(right, n - 1)
  }

  return blink(stone * 2024, n - 1)
}


func hasEvenNumberOfDigits(n int) bool {
  digitCount := len(strconv.Itoa(n))
  return digitCount % 2 == 0
}


func split(n int) (int, int) {
  s := strconv.Itoa(n)

  mid := len(s) / 2
  left, _ := strconv.Atoi(s[:mid])
  right, _ := strconv.Atoi(s[mid:])

  return left, right
}
