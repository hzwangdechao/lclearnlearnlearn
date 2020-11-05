/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (62.72%)
 * Likes:    606
 * Dislikes: 0
 * Total Accepted:    51K
 * Total Submissions: 77.7K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 *
 * 一个数独的解法需遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 *
 *
 * 空白格用 '.' 表示。
 *
 *
 *
 * 一个数独。
 *
 *
 *
 * 答案被标成红色。
 *
 * Note:
 *
 *
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
	// 记录元素是否是数字
	var rows, cols [9][9]bool
	var boxes [3][3][9]bool
	// 记录空白元素的坐标
	var spaces [][2]int

	// 首先遍历整个二维数组， 将spaces 和 rows…… 进行填充、
	for row:=0; row<9; row++{
		for col:=0; col<9; col++{
			if board[row][col] == '.'{
				spaces = append(spaces, [2]int{row, col})
			}else{
				n := board[row][col] - '1'
				rows[row][n] = true
				cols[col][n] = true
				boxes[row/3][col/3][n] = true
			}
		}
	}

	var dfs func(int)bool
	dfs = func (pos int) bool {
		if pos == len(spaces){
			return true
		}
		r, c := spaces[pos][0], spaces[pos][1]
		for digit:=byte(0); digit<9; digit++{
			if !rows[r][digit] && !cols[c][digit] && !boxes[r/3][c/3][digit]{
				rows[r][digit] = true
				cols[c][digit] = true
				boxes[r/3][c/3][digit] = true
                board[r][c] = digit + '1'

					if dfs(pos+1){
						return true
					}
				rows[r][digit] = false
				cols[c][digit] = false
				boxes[r/3][c/3][digit] = false
			}
		}
		return false
	}

	dfs(0)

}
// @lc code=end
