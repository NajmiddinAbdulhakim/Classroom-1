package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scan(&a,&b)
	var ma []int
	for ; 0 < a ; {
		c = c * 10 + a % 10
		a /= 10
	}
	for i := 0; 0 < c ; i++ {
		ma = append(ma, c % 10)
		c /= 10
	}
	for _,val := range ma {
		if val != b {
			fmt.Print(val)
		}
	}

}