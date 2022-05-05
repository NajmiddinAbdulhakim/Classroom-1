package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    switch {
    case n < 0 : fmt.Print("Число отрицательное")
    case n > 0 : fmt.Print("Число положительное")
    case n == 0 : fmt.Print("Ноль")
    }

}