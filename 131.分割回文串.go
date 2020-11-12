/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (68.59%)
 * Likes:    344
 * Dislikes: 0
 * Total Accepted:    41.5K
 * Total Submissions: 60.5K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */

// @lc code=start
func partition(s string) [][]string {
	// 回溯方法分割回文串
	var (
		traceBack func(path []string, start int)
		res       [][]string
	)
	traceBack = func(path []string, start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			if judgePalindrome(s[start : i+1]) {
				path = append(path, s[start:i+1])
				traceBack(path, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	traceBack([]string{}, 0)
	return res

}

// 判断字符串是否是回文串
func judgePalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// @lc code=end
