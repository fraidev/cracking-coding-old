package strings

import "strings"

func StringRotation(s1, s2 string) bool {
	if len(s1) == len(s2) && len(s1) > 0 {
		s1s2 := s1 + s1
		return isSubstring(s1s2, s2)
	}

	return false
}

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}
