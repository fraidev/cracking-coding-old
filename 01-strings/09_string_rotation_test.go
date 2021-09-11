package strings

import "testing"

func TestStringRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s1: "waterbottle", s2: "erbottlewat"}, want: true},
		{args: args{s1: "waterbottle", s2: "erbottlewatf"}, want: false},
		{args: args{s1: "waterbottlef", s2: "erbottlewat"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.args.s1+" "+tt.args.s2, func(t *testing.T) {
			if got := StringRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("StringRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
