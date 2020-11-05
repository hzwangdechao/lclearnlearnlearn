/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (47.61%)
 * Likes:    277
 * Dislikes: 0
 * Total Accepted:    28.8K
 * Total Submissions: 58.1K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i:=1; i*2 <= n; i++{
		if n%i == 0{
			match := true
			for j:=i; j<n; j++{
				if s[j] != s[j-i]{
					match = false
					break
				}
			}
			if match {
				return true
			}
		}

	}
	return false
}
// @lc code=end
