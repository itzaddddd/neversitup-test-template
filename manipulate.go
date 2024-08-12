package main

import "fmt"

func permute(s string) (ans []string) {
	if len(s) == 0 {
		return
	}

	if len(s) == 1 {
		ans = append(ans, s)
		return
	}

	existedPermuted := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		curChar := s[i]
		remainChar := s[:i] + s[i+1:]

		remainCharPermuted := permute(remainChar)

		for j := 0; j < len(remainCharPermuted); j++ {
			newStr := fmt.Sprintf("%c%s", curChar, remainCharPermuted[j])

			if _, ok := existedPermuted[newStr]; !ok {
				existedPermuted[newStr] = true
			} else {
				continue
			}

			ans = append(ans, newStr)
		}
	}

	return
}
