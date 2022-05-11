package main

import "fmt"

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return result
  }
func main() {
	var a string
	fmt.Scan(&a)
	b := Reverse(a)
	if a == b {
		fmt.Print("Палиндром")
	}else {
		fmt.Print("Нет")
	}
}

