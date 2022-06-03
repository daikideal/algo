package main

import (
	"fmt"
)

func palindrome(x int) bool {
	// palindromeにならない特例:
	// 負の数: 数値の先頭が"-"なので、ひっくり返してもpalindromeにならない
	// 最後の桁が0の0ではない数: 最後が0だと、最初の桁の数字も0でないといけなくなる
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedX int
	for x > revertedX {
		revertedX = revertedX*10 + x%10
		x /= 10
	}

	return x == revertedX || x == revertedX/10
}

func main() {
	pattern := []int{
		-121,
		1,
		121,
		1221,
		1223,
		10,
	}

	for _, v := range pattern {
		fmt.Printf("input: %v, output: %v", v, palindrome(v))
		fmt.Println()
	}
}
