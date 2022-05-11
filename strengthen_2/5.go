package main

import ("fmt"
"math")
var k,p,v float64
fmt.Scan(&k,&p,&v)
func T(k,p,v float64) float64 {
	return 6/math.Sqrt(k/(p*v))
}
func main() {

}