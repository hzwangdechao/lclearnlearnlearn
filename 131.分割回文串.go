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
	res := make([][]string, 0)
	path := make([]string, 0)
	backTrack(path, &res, 0, s)
	return res
}

// 递归回溯算法
func backTrack(path []string, res *[][]string, index int, s string) {
	// 如果索引和s的长度相等的话， 说明找到了符合条件的回文串列表
	if index == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := index; i < len(s); i++ {
		if judgePalindrome(s[index : i+1]) {
			path = append(path, s[index:i+1])
			backTrack(path, res, i+1, s)
			path = path[:len(path)-1]
		}
	}
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
