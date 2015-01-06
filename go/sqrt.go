package main

import (
	"fmt"
  "math"
	"math/rand"
)

func Sqrt(x float64) float64 {
	z := rand.Float64()
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	var i float64 = 2
	fmt.Printf("Sqrt(%d) = %f, math.Sqrt(%d) = %f\n",
		int(i), Sqrt(i), int(i), math.Sqrt(i))
  var j = 3
  v, ok := j.(float64)
  fmt.Println(v, ok)
}
