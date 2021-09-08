package strings

func RotateMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false
	}

	n := len(matrix)

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i] // save top

			// left -> top
			matrix[first][i] = matrix[last-offset][first]

			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			// top -> right
			matrix[i][last] = top //right <- saved top
		}
	}

	return true
}
