/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (66.77%)
 * Likes:    630
 * Dislikes: 0
 * Total Accepted:    114.6K
 * Total Submissions: 171.6K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
 *
 *
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	return method2(triangle)
}

// 从下往上计算
func method2(triangle [][]int) int {
	n := len(triangle)
	dp := triangle[n-1] // 最后一行的数据作为base case

	for i := n - 2; i >= 0; i-- { // 从倒数第二行开始计算
		for j := 0; j <= i; j++ { // 从左往右计算
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]

}

// 从上往下计算
func method1(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	// 从第二层开始遍历
	for i := 1; i < n; i++ {
		// 计算开头元素
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		// 计算中间元素
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}

		// 计算结尾元素
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := math.MaxInt32
	for _, num := range dp[n-1] { // 从最后一行中找到最小的元素即为正确答案
		res = min(res, num)
	}
	return res

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
