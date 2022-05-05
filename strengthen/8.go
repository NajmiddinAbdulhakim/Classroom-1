package main
import "fmt"
func main() {
	var a,b,cou int
	fmt.Scan(&a)
	arr := make([]int,a)
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
	}
    b = arr[0]
	for i := 0; i < a; i++ {
   		if b > arr[i] {
			b = arr[i]
		}
	}
    for i := 0 ; i < a ; i++ {
        if b == arr[i] {
			cou++
        }
    }
	fmt.Print(cou)
}
