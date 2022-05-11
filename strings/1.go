package main

import (
    "bufio"
    "fmt"
    "os"
	"unicode"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	tex := []rune(text)
	if unicode.IsUpper(tex[0]) && strings.HasSuffix(text,".") {
		fmt.Println("Right")
	}else {
		fmt.Println("Wrong")
	}
}
    