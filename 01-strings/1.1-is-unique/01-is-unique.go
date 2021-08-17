package isunique

import (
	"github.com/fraidev/cracking-coding/utils"
)

// IsUnique results if a string has all unique chars
// This solution is using HashMap data struct
// Big O (N)
func IsUnique(input string) bool {
	hashMap := make(map[rune]bool)

	for _, r := range input {
		if hashMap[r] {
			return false
		}

		hashMap[r] = true
	}

	return true
}

// IsUniqueSimple results if a string has all unique chars
// This solution is using none data struct
// Big O (N^2)
func IsUniqueSimple(input string) bool {
	for i := range input {
		checked := false

		for j := range input {
			if input[i] == input[j] {
				if checked {
					return false
				}

				checked = true
			}
		}
	}

	return true
}

// IsUniqueSimple results if a string has all unique chars
// This solution is using none data struct
// Big O (n log(n))
func IsUniqueUsingSort(input string) bool {
	lenInput := len(input)
	input = utils.SortString(input)

	for i := range input {
		if i == lenInput || input[i] == input[i+1] {
			return false
		}
	}

	return true
}
