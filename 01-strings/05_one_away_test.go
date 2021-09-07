package strings

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	type args struct {
		i1 string
		i2 string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{i1: "pale", i2: "ple"}, want: true},
		{args: args{i1: "pales", i2: "pale"}, want: true},
		{args: args{i1: "pale", i2: "bale"}, want: true},
		{args: args{i1: "pale", i2: "bake"}, want: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.args.i1+" "+tt.args.i2, func(t *testing.T) {
			if got := OneAway(tt.args.i1, tt.args.i2); got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
