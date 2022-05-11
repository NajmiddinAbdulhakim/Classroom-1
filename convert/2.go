func adding(a, b string) int64 {
	x := ""
	z := ""
	for _, c := range a {
		if unicode.IsDigit(c) {
			x = x + string(c)
		}
	}
	for _, c := range b {
		if unicode.IsDigit(c) {
			z = z + string(c)
		}
	}
	c, _ := strconv.Atoi(x)
	d, _ := strconv.Atoi(z)
	return int64(c + d)

}

