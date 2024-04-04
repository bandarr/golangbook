package main

import "fmt"

func main() {

	blorg := [...]int{1, 3, 5, 7, 11, 13, 17, 19}

	blarg := rotateLeft(blorg[:], 2)

	fmt.Println(blarg)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeft(s []int, x int) []int {
	return append(s[x:], s[:x]...)
}
