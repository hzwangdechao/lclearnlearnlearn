/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 *
 * https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (47.00%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    25.6K
 * Total Submissions: 51.7K
 * Testcase Example:  '5\n7'
 *
 * 给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
 *
 * 示例 1: 
 *
 * 输入: [5,7]
 * 输出: 4
 *
 * 示例 2:
 *
 * 输入: [0,1]
 * 输出: 0
 *
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	for m < n{
		n &= n-1
	}
	return n
	// shift := 0
	// for m < n{
	// 	m, n = m>>1, n>>1
	// 	shift ++
	// }
	// return m << shift
}
// @lc code=end
