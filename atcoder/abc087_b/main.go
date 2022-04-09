package main

import (
	"fmt"
)

// 全探索を用いた解答例
//
// 参考: https://qiita.com/drken/items/fd4e5e3630d0f5859067#%E7%AC%AC-4-%E5%95%8F--abc-087-b---coins-200-%E7%82%B9
func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	var count int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if (500*i + 100*j + 50*k) == x {
					count++
				}
			}
		}
	}

	fmt.Print(count)
}
