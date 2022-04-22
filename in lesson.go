package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	fmt.Printf("It is %d hours %d minuts \n", num / 30,num % 30 )	

}