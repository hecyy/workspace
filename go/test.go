package main

import (
	"fmt"
	"math"
  "net/http"
	//"math/cmplx"
)

func printSlice(s string, x []int) {
  fmt.Printf("%s, len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func adder(s int) func(int) int {

  return func(x int) int {
    s+=x
    return s
  }
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type DaFloat float64

func (f DaFloat) Abs() float64 {
  if f  < 0 {
    f = -f
  }
  return float64(f)
}

type Abser interface {
  Abs() float64
}

type Hello struct {}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello!")
}

func main() {
  arr := [2]string{"abc", "def"}
  fmt.Println(len(arr))
}
