package main
import "fmt"
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(vote(a, b, c))
}
func vote(x int, y int, z int) int {
	// var h int
    // if x == y && x == 0 {
	// 	h = 0
	// }else if x == z && x == 0 {
	// 	h = 0
	// }else if y == z && y == 0 {
	// 	h = 0
	// }else{
	// 	h = 1
	// }

	// return h
	return x & y | x & z | y & z 
}