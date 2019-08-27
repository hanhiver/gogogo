package main

import (
    "./popcount"
    "fmt"
    //"sync"
)

func main() {
    var x uint64 = 186
    y := [...]uint64{123, 44, 565, 778, 1234, 55678, 12, 1, 44445, 63, 678, 2345, 43, 22}

    res := popcount.PopCount(x)
    fmt.Printf("%d\n", res)

    inputs := make(chan uint64)
    results := make(chan int)

    go func() {
        for _, item := range y {
            inputs <- item
        }
        close(inputs)
    }()

    go func() {
        for x := range inputs {
            results <- popcount.PopCount(x)
        }
        close(results)
    }()

    /*
       for _, item := range y {
           go func(s uint64) {
               in := s
               res := popcount.PopCount(in)
               results <- res
           }(item)
       }
    */

    /*
       go func() {
           for {
               out, ok := <-results
               if !ok {
                   break
               }
               fmt.Printf("PopCount %d\n", out)
           }
           close(results)
       }()
    */

    for x := range results {
        fmt.Printf("PopCount %d\n", x)
    }

    //time.Sleep(1 * time.Second)
    fmt.Println("Done. ")
}
