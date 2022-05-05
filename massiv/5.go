package main

import "fmt"

func main() {
	var a,cou int
	fmt.Scan(&a)
	mass := make([]int,a)
	for i := 0; i < a; i++ {
		fmt.Scan(&mass[i])
	}
	for i := 0; i < a; i++ {
		if mass[i] > 0 {
			cou++
		}
	}
	fmt.Print(cou)
}