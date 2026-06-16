type NumMatrix struct {
	sumMatrix [][]int;
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	var sumMatrix [][]int = make([][]int, rows)
	for i, _ := range sumMatrix {
		sumMatrix[i] = make([]int, cols)
	}

	for i, row := range matrix {
		for j, _ := range row {
			if i == j && i == 0 {
				sumMatrix[i][j] = matrix[i][j]
				continue
			}

			if i == 0 {
				sumMatrix[i][j] = sumMatrix[i][j-1] + matrix[i][j]
				continue
			}

			if j == 0 {
				sumMatrix[i][j] = sumMatrix[i-1][j] + matrix[i][j]
				continue
			}

			sumMatrix[i][j] = sumMatrix[i-1][j] + sumMatrix[i][j-1] - sumMatrix[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{sumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == col1 && row1 == 0 {
		return this.sumMatrix[row2][col2]
	}

	if row1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row2][col1-1]
	}
	if col1 == 0 {
		return this.sumMatrix[row2][col2] - this.sumMatrix[row1-1][col2]
	}

	return this.sumMatrix[row2][col2] - this.sumMatrix[row2][col1-1] - this.sumMatrix[row1-1][col2] + this.sumMatrix[row1-1][col1-1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
