package main

import "fmt"

func main() {

	blorg := [8]int{1, 3, 5, 7, 11, 13, 17, 19}

	reverse(&blorg)

	fmt.Println(blorg)
}

func reverse(s *[8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
