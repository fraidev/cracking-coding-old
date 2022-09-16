package strings

import (
	"strings"
)

func Validate(input string) bool {
	str := strings.ReplaceAll(strings.ToLower(input), " ", "")
	hashMap := map[rune]int{}
	alreadyHasMiddleChar := false

	for _, r := range str {
		hashMap[r]++
	}

	for _, m := range hashMap {
		if m%2 == 1 {
			if alreadyHasMiddleChar {
				return false
			}
			alreadyHasMiddleChar = true
		}
	}

	return true
}
