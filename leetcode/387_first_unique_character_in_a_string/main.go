package main

func firstUniqChar(s string) int {
	cache := make(map[rune]int)
	for _, v := range s {
		cache[v] += 1
	}

	for i, v := range s {
		if cache[v] == 1 {
			return i
		}
	}
	return -1
}
