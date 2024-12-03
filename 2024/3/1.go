package main

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "regexp"
)


func main() {
  b, err := os.ReadFile("3/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  s := string(b)


  r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
  mulsSlice := r.FindAllString(s, -1)
  muls := strings.Join(mulsSlice, " ")

  r, _ = regexp.Compile(`\d+`)
  numsSlice := r.FindAllString(muls, -1)

  sum := 0
  for i := 0; i < len(numsSlice) - 1; i += 2 {
    a, _ := strconv.Atoi(numsSlice[i])
    b, _ := strconv.Atoi(numsSlice[i + 1]) 
    sum += a * b
  }

  fmt.Println(sum)
}
