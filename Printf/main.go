package main

import "fmt"

func main() {
	var v float64
	fmt.Scan(&v)
	if v < 1 {
		fmt.Printf("число %.2f не подходит",v)
	}else if v > 10000 {
		fmt.Printf("%e",v)
	}else if v > 0 && v < 10000{
		fmt.Printf("%.4f",v*v)
	}
}