package strings

func ZeroMatrix(matrix [][]int) {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 0 {
				defer SetZero(matrix, i, j)
			}
		}
	}
}

func SetZero(matrix [][]int, x, y int) {
	for i := 0; i < len(matrix[x]); i++ {
		matrix[x][i] = 0
	}

	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
}
