package strings

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{str: "Tact Coa"}, want: true},
		{args: args{str: "Tact boa"}, want: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.args.str, func(t *testing.T) {
			if got := Validate(tt.args.str); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
