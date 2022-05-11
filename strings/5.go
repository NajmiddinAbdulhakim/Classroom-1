package main

import ("fmt"
"strings")
func main() {
	var a,c string
	fmt.Scan(&a)
	for _, v := range a {
		h := strings.Count(a,string(v))
		if h == 1 {
			c = c + string(v)
		}
	}
	fmt.Print(c)
}