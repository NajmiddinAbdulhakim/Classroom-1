package main

import ("fmt"
"strconv")
func main() {
	var num int
	fmt.Scan(&num)
	a := strconv.Itoa(num)
	var sum string
	for _,val := range a {
		kv,_ := strconv.Atoi(string(val))
		kv = kv * kv
		sum = sum + strconv.Itoa(kv)
	}
	fmt.Println(sum)

}