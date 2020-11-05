/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode-cn.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (56.38%)
 * Likes:    290
 * Dislikes: 0
 * Total Accepted:    41.6K
 * Total Submissions: 71.9K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 *
 * 示例 1:
 *
 * 输入: 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 * 输入: 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 * 说明: 你可以假设 n 不小于 2 且不大于 58。
 *
 */

 /*
 2   1+1  1
 3   1+2  2
 4   2+2  4
 5   2+3  6
 6   2+4  8
 7   3+4  12
 8   4+4  16
 9   3+3+3  27
 10  3+3+4  36
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i:=2; i<=n; i++{
		// 当前最大值
		curMax := 0

		// 从1开始遍历， 寻找最大的乘积
		for j:=1; j<i; j++{
			x := j * (i - j)
			y := j * dp[i - j]
			curMax = max(curMax, max(x, y))
		}
		dp[i] = curMax
	}
	fmt.Println(dp)
	return dp[n]
}

func max(x , y int)int  {
	if x > y{
		return x
	}
	return y
}

// @lc code=end
