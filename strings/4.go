package main

import "fmt"

func main() {
	var a,c string	
	fmt.Scan(&a)
	for i, v := range a {
		if i % 2 != 0 {
			c = c + string(v)
		}
	}
	fmt.Print(c)
}