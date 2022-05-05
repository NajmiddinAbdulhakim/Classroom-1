package main

import "fmt"

func main() {
	var a,num,sum int
	fmt.Scan(&a)
	for ; a > 0; a-- {
		fmt.Scan(&num)
		if num > 9 && num < 100 && num % 8 == 0 {
			sum += num
		}
	}
	fmt.Print(sum)
}