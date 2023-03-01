package main

import "fmt"

func main() {
    m := map[string]int{"uno": 1}
    p := &m["uno"]
    *p = 2
    fmt.Println(m["uno"])
}