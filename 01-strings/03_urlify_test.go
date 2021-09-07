package strings

import "testing"

func TestURLify(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{input: "Mr John Smith    "}, want: "Mr%20John%20Smith"},
	}
	for _, tt := range tests {
		t.Run(tt.args.input+" "+tt.want, func(t *testing.T) {
			if got := URLify(tt.args.input); got != tt.want {
				t.Errorf("URLify() = %v, want %v", got, tt.want)
			}
		})
	}
}
