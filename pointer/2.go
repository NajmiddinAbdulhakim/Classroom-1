package main
import "fmt"
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	test(&a, &b)
}
func test(x1 *int, x2 *int) {
	v := *x1
	*x1 = *x2
	*x2 = v
	fmt.Print(*x1,*x2)
 }
 
 
 
 
 