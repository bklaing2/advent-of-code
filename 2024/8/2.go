package main

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
)



const S = 50


func main() {
  // Setup
  file, err := os.Open("8/input.txt")

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var antennas [123][][2]int
  r, _ := regexp.Compile(`[a-zA-Z0-9]`)

  scanner := bufio.NewScanner(file)

  for y := 0; scanner.Scan(); y++ {
    row := scanner.Text()

    for _, r := range r.FindAllStringIndex(row, -1) {
      x := r[0]
      i := row[x]
      antennas[i] = append(antennas[i], [2]int{x, y})
    }
  }


  // Logic
  locations := make([]bool, S * S)
  sum := 0

  for _, antenna := range antennas {
    findAntinodes(antenna, &sum, &locations)
  }

  fmt.Println(sum)
}



func findAntinodes(antenna [][2]int, sum *int, locations *[]bool) {
  for i := 0; i < len(antenna); i++ {
    for j := i + 1; j < len(antenna); j++ {
      a, b := antenna[i], antenna[j]
      dx, dy := antenna[i][0] - antenna[j][0], antenna[i][1] - antenna[j][1]

      c, d := [2]int{a[0], a[1]}, [2]int{b[0], b[1]}

      for inBounds(c) {
        if !(*locations)[toIndex(c)] {
          *sum++
          (*locations)[toIndex(c)] = true
        }

        c[0] += dx
        c[1] += dy
      }

      for inBounds(d) {
        if !(*locations)[toIndex(d)] {
          *sum++
          (*locations)[toIndex(d)] = true
        }

        d[0] -= dx
        d[1] -= dy
      }
    }
  }
}


func toIndex(p [2]int) int {
  return p[1] * S + p[0]
}

func inBounds(p [2]int) bool {
  return p[0] >= 0 && p[0] < S && p[1] >= 0 && p[1] < S
}
