package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    rmap := make(map[string]int)
    words := strings.Fields(s)
    for _, word := range(words) {
        rmap[word] += 1
    }
    return rmap
}

func main() {
    wc.Test(WordCount)
}
