package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading | ", err)
	}
	b := strings.Split(strings.ReplaceAll(strings.ReplaceAll(a, " ", ""), ",", "."), ";")
	g, _ := strconv.ParseFloat(b[0], 64)
	h, _ := strconv.ParseFloat(b[1], 64)
	fmt.Printf(strconv.FormatFloat(g/h, 'f', 4, 64))

}
