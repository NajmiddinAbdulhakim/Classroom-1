package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    mass := make([]int,n)
	for i := 0; i < n ;i++ {
		fmt.Scan(&mass[i])
	}
	fmt.Print(mass[3])
}