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
  file, err := os.Open("5/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var rules [100][]int
  var updates [][]int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if scanner.Text() == "" {
      break
    }
    pages := strings.Split(scanner.Text(), "|")
    x, _ := strconv.Atoi(pages[0])
    y, _ := strconv.Atoi(pages[1])
    rules[x] = append(rules[x], y)
  }

  for scanner.Scan() {
    strings := strings.Split(scanner.Text(), ",")
    updates = append(updates, toIntArray(strings))
  }


  // Logic
  sum := 0

  for _, pages := range updates {
    if !updatesValid(pages, rules) {
      continue
    }

    sum += pages[len(pages) / 2]
  }

  fmt.Println(sum)
}




func updatesValid (pages []int, rules [100][]int) bool {
  var pageCount [100]int

  for _, page := range pages {
    for _, i := range rules[page] {
      if pageCount[i] > 0 {
        return false
      }
    }

    pageCount[page]++
  }

  return true
}




func toIntArray(strs []string) []int {
  ints := make([]int, len(strs))

  for i, s := range strs {
    num, _ := strconv.Atoi(s)
    ints[i] = num
  }

  return ints
}
