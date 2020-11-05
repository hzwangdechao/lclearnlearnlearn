/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (47.26%)
 * Likes:    625
 * Dislikes: 0
 * Total Accepted:    83.7K
 * Total Submissions: 177K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 *
 * 说明：
 *
 *
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// 创建一个map， 判断字符串是否在wordDict出现过
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for end := 1; end <= n; end++ {
		for start := 0; start < end; start++ {
			if m[s[start:end]] && dp[start] {  // 以start结尾的字符串可以在wordDict出现过
				dp[end] = true
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end
