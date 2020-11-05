/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (49.09%)
 * Likes:    341
 * Dislikes: 0
 * Total Accepted:    51K
 * Total Submissions: 101.8K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 * 说明：
 *
 *
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 *
 *
 * 示例 1:
 *
 * 输入: n = 3, k = 3
 * 输出: "213"
 *
 *
 * 示例 2:
 *
 * 输入: n = 4, k = 9
 * 输出: "2314"
 *
 *
 */

// @lc code=start
func getPermutation(n int, k int) string {
	factorial := make([]int, n+1) // 计算n个数的阶乘
	factorial[0] = 1
	tokens := make([]int, n)  // 存放1~n

	for i:=1; i<n+1; i++{
		factorial[i] = factorial[i-1]*i
		tokens[i-1] = i
	}

	fmt.Println(factorial)
	fmt.Println(tokens)

	k --

	var b strings.Builder
	for n > 0{
		n --
		a := k/factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}
// @lc code=end
