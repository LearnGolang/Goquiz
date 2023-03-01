package main

import "fmt"

func main() {
    s := []string{"A", "B", "C"}

    t := s[:1]
    fmt.Println(&s[0] == &t[0])

    u := append(s[:1], s[2:]...)
    fmt.Println(&s[0] == &u[0])
}