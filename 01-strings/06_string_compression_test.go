package strings

import "testing"

func TestStringCompression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{input: "aabcccccaaa"}, want: "a2b1c5a3"},
		{args: args{input: "aabcccddaaa"}, want: "a2b1c3d2a3"},
		{args: args{input: "abc"}, want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			if got := StringCompression(tt.args.input); got != tt.want {
				t.Errorf("StringCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
