package main

import ("fmt"
"strings")

func main() {
	var X,S string
	fmt.Scan(&X,&S)
	fmt.Print(strings.Index(X,S))
}