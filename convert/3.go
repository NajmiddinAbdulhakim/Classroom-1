package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.Replace(a, ",", ".", -1)
	a = strings.Replace(a, " ", "", -1)
	// var g float64
	var b []string
	b = strings.Split(a, ";")
	fmt.Println(b[1])
	g, _ := strconv.ParseFloat(b[0], 64)
	h, _ := strconv.ParseFloat(b[1], 64)
	fmt.Println(g, h)
	fmt.Println(g / h)
	// for _, c := range b {
	// 	g = c
	// }
	// fmt.Printf("%f", 1149.6088607594936/1179.0666666666666)
}
