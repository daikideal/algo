package main

import (
	"strings"
)

// 1: horizontal scannin
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
			return ""
	}
	
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
			for strings.Index(strs[i], prefix) != 0 {
					prefix = string([]rune(prefix)[:len(prefix)-1])
					if prefix == "" {
							return ""
					}
			}
	}

	return prefix
}

// 2: vertical scanning
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) < 1 {
// 			return ""
// 	}
	
// 	for i := 0; i < len(strs[0]); i++ {
// 			c := string([]rune(strs[0])[i])

// 			// 一致しなくなるまでstrs[0]から切り出す範囲を拡大する
// 			for j := 1; j < len(strs); j++ {
// 					if i == len(strs[j]) || string([]rune(strs[j])[i]) != c {
// 							return string([]rune(strs[0])[:i])
// 					}
// 			}
// 			fmt.Println(strs[0])
// 	}

// 	return strs[0]
// }
