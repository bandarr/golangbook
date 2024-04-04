package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := 6

	c := appendInt(a, b)

	fmt.Println(c)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// there is room to grow.  extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // builtin
	}
	z[len(x)] = y
	return z
}
