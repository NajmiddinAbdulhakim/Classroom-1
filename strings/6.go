package main

import ("fmt"
// "strings"
"unicode"
)

func main() {
	var a string
	fmt.Scan(&a)
	// b := []rune(a)
	// fmt.Printf("%T",b)
	if len(a) >= 5 {
		gh := 0
		for _, val := range a {
			if unicode.Is(unicode.Latin,val) || unicode.IsDigit(val){
				gh++
			}
		}
		if gh == len(a) {
			fmt.Print("Ok")
		}else{
			fmt.Print("Wrong password")
		}
	}else{
		fmt.Print("Wrong password")
	}
}