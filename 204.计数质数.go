/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (33.47%)
 * Likes:    367
 * Dislikes: 0
 * Total Accepted:    61.8K
 * Total Submissions: 181.6K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 2{
		return 0
	}
	primes := make([]bool, n)

	count := 0
	for i:=2; i < n ; i++{
		if primes[i]{
			continue
		}
		count ++
		for j:=2 * i; j<n; j+=i{
			primes[j] = true
		}
	}

	return count

}
// @lc code=end
