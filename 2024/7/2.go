package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "bufio"
)


func main() {
  // Setup
  file, err := os.Open("7/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var equations [][]int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    row := strings.Fields(scanner.Text())
    row[0] = row[0][:len(row[0]) - 1]

    ints := make([]int, len(row))
    for i, s := range row {
      num, _ := strconv.Atoi(s)
      ints[i] = num
    }

    equations = append(equations, ints)
  }

  // Logic
  sum := 0

  for _, equation := range equations {
    sum += solvable(equation)
  }

  fmt.Println(sum)
}



func solvable(equation []int) int {
  left := equation[0]
  right := equation[1:]

  l := len(right) - 1
  ops := make([]int, len(right))
  ops[0] = -1
  if checkPermutations(l, ops, left, right) {
    return left
  }

  return 0
}



func checkPermutations(level int, ops []int, left int, right []int) bool {
  if (level == 0) {
    ans := right[0]

    for i := 1; i < len(right); i++ {
      if ops[i] == 0 {
        ans += right[i]
      } else if ops[i] == 1 {
        ans *= right[i]
      } else {
        concat, _ := strconv.Atoi(strconv.Itoa(ans) + strconv.Itoa(right[i]))
        ans = concat
      }
    }

    return ans == left
  }


  for i := 0; i < 3; i++ {
    ops[level] = i
    if checkPermutations(level - 1, ops, left, right) {
      return true
    }
  }

  return false
}
