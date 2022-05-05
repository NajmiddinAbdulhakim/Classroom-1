package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	for i:=0; i < 5; i++ {
        if array[0] < array[i]{
			array[0] = array[i]
	    }
    }
	fmt.Print(array[0])
}