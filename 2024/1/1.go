package main

import (
  "strings"
  "strconv"
  "sort"
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

  sort.Ints(left)
  sort.Ints(right)

  var distances []int
  for i := 0; i < len(left); i++ {
    distances = append(distances, abs(left[i] - right[i]))
  }

  var sum int
  for _, v := range distances {
    sum += v
  }
  fmt.Println(sum)
}



func abs(a int) int {
  if a >= 0 {
    return a
  }
  return -a
}
