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
  leftIndex, rightIndex, freeIndex := 0, len(files) - 1, 0
  position := 0
  checkSum := 0

  for leftIndex < rightIndex {
    for i := 0; i < files[leftIndex]; i++ {
      checkSum += position * leftIndex
      position++
    }

    for i := 0; i < free[freeIndex]; i++ {
      checkSum += position * rightIndex
      files[rightIndex]--
      if files[rightIndex] == 0 {
        rightIndex--
      }
      position++
    }

    leftIndex++
    freeIndex++
  }


  for rightIndex < len(files) {
    for i := 0; i < files[rightIndex]; i++ {
      checkSum += position * rightIndex
      position++
    }

    rightIndex++ 
  }

  fmt.Println(checkSum)
}
