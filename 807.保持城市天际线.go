/*
 * @lc app=leetcode.cn id=807 lang=golang
 *
 * [807] 保持城市天际线
 *
 * https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/description/
 *
 * algorithms
 * Medium (82.71%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    11.6K
 * Total Submissions: 14K
 * Testcase Example:  '[[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]'
 *
 * 在二维数组grid中，grid[i][j]代表位于某处的建筑物的高度。 我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。 高度 0
 * 也被认为是建筑物。
 *
 * 最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。
 * 城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。 请看下面的例子。
 *
 * 建筑物高度可以增加的最大总和是多少？
 *
 *
 * 例子：
 * 输入： grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
 * 输出： 35
 * 解释：
 * The grid is:
 * [ [3, 0, 8, 4],
 * ⁠ [2, 4, 5, 7],
 * ⁠ [9, 2, 6, 3],
 * ⁠ [0, 3, 1, 0] ]
 *
 * 从数组竖直方向（即顶部，底部）看“天际线”是：[9, 4, 8, 7]
 * 从水平水平方向（即左侧，右侧）看“天际线”是：[8, 7, 9, 3]
 *
 * 在不影响天际线的情况下对建筑物进行增高后，新数组如下：
 *
 * gridNew = [ [8, 4, 8, 7],
 * ⁠           [7, 4, 7, 7],
 * ⁠           [9, 4, 8, 7],
 * ⁠           [3, 3, 3, 3] ]
 *
 *
 * 说明:
 *
 *
 * 1 < grid.length = grid[0].length <= 50。
 * grid[i][j] 的高度范围是： [0, 100]。
 * 一座建筑物占据一个grid[i][j]：换言之，它们是 1 x 1 x grid[i][j] 的长方体。
 *
 *
 */

// @lc code=start
func maxIncreaseKeepingSkyline(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	// 首先求出从上部或底部看到的最大值列表
	top := make([]int, cols)
	for col := 0; col < cols; col++ {
		maxNum := 0
		for row := 0; row < rows; row++ {
			maxNum = max(grid[row][col], maxNum)
		}
		top[col] = maxNum
	}

	// 然后求出从左部或者右部看到的最大值列表
	left := make([]int, rows)
	for row := 0; row < rows; row++ {
		maxNum := 0
		for col := 0; col < cols; col++ {
			maxNum = max(grid[row][col], maxNum)
		}
		left[row] = maxNum
	}

	res := 0
	// grid[row][col] 单元格最大能增加到  min(top[col], left[row]) 高度
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			full := min(top[col], left[row])
			res += full - grid[row][col]
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
