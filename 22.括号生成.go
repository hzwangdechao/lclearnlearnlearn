/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.25%)
 * Likes:    1320
 * Dislikes: 0
 * Total Accepted:    181K
 * Total Submissions: 237.3K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	return methodWithBackTrack(n)
}

func methodWithBackTrack(n int) []string {
	res := []string{}

	var backTrack func(S []string, left int, right int)
	backTrack = func (S []string, left int, right int)  {
		if len(S) == 2*n{
			res = append(res, strings.Join(S, ""))
			return
		}

		if left < n{
			S = append(S, "(")
			backTrack(S, left+1, right)
			S = S[:len(S)-1]
		}

		if right < left{
			S = append(S, ")")
			backTrack(S, left, right+1)
			S = S[:len(S)-1]
		}
	}
	backTrack([]string{}, 0, 0)
	return res
}

// @lc code=end
