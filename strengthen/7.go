package main

import "fmt"

func main() {
	var a,cou int
	fmt.Scan(&a)
	arr := make([]int,a)
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < a; i++ {
		if arr[i] == 0 {
			cou++
		}
	}
	fmt.Print(cou)



}