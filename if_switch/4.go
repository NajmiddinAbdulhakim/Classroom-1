package main
import "fmt"
func main(){
    	var num,a,b int
    fmt.Scan(&num)
    a = num % 10
	num /= 10
    a = a + num % 10
	num /= 10
    a = a + num % 10
	num /= 10
    
    b = num % 10
	num /= 10
	b =  b + num % 10
	num /= 10
	b = b + num % 10
	num /= 10
    //fmt.Print(a,b) 
    if a == b {
        fmt.Print("YES")
     
    }else {
        fmt.Print("NO")
    }

}