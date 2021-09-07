package strings

import (
	"fmt"
	"strings"
)

func StringCompression(input string) string {
	var sb strings.Builder
	counter := 0

	for i := range input {
		counter++

		if i+1 >= len(input) || input[i] != input[i+1] {
			fmt.Fprint(&sb, string(input[i]))
			fmt.Fprint(&sb, counter)
			counter = 0
		}
	}

	if sb.Len() > len(input) {
		return input
	}

	return sb.String()

}
