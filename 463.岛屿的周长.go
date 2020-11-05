/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 *
 * https://leetcode-cn.com/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (68.42%)
 * Likes:    287
 * Dislikes: 0
 * Total Accepted:    33.1K
 * Total Submissions: 47.2K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * 给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
 *
 * 网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
 *
 * 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100
 * 。计算这个岛屿的周长。
 *
 *
 *
 * 示例 :
 *
 * 输入:
 * [[0,1,0,0],
 * ⁠[1,1,1,0],
 * ⁠[0,1,0,0],
 * ⁠[1,1,0,0]]
 *
 * 输出: 16
 *
 * 解释: 它的周长是下面图片中的 16 个黄色的边：
 *
 *
 *
 *
 */

 3+3+0+3+2+2+3

// @lc code=start
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	res := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1{
				isLandCnt := searchIsland(grid, row, col)
				curLength := 4 - isLandCnt
				res += curLength
				if isLandCnt == 0{
					return res
				}
			}

		}
	}

	return res
}

func searchIsland(grid [][]int, row int, col int) int {
	islandCnt := 0
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range ds {
		r, c := row+d[0], col+d[1]
		if r < len(grid) && r >= 0 && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 {
			islandCnt += 1
		}
	}
	return islandCnt
}

// 遍历网格地图， 如果碰到了岛屿， 判断该岛屿周围岛屿的数量， 如
// @lc code=end
