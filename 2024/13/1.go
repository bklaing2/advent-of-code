package main

import (
  "strconv"
  "fmt"
  "os"
  "bufio"
  "regexp"
  "math"
)


type Point struct {
  X, Y float64
}

type Machine struct {
  A, B, Prize Point
}


func main() {
  // Setup
  fileName := "input.txt"
  if len(os.Args) > 1 {
    fileName = os.Args[1]
  }
  file, err := os.Open(fileName)

  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  r, _ := regexp.Compile(`\d+`)

  var machines []Machine
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    var m Machine

    rowA := scanner.Text()
    scanner.Scan()
    rowB := scanner.Text()
    scanner.Scan()
    rowP := scanner.Text()
    scanner.Scan()

    a := r.FindAllString(rowA, -1)
    b := r.FindAllString(rowB, -1)
    p := r.FindAllString(rowP, -1)

    m.A.X, _ = strconv.ParseFloat(a[0], 64)
    m.A.Y, _ = strconv.ParseFloat(a[1], 64)
    m.B.X, _ = strconv.ParseFloat(b[0], 64)
    m.B.Y, _ = strconv.ParseFloat(b[1], 64)
    m.Prize.X, _ = strconv.ParseFloat(p[0], 64)
    m.Prize.Y, _ = strconv.ParseFloat(p[1], 64)

    machines = append(machines, m)
  }


  // Logic
  tokens := 0

  for _, m := range machines {
    b := (m.Prize.X - ((m.A.X * m.Prize.Y) / m.A.Y)) / (m.B.X - ((m.A.X * m.B.Y) / m.A.Y))
    a := (m.Prize.X - (b * m.B.X)) / m.A.X

    if !isRound(a) || !isRound(b) {
      continue
    }

    tokens += (int(math.Round(a)) * 3) + (int(math.Round(b)) * 1)
  }

  fmt.Println(tokens)
}



func isRound(f float64) bool {
  // Define a small tolerance
  const epsilon = 1e-9

  // Check if the fractional part is effectively zero
  return math.Abs(f-math.Round(f)) <= epsilon
}
