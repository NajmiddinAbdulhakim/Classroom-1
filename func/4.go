package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(fibonacci(a))
}
func fibonacci(n int) int {
    if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}