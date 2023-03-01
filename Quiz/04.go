package main

import "fmt"

func main() {
    var m map[string]int
    delete(m, "oh noes!")
    fmt.Println(m)
}