package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func checkTask(a interface{}) (string, bool) {
	switch v := a.(type) {
	case string:
		return `string`, true
	case float64:
		return `float64`, true
	default:
		fmt.Printf(`value=%v: %T`, v, v)
	}
	return "", false
}
func can(nf1, nf2 float64, op string) float64 {
	switch op {
	case `+`:
		return nf1 + nf2
	case `-`:
		return nf1 - nf2
	case `*`:
		return nf1 * nf2
	case `/`:
		return nf1 / nf2
	}
	return 0.0
}

func k(a, b, c interface{}) {

	aa, _ := checkTask(a)
	bb, _ := checkTask(b)
	cc, _ := checkTask(c)

	if aa == `float64` && bb == `float64` && cc == `string` {
		n1, first := a.(float64)
		n2, second := b.(float64)
		op, opt := c.(string)
		if first && second && opt {
			fmt.Printf("%.4f", can(n1, n2, op))
		}
	}
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	k(value1, value2, operation)            // все полученные значения имеют тип пустого интерфейса
}
