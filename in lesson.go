package main

import "fmt"

func main() {

	// var c float32 = 9
	// switch c {
	// case 1 <= c && c <= 9:
	// 	fmt.Print("1 dan 9 gachca")
	// 	c--
	// 	fullthrough
	// case c == 10.2:
	// 	fmt.Print("Ok")
	// case 100 <= c && c <= 250:
	// 	fmt.Print("100 dan 200 gachca")
	// case 1000 <= c && c <= 6000:
	// 	c += 999
	// 	fmt.Print("1000 dan 6000 gachca")
	// 	fullthrough
	// default:
	// 	fmt.Print("default")

	// }

	// var n int
	// for fmt.Scan(&n);n != ; fmt.Scan(&n) {
	// var word string
	// fmt.Scanln(&word)
	// fmt.Println("Hello, World")
	// fmt.Println(word)
	

    // var hello string

    // hello = "Hello Go!"

    // var a int = 2019

    // fmt.Println(hello,a)
    // fmt.Println(a)

	// var symbol int32 = 'c'
	// fmt.Println(symbol)
	// a := 5
	// var b int = 6
	// var c = 7
	// fmt.Println(a,b,c)
	
	// var name string
    // var age int
    // fmt.Print("Введите имя: ")
    // fmt.Print("Введите возраст: ")
    // fmt.Scan(&name,&age)
    
	// fmt.Print("Ivan", 27) // Ivan27

	// fmt.Println("Ivan", 27) // Ivan 27

	// fmt.Print(33, 27) // 33 27	
	// var a int
    // fmt.Scan(&a)
    // fmt.Println(a*2+100)

	// var a,b,c int
    // fmt.Scan(&a)
    // b = a / 30
    // c = a % 30
    // fmt.Printf("It is %d hours %d minutes.",b,c)

	// const (
	// 	u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
	// 	v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
	// 	w         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
	// )
	// fmt.Print(u, v, w)
	
	// const (
	// 	a = iota + 1
	// 	_
	// 	b
	// 	c
	// 	d = c + 2
	// 	t
	// 	i
	// 	i2 = iota + 2
	// )
	
	// fmt.Println(a,b,c,d,t,i,i2)

	// if v := 9; v < 9 {
	// 	fmt.Print("ok")
	// }else{
	// 	fmt.Print("okk")
	// }

	// v := 50
	// switch v {
	// case 50:
	// 	fmt.Println(100)
	// 	fallthrough
	// case 42:
	// 	fmt.Println(42)
	// 	fallthrough
	// case 1:
	// 	fmt.Println(1)
	// 	fallthrough
	// default:
	// 	fmt.Println("default")
	
	// var a int
    // fmt.Scan(&a)
    // switch a {
    //     case a < 0: 
	// 		fmt.Println("Oтрицательное число")
    //     case a > 0: 
	// 		fmt.Println("Число положительное")
    //     case a == 0: 
	// 		fmt.Println("ноль")

	// var c uint32
	// fmt.Scan(&c)
	// switch {
	// case 1 <= c && c <= 9:
	// 	fmt.Println("от 1 до 9")
	// case 100 <= c && c <= 250:
	// 	fmt.Println("от 100 до 250")
	// case 1000 <= c && c <= 6000:
	// 	fmt.Println("от 1000 до 6000")

	// var num,a,b,c int
    // fmt.Scan(&num)
    // a = num / 100
    // b = num / 10 % 10
    // c = num % 10
    // if a == b || b == c {
    //     fmt.Print("NO")
    // }else {
    //     fmt.Print("YES")
    
	// fmt.Print(10000%10)
	// fmt.Print(1000%10)
	// fmt.Print(100%10)
	// fmt.Print(10%10)
	// fmt.Print(1%10)
	
	// var n int
	// fmt.Scan(&n)
    // if n >= 10 {
    //     n = n / 10
    // }
	// if n >= 10 {
    //     n = n / 10
    // }
	// if n >= 10 {
    //     n = n / 10
    // }
	// if n >= 10 {
    //     n = n / 10
    // }
	// if n >= 10 {
    //     n = n / 10
    // }
        
    // fmt.Print(n)

	// var num,a,b int
    // fmt.Scan(&num)
    // a = num % 10
	// num /= 10
    // a = a + num % 10
	// num /= 10
    // a = a + num % 10
	// num /= 10
    
    // b = num % 10
	// num /= 10
	// b =  b + num % 10
	// num /= 10
	// b = b + num % 10
	// num /= 10
    // fmt.Print(a,b) 

    // var s string
    // fmt.Scan(&s)
    // if s[0] + s[1] + s[2] == s[3] + s[4] + s[5] {
	// 	fmt.Println("YES")
	// 	} else {
	// 		fmt.Println("NO")
	// 	}
		
	// 	fmt.Print(s[0] + s[1] + s[2],s[3] + s[4] + s[5])
	var a int
	fmt.Scan(&a)
	fmt.Print(a%1000%9, a/1000%9)
}