package spiralmatrix

func SpiralMatrix(size int) [][]int {
	if size <= 0 {
		return [][]int{}
	}

	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	directions := [][]int{
		{0, 1}, 
		{1, 0}, 
		{0, -1},
		{-1, 0}, 
	}

	row, col := 0, 0        
	dirIdx := 0             
	start, end := 1, size*size 

	for num := start; num <= end; num++ {
		matrix[row][col] = num


		nextRow, nextCol := row+directions[dirIdx][0], col+directions[dirIdx][1]

		// Check if the next position is out of bounds or already filled
		if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size || matrix[nextRow][nextCol] != 0 {
			// Change direction
			dirIdx = (dirIdx + 1) % 4
			nextRow, nextCol = row+directions[dirIdx][0], col+directions[dirIdx][1]
		}

		// Update the position
		row, col = nextRow, nextCol
	}

	return matrix
}
