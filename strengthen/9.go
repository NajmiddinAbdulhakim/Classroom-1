
package main
import "fmt"
func main() {
	var n,s int
	fmt.Scan(&n)
	for ;n > 1; {
		s = s + n % 10
		n /= 10
	}
	for ;s > 0; {
		n += s % 10
		s /= 10
	}
	fmt.Print(n)
}





