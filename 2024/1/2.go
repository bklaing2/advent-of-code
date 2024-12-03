package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "bufio"
)


func main() {
  file, err := os.Open("1/input.txt")
  defer file.Close()

  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }

  var left, right []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    strings := strings.Fields(scanner.Text())
    leftNum, _ := strconv.Atoi(strings[0])
    rightNum, _ := strconv.Atoi(strings[1])

    left = append(left, leftNum)
    right = append(right, rightNum)
  }

  var leftCounts, rightCounts [99999]int
  for i := 0; i < len(left); i++ {
    leftNum := left[i]
    rightNum := right[i]

    leftCounts[leftNum]++
    rightCounts[rightNum]++
  }

  var sum int
  for i := 0; i < len(leftCounts); i++ {
    sum += i * leftCounts[i] * rightCounts[i]
  }

  fmt.Println(sum)
}



func abs(a int) int {
  if a >= 0 {
    return a
  }

  return -a
}



func countOccurrences(slice []int, target int) int {
  count := 0
  for _, num := range slice {
    if num == target {
      count++
    }
  }

  return count
}
