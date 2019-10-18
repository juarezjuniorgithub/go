package main

import (
    "fmt"
    "math"
)

const s string = "Constants Example"

func main() {

    fmt.Println(s)

    const n = 800000000

    const d = 8e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}