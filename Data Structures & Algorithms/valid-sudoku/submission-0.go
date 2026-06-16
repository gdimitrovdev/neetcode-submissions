func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var seen map[byte]bool = make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			_, ok := seen[board[i][j]]
			if ok {
				return false
			}

			seen[board[i][j]] = true
		}
	}

	for i := 0; i < 9; i++ {
		var seen map[byte]bool = make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			
			_, ok := seen[board[j][i]]
			if ok {
				return false
			}

			seen[board[j][i]] = true
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var seen map[byte]bool = make(map[byte]bool, 9)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[i*3 + x][j*3 + y] == '.' {
						continue
					}
					
					_, ok := seen[board[i*3 + x][j*3 + y]]
					if ok {
						return false
					}

					seen[board[i*3 + x][j*3 + y]] = true
				}
			}
		}
	}

	return true
}
