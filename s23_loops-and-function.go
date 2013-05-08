package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (z float64) {
    delta := 0.001
    z = x
    oldZ := z + 2*delta
    for math.Abs(z - oldZ) > delta {
        oldZ = z
        z = z - (z*z - x)/(2*z)
    }
    return
}

func main() {
    fmt.Println(Sqrt(4), Sqrt(25))
}
