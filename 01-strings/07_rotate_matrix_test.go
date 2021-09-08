package strings

import (
	"testing"

	"github.com/fraidev/cracking-coding/utils"
)

func TestRotateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		changedTo [][]int
	}{
		{name: "work", args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 12, 13}, {14, 15, 16, 17}}}, 
		want: true, changedTo: [][]int{{14,9,5,1},{15,10,6,2},{16,12,7,3},{17,13,8,4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateMatrix(tt.args.matrix)
			if got != tt.want {
				t.Errorf("RotateMatrix() = %v, want %v", got, tt.want)
			}

			if !utils.EqualMatrix(tt.args.matrix, tt.changedTo) {
				t.Errorf("matrix input = %v, changedTo %v", tt.args.matrix, tt.changedTo)
			}
		})
	}
}
