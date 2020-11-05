/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (62.22%)
 * Likes:    692
 * Dislikes: 0
 * Total Accepted:    148.9K
 * Total Submissions: 239.2K
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 *
 *
 *
 * 示例 1:
 *
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 2:
 *
 * 输入: m = 7, n = 3
 * 输出: 28
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 100
 * 题目数据保证答案小于等于 2 * 10 ^ 9
 *
 *
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	return methodWithDp(m, n)
}

// 使用dp数组进行求解
func methodWithDp(m int, n int) int {
	/*
		首先构造一个dp数组
		1 1 1 1
		1 0 0 0
		1 0 0 0

		dp[i][j] = dp[i-1][j] + dp[i][j-1]

		最终要求的结果就是dp[m-1][n-1]
	*/
	dp := make([][]int, m)
	for i:=0; i<m; i++{
		rows := make([]int, n)
		if i == 0{
			for j:=0; j<n; j++{
				rows[j] = 1
			}
		}else{
			rows[0] = 1
		}
		dp[i] = rows
	}

	for i:=1; i<m; i++{
		for j:=1; j<n; j++{
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
// @lc code=end
