package main

import "fmt"

func main() {
	var a,num int
	cou := 1
	for fmt.Scan(&a) ; a != 0 ; fmt.Scan(&a) {
		if a == num {
			cou++
		}else {
			cou = 1
		}
		if a > num {
			num = a
		}
	}
	fmt.Print(cou)
}