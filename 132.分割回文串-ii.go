/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning-ii/description/
 *
 * algorithms
 * Hard (43.97%)
 * Likes:    198
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 36.9K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回符合要求的最少分割次数。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出: 1
 * 解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
 *
 *
 */

// @lc code=start
func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1) // dp[i]存放0~i最少的切割次数
	dp[0] = -1
	dp[1] = 0 // s的第一个字符的最小切割次数为0

	for i := 2; i <= n; i++ { // 从s的第二个字符开始计算
		dp[i] = i - 1 // 第i个字符的最多需要切割 i-1次
		for j := 0; j < i; j++ {
			// 例如 aab字符串 aa是回文字符串 因此dp[2]其实就是dp[0]+1=0次
			if judgePalindrome(s[j:i]) { // 如果s[j:i]是回文字符串的话， 那么dp[i]的做少切割次数为dp[j]+1
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
