/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (36.75%)
 * Likes:    414
 * Dislikes: 0
 * Total Accepted:    102.1K
 * Total Submissions: 277.8K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	return methodWithDp(obstacleGrid)
}


func methodWithDp(obstacleGrid [][]int) int {
	/*
		与不同路径解法基本相同， 只不过障碍物的位置， 将路径数量变成0
	*/

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1{
		return 0
	}

	dp := make([][]int, m)
	for i:=0; i<m; i++{
		rows := make([]int, n)
		dp[i] = rows
	}
	// 改变第一行的数组
	for j:=0; j<n; j++{
		if obstacleGrid[0][j] != 1{
			dp[0][j] = 1
		}else{
			break
		}
	}
	// 改变第一列的数组
	for i:=0; i<m; i++{
		if obstacleGrid[i][0] != 1{
			dp[i][0] = 1
		}else{
			break
		}
	}

	for i:=1; i<m; i++{
		for j:=1; j<n; j++{
			if obstacleGrid[i][j] != 1{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]

}
// @lc code=end
