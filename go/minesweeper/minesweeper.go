package minesweeper

func Annotate(board []string) []string {
	// Early return for empty boards
	if len(board) == 0 {
		return board
	}

	rows, cols := len(board), len(board[0])
	result := make([]string, rows)
	
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	runeBoard := make([][]rune, rows)
	for i, row := range board {
		runeBoard[i] = []rune(row)
	}

	// Process each cell on the board
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if runeBoard[r][c] == '*' {
				continue
			}
			
			// Count mines in adjacent cells
			count := 0
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && board[nr][nc] == '*' {
					count++
				}
			}
			
			// Update cell if mines were found
			if count > 0 {
				runeBoard[r][c] = rune('0' + count)
			}
		}
	}

	// Convert rune slices back to strings
	for i, row := range runeBoard {
		result[i] = string(row)
	}
	
	return result
}