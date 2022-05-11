package main

import ("fmt"
)

func main() {
	var str,a string
	fmt.Scan(&str)
	for i, val := range str {
		a = a + string(val)
		if len(str)-1 != i {
			a = a + "*" 
		}
	}
	fmt.Print(a)
}
// fmt.Println(strings.Join(strings.Split(s, ""), "*"))