package main

import "fmt"

func main() {
    type pos [2]int
    a := pos{4, 5}
    b := pos{4, 5}
    fmt.Println(a == b)
}