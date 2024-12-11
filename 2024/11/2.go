package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
)


const NUM_BLINKS = 75


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
  cache := make(map[[2]int]int)

  for _, n := range strings.Fields(s) {
    num, _ := strconv.Atoi(string(n))
    stones = append(stones, num)
  }


  // Logic
  stoneCount := 0
  for _, s := range stones {
    stoneCount += blink(s, NUM_BLINKS, cache)
  }
  fmt.Println(stoneCount)
}


func blink(stone, n int, cache map[[2]int]int) int {
  if n == 0 {
    return 1
  }

  key := [2]int{n, stone}
  count, exists := cache[key]

  if exists {
    return count
  } else if stone == 0 {
    count = blink(1, n - 1, cache)
  } else if hasEvenNumberOfDigits(stone) {
    left, right := split(stone)
    count = blink(left, n - 1, cache) + blink(right, n - 1, cache)
  } else {
    count = blink(stone * 2024, n - 1, cache)
  }

  cache[key] = count
  return count
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
