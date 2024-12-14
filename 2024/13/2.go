package main

import (
  "strconv"
  "fmt"
  "os"
  "bufio"
  "regexp"
  "math/big"
)


const OFFSET = 10000000000000
const MaxPrecision = uint(2048)


type Point struct {
  X, Y *big.Float
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

    aX, _ := strconv.ParseFloat(a[0], 64)
    aY, _ := strconv.ParseFloat(a[1], 64)
    bX, _ := strconv.ParseFloat(b[0], 64)
    bY, _ := strconv.ParseFloat(b[1], 64)
    prizeX, _ := strconv.ParseFloat(p[0], 64)
    prizeY, _ := strconv.ParseFloat(p[1], 64)

    // Assign the large values to the struct fields with maximum precision
    m.A.X = new(big.Float).SetPrec(MaxPrecision).SetFloat64(aX)
    m.A.Y = new(big.Float).SetPrec(MaxPrecision).SetFloat64(aY)
    m.B.X = new(big.Float).SetPrec(MaxPrecision).SetFloat64(bX)
    m.B.Y = new(big.Float).SetPrec(MaxPrecision).SetFloat64(bY)
    m.Prize.X = new(big.Float).SetPrec(MaxPrecision).SetFloat64(prizeX)
    m.Prize.Y = new(big.Float).SetPrec(MaxPrecision).SetFloat64(prizeY)

    m.Prize.X.Add(m.Prize.X, big.NewFloat(OFFSET))
    m.Prize.Y.Add(m.Prize.Y, big.NewFloat(OFFSET))

    machines = append(machines, m)
  }


  // Logic
  tokens := new(big.Float).SetPrec(MaxPrecision).SetFloat64(0)

  for _, m := range machines {
    a, b := calculateBig(m)

    if !isRound(a) || !isRound(b) {
      continue
    }

    // Round a and b
    roundedA := round(a)
    roundedB := round(b)

    // Multiply rounded values by 3 and 1 respectively
    mulA := new(big.Float).SetPrec(MaxPrecision).Mul(roundedA, big.NewFloat(3)) // roundedA * 3
    mulB := new(big.Float).SetPrec(MaxPrecision).Mul(roundedB, big.NewFloat(1)) // roundedB * 1

    // Add to tokens
    tokens.Add(tokens, mulA)
    tokens.Add(tokens, mulB)
  }

  fmt.Println(tokens.Text('f', 0))
}

func round(f *big.Float) *big.Float {
  rounded := new(big.Float).Set(f)
  rounded.Add(rounded, new(big.Float).SetFloat64(0.5))
  roundedInt, _ := rounded.Int(nil)
  rounded.SetInt(roundedInt)

  return rounded
}

func isRound(f *big.Float) bool {
  const epsilon = 1e-9

  rounded := round(f)
  diff := new(big.Float).Sub(f, rounded)

  return diff.Abs(diff).Cmp(big.NewFloat(epsilon)) <= 0
}



func calculateBig(m Machine) (*big.Float, *big.Float) {
  // b := (m.Prize.X - ((m.A.X * m.Prize.Y) / m.A.Y)) / (m.B.X - ((m.A.X * m.B.Y) / m.A.Y))
  // (m.A.X * m.Prize.Y)
  temp1 := new(big.Float).Mul(m.A.X, m.Prize.Y)

  // (m.A.X * m.B.Y)
  temp2 := new(big.Float).Mul(m.A.X, m.B.Y)

  // (m.A.X * m.Prize.Y) / m.A.Y
  temp3 := new(big.Float).Quo(temp1, m.A.Y)

  // (m.A.X * m.B.Y) / m.A.Y
  temp4 := new(big.Float).Quo(temp2, m.A.Y)

  // m.Prize.X - ((m.A.X * m.Prize.Y) / m.A.Y)
  num := new(big.Float).Sub(m.Prize.X, temp3)

  // m.B.X - ((m.A.X * m.B.Y) / m.A.Y)
  den := new(big.Float).Sub(m.B.X, temp4)

  // b = num / den
  b := new(big.Float).Quo(num, den)

  // a := (m.Prize.X - (b * m.B.X)) / m.A.X
  // b * m.B.X
  temp5 := new(big.Float).Mul(b, m.B.X)

  // m.Prize.X - (b * m.B.X)
  numA := new(big.Float).Sub(m.Prize.X, temp5)

  // a = (m.Prize.X - (b * m.B.X)) / m.A.X
  a := new(big.Float).Quo(numA, m.A.X)

  return a, b
}
