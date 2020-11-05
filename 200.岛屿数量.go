/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    575
 * Dislikes: 0
 * Total Accepted:    108.8K
 * Total Submissions: 219.8K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */

// @lc code=start
// [
// 	["1","1","1","1","0"],
// 	["1","1","0","1","0"],
// 	["1","1","0","0","0"],
// 	["0","0","0","0","0"]
// ]

var ds = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numIslands(grid [][]byte) int {
	var res int
	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[0]); j++{
			if grid[i][j] == '1'{
				res += 1
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	if grid[r][c] == '1'{
		grid[r][c] = '0'
		for _, d := range ds{
			x, y := r+d[0], c+d[1]
			if x >=0 && x < len(grid) && y>=0 && y< len(grid[0]){
				dfs(grid, x, y)
			}
		}
	}
	return
}

// func numIslands(grid [][]byte) int {
// 	res := 0
// 	for i := 0; i < len(grid); i++ {
// 		// 遍历行
// 		for j := 0; j < len(grid[i]) ; j++{
// 			// 遍历列
// 			if grid[i][j] == '1'{
// 				res ++
// 				SetIsLands(&grid, i, j)
// 			}
// 		}
// 	}
// 	return res
// }


// func SetIsLands(grid *[][]byte, x int , y int )  {
// 	if (*grid)[x][y] == byte('1'){
// 		(*grid)[x][y] = byte('0')  // 将节点设置成0

// 		if x - 1 >= 0 {
// 			SetIsLands(grid,x-1,y)
// 		}

// 		if x + 1 < len(*grid){
// 			SetIsLands(grid, x+1, y)
// 		}

// 		if y - 1 >= 0{
// 			SetIsLands(grid, x, y-1)
// 		}
// 		if y + 1 < len((*grid)[x]){
// 			SetIsLands(grid, x, y+1)
// 		}
// 	}
// 	return

// }
// @lc code=end
