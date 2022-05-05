package main

import "fmt"

func main() {

	var x,p,y float32
	var	g int
	fmt.Scan(&x,&p,&y)
	for ; y > x  ; g++ {
		x += x / 100 * p
	}
		fmt.Print(g)

}