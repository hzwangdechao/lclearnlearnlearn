/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (70.70%)
 * Likes:    577
 * Dislikes: 0
 * Total Accepted:    69.6K
 * Total Submissions: 96.1K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 *
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例：
 *
 * 输入：4
 * 输出：[
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
 * 皇后可以攻击同一行、同一列、同一条斜线上的元素
 *
 *
 */

// @lc code=start
var res [][]string
func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	board := make([][]string, 0)


	for i:=0; i<n; i++{
		temp := make([]string, 0)
		for j:=0; j<n; j++{
			temp = append(temp, ".")
		}
		board = append(board, temp)
	}
	backtrack(board, 0)
	return res
}

// 结束条件： row等于board的长度
func backtrack(board [][]string, row int)  {
	if row == len(board){
		temp := make([]string, 0)
		for _, v := range board{
			str := ""
			for _, e := range v{
				str += e
			}
			temp = append(temp, str)
		}
		res = append(res, temp)
		return
	}

	n := len(board[row])
	for col:=0; col < n; col++{
		if !isValid(board, row, col){
			continue
		}

		// 做选择
		board[row][col] = "Q"
		// 进行下一个决策
		backtrack(board, row+1)
		// 撤销选择
		board[row][col] = "."
	}

}


// 判断第row行第col列是否可以放置皇后
func isValid(board [][]string, row int, col int) bool {
	n := len(board)

	// 检查第row行上是否有冲突的元素
	for i:=0; i<n; i++{
		if board[row][i] =="Q"{
			return false
		}
	}
	// 检查第col列上是否有冲突的元素
	for i:=0; i<n; i++{
		if board[i][col] == "Q"{
			return false
		}
	}
	// 检查右上方是否有冲突的元素
	for i,j:=row-1,col+1; i>=0 && j<n ; i, j=i-1, j+1{
		if board[i][j] == "Q"{
			return false
		}
	}
	// 检查左上方是否有冲突的元素
	for i,j:=row-1,col-1; i>=0 && j >=0; i, j=i-1,j-1{
		if board[i][j] == "Q"{
			return false
		}
	}

	return true
}
// @lc code=end
