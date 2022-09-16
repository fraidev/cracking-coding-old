package strings

import (
	"github.com/fraidev/cracking-coding/utils"
)

func IsPermutation(str, input string) bool {

	if len(str) != len(input) {
		return false
	}

	return utils.SortString(str) == utils.SortString(input)
}

func IsPermutation2(str, input string) bool {

	if len(str) != len(input) {
		return false
	}

	letters := make(map[rune]int, 128)

	for _, r := range str {
		letters[r]++
	}

	for _, r := range input {
		letters[r]--
		if letters[r] < 0 {
			return false
		}
	}

	return true
}
