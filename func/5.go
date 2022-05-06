package main
import "fmt"
func main() {
	a,b := sumInt(1,3)
	fmt.Print(a,b)
}
func sumInt(a ...int) (c int ,d int)  {
	for _, b := range a {
	    c++
		d += b
	}
	return c, d
}