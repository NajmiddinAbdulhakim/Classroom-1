package main
import "fmt"
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Print(minimumFromFour())
}
func minimumFromFour() int {
    var a,b,c,d int
    fmt.Scan(&a,&b,&c,&d)
	switch {
	case a < b && a < c && a < b : return a
	case b < c && b < d && b < a : return b
	case c < a && c < b && c < d : return c
	case d < a && d < b && d < c : return d	
	}
    return d
}