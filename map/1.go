package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	var a int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if _, ok := m[a]; ok {
			fmt.Print(m[a], " ")
		} else {
			m[a] = work(a) // ready func
			fmt.Print(m[a], " ")
		}
	}
}
