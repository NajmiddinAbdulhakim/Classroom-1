package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	for ;a < b; b-- {
		if b % 7 == 0 {
			fmt.Print(b)
			return
		}
	}
	fmt.Print("NO")
}

