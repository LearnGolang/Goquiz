package main

import "fmt"

func main() {
    var p [100]int
    var m interface{} = [...]int{99: 0}
    fmt.Println(p == m)
}