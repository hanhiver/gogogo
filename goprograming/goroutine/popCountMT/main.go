package main

import (
    "./popcount"
    "fmt"
    "sync"
)

func main() {
    var x uint64 = 186
    y := [...]uint64{123, 44, 565, 778, 1234, 55678, 12, 1, 44445, 63, 678, 2345, 43, 22}

    res := popcount.PopCount(x)
    fmt.Printf("%d\n", res)

    var wg sync.WaitGroup

    for _, item := range y {
        wg.Add(1) // sync.WaitGroup add one.

        go func(in uint64) {
            defer wg.Done()
            res := popcount.PopCount(in)
            fmt.Printf("PopCount(%d) = %d \n", in, res)
        }(item)
    }

    // Wait all go routines return.
    wg.Wait()

    fmt.Println("Done. ")
}
