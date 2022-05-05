
package main

import "fmt"

func main() {
	var n,r int
	fmt.Scan(&n)
	for ;n > 0; {
		r = r * 10 + n % 10
		n /= 10
	}
	fmt.Print(r)
	}	



