/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (42.17%)
 * Likes:    349
 * Dislikes: 0
 * Total Accepted:    66K
 * Total Submissions: 156.4K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 *
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 *
 * 示例:
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * 运行你的函数后，矩阵变为：
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * 解释:
 *
 * 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 */

// @lc code=start

func solve(board [][]byte)  {
	SolveMethodWithDfs(board)
}

func SolveMethodWithDfs(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {

		return
	}
	rows, cols := len(board), len(board[0])
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0)
		dfs(row, cols-1)
	}

	for col := 1; col < cols-1; col++ {
		dfs(0, col)
		dfs(rows-1, col)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'A' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}

}

func SolveMethodWithBfs(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	ds := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 与深度优先搜索相似， 寻找边界上的O元素， 将元素的位置放到队列中， 循环遍历队列， 将与边界元素O相邻的元素继续放到队列中
	rows, cols := len(board), len(board[0])

	var Q [][]int

	// 遍历每一行的第一个元素和最后一个元素
	for row := 0; row < rows; row++ {
		if board[row][0] == 'O' {
			Q = append(Q, []int{row, 0})
		}
		if board[row][cols-1] == 'O' {
			Q = append(Q, []int{row, cols - 1})
		}
	}
	// 遍历第一行和最后一行剩余的元素
	for col := 1; col < cols-1; col++ {
		if board[0][col] == 'O' {
			Q = append(Q, []int{0, col})
		}
		if board[rows-1][col] == 'O' {
			Q = append(Q, []int{rows - 1, col})
		}
	}

	// 队列中的元素存放的是边界上的元素或者与边界元素相邻的元素的坐标
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		x, y := cur[0], cur[1]

		board[x][y] = 'A'      // 将元素暂时变为A
		for _, d := range ds { // 继续遍历元素相邻的元素
			r, c := x+d[0], y+d[1]

			if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
				continue
			}
			Q = append(Q, []int{r, c})
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'A' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}

// func SolveMethodWithUnionFind(board [][]byte) {
// 	if len(board) == 0 || len(board[0]) == 0 {
// 		return
// 	}
// 	ds := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
// 	rows, cols := len(board), len(board[0])
// 	product := rows + cols
// 	uf := make([]int, product+1)
// 	var find func(x int) int
// 	find = func(x int) int {
// 		if uf[x] != x {
// 			uf[x] = find(uf[x])
// 		}
// 		return uf[x]
// 	}

// 	var union func(x, y int)
// 	union = func(x, y int) {
// 		uf[find(x)] = find(y)
// 	}

// 	var node func(x, y int)int
// 	node = func (x, y int) int {
// 		return x*cols+y
// 	}

// 	for row := 0; row < rows; row++ {
// 		for col := 0; col < cols; col++ {
// 			if board[row][col] == 'O' {
// 				if row == 0 || row == rows-1 || col == 0 || col == cols-1 {
// 					// 如果O是边界的元素
// 					union(node(row, col), product)
// 				} else {
// 					for _, d := range ds {
// 						x := row+d[0]
// 						y := col+d[1]
// 						if (x > 0 || x < rows || y > 0 || y<cols) && board[x][y] == 'O'{
// 							union(node(row, col), node(x, y))
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if find(product) == find(i*cols+j) {
// 				board[i][j] = 'O'
// 			} else {
// 				board[i][j] = 'X'
// 			}
// 		}
// 	}

// }

// @lc code=end
