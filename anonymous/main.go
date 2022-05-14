package main

import (
	"fmt"
	"strconv"
)

func main() {
	// }

	fn := func(i uint) uint {
		if i == 0 {
			return 100
		} else {
			j := strconv.FormatUint(uint64(i), 10)
			srt := ""
			for _, v := range j {
				k, _ := strconv.Atoi(string(v))
				if k%2 != 1 && k != 0 {
					srt = srt + string(strconv.Itoa(k))
				}
			}
			l, _ := strconv.ParseUint(srt, 10, 64)
			if l == 0 {
				return 100
			}
			return uint(l)
		}
	}

}
