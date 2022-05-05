package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
    mass := make([]int,a)
	for i := 0; i < a; i++ {
		fmt.Scan(&mass[i])
	}
	for i := 0; i < a; i++{
		if i % 2 == 0 {
			fmt.Print(mass[i]," ")
		}	
	}
}

