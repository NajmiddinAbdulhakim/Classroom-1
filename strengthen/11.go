package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	switch  {
		case a % 10 == 1 && a != 11 : fmt.Printf("%d korova", a)
		case (a % 10 == 2 || a % 10 == 3 || a % 10 == 4) && (a != 12 && a != 13 && a != 14) : fmt.Printf("%d korovy",a)
		default : fmt.Printf("%d korov",a)
	}
}
