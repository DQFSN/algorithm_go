package main

func Find(board [][]int, target int) int {
	rlen := len(board)
	if rlen == 0 {
		return false
	}
	clen := len(board[0])

	for r, c := 0, clen - 1; r < rlen && c >= 0; {
		if board[r][c] == target {
			return true
		}else if target >board[r][c] {
			r++
		}else {
			c--
		}
	}

	return false
}