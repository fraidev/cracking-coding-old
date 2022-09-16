package strings

import (
	"testing"

	"github.com/fraidev/cracking-coding/utils"
)

func TestZeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name      string
		args      args
		changedTo [][]int
	}{
		{
			name: "4x4", args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 0, 12, 13}, {14, 15, 16, 17}}},
			changedTo: [][]int{{1, 0, 3, 4}, {5, 0, 7, 8}, {0, 0, 0, 0}, {14, 0, 16, 17}},
		},
		{
			name: "3x4", args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 0, 12, 13}}},
			changedTo: [][]int{{1, 0, 3, 4}, {5, 0, 7, 8}, {0, 0, 0, 0}},
		},
		{
			name: "4x3", args: args{matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 0, 12}, {14, 15, 16}}},
			changedTo: [][]int{{1, 0, 3}, {5, 0, 7}, {0, 0, 0}, {14, 0, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ZeroMatrix(tt.args.matrix)

			if !utils.EqualMatrix(tt.args.matrix, tt.changedTo) {
				t.Errorf("matrix input = %v, changedTo %v", tt.args.matrix, tt.changedTo)
			}
		})
	}
}
