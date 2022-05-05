package main
import "fmt"
func main() {
	var a int
	one := 1
	two := 1
	count := 2
    empty := 0 
	fmt.Scan(&a)
	if count >= 2 {
		for  ; empty < a ; count++ {
			empty = one + two
			one = two
			two = empty
		}
    }
	if empty == a {
		fmt.Print(count)
	}else{
		fmt.Print(-1)
	}


}