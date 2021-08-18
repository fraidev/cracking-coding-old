package urlify

import (
	"fmt"
	"strings"
)

func URLify(input string) string {
	input = strings.TrimSpace(input)

	var sb strings.Builder

	for i, r := range input {
		if r == ' ' {
			fmt.Fprint(&sb, "%20")
		} else {
			fmt.Fprintf(&sb, "%c", input[i])
		}
	}

	return sb.String()
}
