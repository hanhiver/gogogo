package main

import (
    "fmt"
)

const (
    a = iota
    b
    c
    d = "ha"
    e
    f = 100
    g
    h = iota
    i
)

func main() {
    fmt.Println(a, b, c, d, e, f, g, h, i)
}

