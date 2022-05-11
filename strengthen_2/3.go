package main
import ("fmt"
"strconv"
)
func main() {
	var num string
	var max int
	fmt.Scan(&num)
	for _, v := range num {
		a,_ := strconv.Atoi(string(v))
		if a > max {
			max = a
		}
	}
	fmt.Print(max)
}