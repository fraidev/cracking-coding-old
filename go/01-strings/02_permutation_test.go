package strings

import "testing"

func TestIsPermutation(t *testing.T) {
	type args struct {
		str   string
		input string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{str: "ab", input: "dc"}, want: false},
		{args: args{str: "ab", input: "ba"}, want: true},
		{args: args{str: "abcce", input: "cabec"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.args.str+" and "+tt.args.input, func(t *testing.T) {
			if got := IsPermutation(tt.args.str, tt.args.input); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
			if got := IsPermutation2(tt.args.str, tt.args.input); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
