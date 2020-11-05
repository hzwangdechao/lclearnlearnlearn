/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (74.41%)
 * Likes:    347
 * Dislikes: 0
 * Total Accepted:    77.6K
 * Total Submissions: 103.3K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */
// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	temp := []int{}
	var dfs func(int)

	dfs = func (cur int)  {
		if len(temp) + (n-cur+1) < k{
			// temp的长度加上 [cur, n]的长度小于k, 此时不可能构造出长度为k的temp
			return
		}

		// 记录合法的方案
		if len(temp) == k{
			comb := make([]int, k)
			copy(comb, temp)
			res = append(res, comb)
			return
		}

		// 考虑选择当前的位置
		temp = append(temp, cur)
		dfs(cur+1)
		// 不考虑当前位置
		temp = temp[:len(temp)-1]
		dfs(cur+1)
	}
	dfs(1)

	return res
}
// @lc code=end
