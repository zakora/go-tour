package main

import (
    "code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    res := make([][]uint8, dy)
    for y := 0 ; y < dy; y++ {
        tmpLine := make([]uint8, dx)
        for x := 0; x < dx; x++ {
            tmpLine[x] = uint8(x^y)
        }
        res[y] = tmpLine
    }
    return res
}

func main() {
    pic.Show(Pic)
}
