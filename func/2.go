package main
import "fmt"
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Print(minimumFromFour(a, b, c, d))
}
func minimumFromFour(a, b, c, d int) int {
	var h int
	switch {
	case a < b && a < c && a < b : 	h = a
	case b < c && b < d && b < a : 	h = b
	case c < a && c < b && c < d :	h = c
	case d < a && d < b && d < c : 	h = d	
	}
	return h
}