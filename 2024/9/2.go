package main

import (
  "strconv"
  "fmt"
  "os"
)



func main() {
  // Setup
  b, err := os.ReadFile("9/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  s := string(b)
  var files, free []int

  for i := 0; i < len(s); i += 2 {
    file, _ := strconv.Atoi(string(s[i]))
    empty, _ := strconv.Atoi(string(s[i + 1]))
    files = append(files, file)
    free = append(free, empty)
  }


  // Logic
  leftIndex, freeIndex := 0, 0
  position := 0
  checkSum := 0

  for leftIndex < len(files) {
    if files[leftIndex] < 0 {
      position -= files[leftIndex]
    }

    for i := 0; i < files[leftIndex]; i++ {
      checkSum += position * leftIndex
      position++
    }


    for rightIndex := len(files) - 1; rightIndex > freeIndex && free[freeIndex] > 0; rightIndex-- {
      if files[rightIndex] > free[freeIndex] || files[rightIndex] < 0 {
        continue
      }

      free[freeIndex] -= files[rightIndex]

      for i := 0; i < files[rightIndex]; i++ {
        checkSum += position * rightIndex
        position++
      }

      files[rightIndex] = -files[rightIndex]
    }

    position += free[freeIndex];

    leftIndex++
    freeIndex++
  }

  fmt.Println(checkSum)
}
