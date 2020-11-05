/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (42.31%)
 * Likes:    518
 * Dislikes: 0
 * Total Accepted:    77.6K
 * Total Submissions: 183.5K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true
 * 给定 word = "SEE", 返回 true
 * 给定 word = "ABCB", 返回 false
 *
 *
 *
 * 提示：
 *
 *
 * board 和 word 中只包含大写和小写英文字母。
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	var (
		rows, cols = len(board), len(board[0])
	)
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] {
				visited[row][col] = true
				if existDFS(board, visited, row, col, word[1:]) {
					return true
				}
				visited[row][col] = false
			}
		}
	}

	return false
}

func existDFS(board [][]byte, visited [][]bool, row, col int, word string) bool {
	if len(word) == 0 {
		return true
	}

	ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var r, c int
	for _, d := range ds {
		r, c = row+d[0], col+d[1]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !visited[r][c] && word[0] == board[r][c] {
			visited[r][c] = true
			if existDFS(board, visited, r, c, word[1:]) {
				return true
			}
			visited[r][c] = false
		}
	}

	return false
}

// func dfs(board [][]byte, row int, col int, used [][]bool, word string) bool {
// 	if len(word) == 0{
// 		return true
// 	}
// 	// 用于上下左右移动
// 	ds := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

// 	var x, y int
// 	for _, d := range ds{
// 		x, y = row+d[0], col+d[1]
// 		if x >=0 && x < len(board) && y>=0 && y<len(board[0]) && !used[x][y] && board[x][y] == word[0]{
// 			used[x][y] = true
// 			if dfs(board, x, y, used, word[1:]){
// 				return true
// 			}
// 			used[x][y] = false
// 		}
// 	}
// 	return false
// }
// @lc code=end
