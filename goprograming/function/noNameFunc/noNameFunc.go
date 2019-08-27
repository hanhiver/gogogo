package main

import (
    "fmt"
)

func main() {
    testFunc := selfAdd()

    fmt.Println(testFunc())
    fmt.Println(testFunc())
    fmt.Println(testFunc())

    testFunc = selfAdd()
    fmt.Println(testFunc())
    fmt.Println(testFunc())
}

func selfAdd() func() int {
    i := 0
    return func () int {
        res := i
        i += 1
        return res
    }
}

