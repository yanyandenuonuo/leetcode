/*
 * @lc app=leetcode.cn id=1275 lang=golang
 *
 * [1275] 找出井字棋的获胜者
 */

// @lc code=start
func tictactoe(moves [][]int) string {
	// solution: 模拟
	if len(moves) < 3 {
		return "Pending"
	}

	chessBoard := [3][3]byte{}
	for idx, move := range moves {
		if idx&0x01 == 0x00 {
			// A
			chessBoard[move[0]][move[1]] = 'X'
		} else {
			// B
			chessBoard[move[0]][move[1]] = 'O'
		}
	}

	resList := []string{
		// row
		string(chessBoard[0][0]) + string(chessBoard[0][1]) + string(chessBoard[0][2]),
		string(chessBoard[1][0]) + string(chessBoard[1][1]) + string(chessBoard[1][2]),
		string(chessBoard[2][0]) + string(chessBoard[2][1]) + string(chessBoard[2][2]),

		// column
		string(chessBoard[0][0]) + string(chessBoard[1][0]) + string(chessBoard[2][0]),
		string(chessBoard[0][1]) + string(chessBoard[1][1]) + string(chessBoard[2][1]),
		string(chessBoard[0][2]) + string(chessBoard[1][2]) + string(chessBoard[2][2]),

		// intersect
		string(chessBoard[0][0]) + string(chessBoard[1][1]) + string(chessBoard[2][2]),
		string(chessBoard[0][2]) + string(chessBoard[1][1]) + string(chessBoard[2][0]),
	}

	for _, res := range resList {
		switch res {
		case "XXX":
			return "A"
		case "OOO":
			return "B"
		}
	}

	if len(moves) != 9 {
		return "Pending"
	}

	return "Draw"
}

// @lc code=end

