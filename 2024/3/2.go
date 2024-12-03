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

  r, _ := regexp.Compile(`mul\(\d+,\d+\)|do(n\'t)?\(\)`)
  instSlice := r.FindAllString(s, -1)
  inst := strings.Join(instSlice, " ")

  r, _ = regexp.Compile(`\d+|do(n\'t)?\(\)`)
  numsSlice := r.FindAllString(inst, -1)

  sum := 0
  doFlag := true
  for i := 0; i < len(numsSlice); i++ {
    inst := numsSlice[i]
    if inst == "do()" {
      doFlag = true
    } else if inst == "don't()" {
      doFlag = false
    } else if doFlag {
      a, _ := strconv.Atoi(numsSlice[i])
      b, _ := strconv.Atoi(numsSlice[i + 1])
      sum += a * b
      i++
    }
  }

  fmt.Println(sum)
}
