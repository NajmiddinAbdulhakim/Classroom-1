package main
import "fmt"

func divide(a, b int) (int,error) {
	return a / b, nil
}

func main() {
    var x,z int
    fmt.Scan(&x,&z)
	a,err := divide(x,z)
	if err != nil {
		fmt.Println("ошибка")
	}else {
		fmt.Println(a)
	}
}
