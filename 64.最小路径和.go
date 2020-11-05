/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (66.10%)
 * Likes:    574
 * Dislikes: 0
 * Total Accepted:    118.1K
 * Total Submissions: 175.8K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i:=0; i<rows; i++{
		row := make([]int, columns)
		dp[i] = row
	}

	// 先将第一行第一列的距离计算出来， 第一行第一列的最短距离呈一条直线
	for i:=0; i<rows; i++{
		for j:=0; j<columns; j++{
			if i == 0 && j == 0{
				dp[i][j] = grid[i][j]
			}else if i == 0{
				dp[i][j] = dp[0][j-1] + grid[i][j]
			}else if j == 0{
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}else{
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[rows-1][columns-1]

}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
// @lc code=end
