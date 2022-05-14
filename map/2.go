package main

import (
	"fmt"
)

func main() {
	// groupCity := map[int][]string{
	// 	10:   []string{"hello", "world", "south"}, // города с населением 10-99 тыс. человек
	// 	100:  []string{"hello", "world"},          // города с населением 100-999 тыс. человек
	// 	1000: []string{"hello", "world", "south"}, // города с населением 1000 тыс. человек и более
	// }

	// cityPopulation := map[string]int{
	// 	"hello": 1,
	// 	"world": 2,
	// 	"south": 3,
	// }

	for k, _ := range cityPopulation {
		cou := 0
		for _, v := range groupCity[100] {
			if k == v {
				cou++
			}
		}
		if cou != 1 {

			delete(cityPopulation, k)
		}
	}
	fmt.Print(cityPopulation)
}

// for _, city := range append(groupCity[10], groupCity[1000]...) {
//     delete(cityPopulation, city)
// }
