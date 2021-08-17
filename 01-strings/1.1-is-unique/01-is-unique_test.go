package isunique

import "testing"

func Test_isUnique(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "abc", want: true},
		{name: "abcd", want: true},
		{input: "abca", want: false},
		{input: "abcc", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsUnique(tt.input); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
			if got := IsUniqueSimple(tt.input); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
			if got := IsUniqueUsingSort(tt.input); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
