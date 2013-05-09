package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() (func() int) {
    a := 1
    b := 2
    return func() int {
        tmpb := b
        b += a
        a = tmpb
        return tmpb
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
