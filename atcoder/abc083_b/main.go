package main

import (
	"fmt"
)

// 渡した数値の各桁の和を返す
//
// 参考: https://hacknote.jp/archives/28531/
func digitSum(num int) int {
	if num < 10 {
		return num
	}

	return (digitSum(num/10) + num%10)
}

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var answer int
	for i := 1; i <= n; i++ {
		sum := digitSum(i)

		if sum >= a && sum <= b {
			answer += i
		}
	}

	fmt.Print(answer)
}
