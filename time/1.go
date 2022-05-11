package main

import ("fmt"
"time"
)

func main() {
    var date string
	fmt.Scan(&date)
	dat,_ := time.Parse(time.RFC3339, date)
	fmt.Print(dat.Format(time.UnixDate))
}