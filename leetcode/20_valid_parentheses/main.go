package main

func isValid(s string) bool {
	if len(s) < 2 {
			return false
	}

	pair := map[rune]rune {
			')': '(',
			'}': '{',
			']': '[',
	}

	stack := []rune{}
	for _, char := range s {
			if char == ')' || char == '}' || char == ']' {
					if len(stack) > 0 && stack[len(stack)-1] == pair[char] {
							stack = stack[:len(stack)-1]
							continue
					}
			}

			stack = append(stack, char)
	}
	
	if len(stack) > 0 {
			return false
	}
	return true
}
