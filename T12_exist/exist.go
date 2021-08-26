package T12_exist

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if findWay(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func findWay(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[0] {
		return false
	}

	// 将已判断元素置空，避免再次判断
	board[i][j] = '0'

	res := findWay(board, word[1:], i+1, j) || findWay(board, word[1:], i-1, j) ||
		findWay(board, word[1:], i, j+1) || findWay(board, word[1:], i, j-1)
	board[i][j] = word[0]

	return res
}
