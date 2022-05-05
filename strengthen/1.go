package main

import "fmt"

func main() {
	var n,s int
	fmt.Scan(&n)
	for ;n > 0; {
		s += n % 10
		n /= 10
	}
	fmt.Print(s)
}



