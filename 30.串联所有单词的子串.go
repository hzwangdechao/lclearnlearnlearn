/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (32.26%)
 * Likes:    364
 * Dislikes: 0
 * Total Accepted:    47.4K
 * Total Submissions: 146.4K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 *
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 *
 *
 * 示例 2：
 *
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 *
 *
 */

// @lc code=start
// ***** 抓住题意： words中每个单词的长度相等
func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return []int{}
	}
	wordMap := make(map[string]int) // 统计words中每个单词的词频
	for _, word := range words {
		wordMap[word] += 1
	}

	wLen := len(words[0])     // 单个单词的长度
	sLen := len(s)            // 整个字符串的长度
	tLen := wLen * len(words) // 可能满足条件的字符串的长度
	var res []int
	for start := 0; start < wLen; start++ {
		left, right := start, start    // 滑动窗口的左口和右口
		match := 0                     // 匹配成功的次数, 如果match==len(words) && right-left+1 == tLen 则向res添加left
		window := make(map[string]int) // 统计字符串中的词频
		for right+wLen <= sLen {
			rightWord := s[right : right+wLen]
			right += wLen
			// 如果word在wordMap中出现的话
			if wordMap[rightWord] > 0 {
				window[rightWord]++
				if window[rightWord] == wordMap[rightWord] {
					match++
				}
			}
			// 如果left ~ right 的长度满足要求的话
			if right-left == tLen {
				// 如果匹配成功的次数等于words的长度的话， 向res中添加left
				if match == len(wordMap) {
					res = append(res, left)
				}
				// 移动左侧窗口
				leftWord := s[left : left+wLen]
				left += wLen

				// 如果左侧的单词在wordMap中有出现的话
				if wordMap[leftWord] > 0 {
					if wordMap[leftWord] == window[leftWord] {
						// 减少匹配成功的次数
						match--
					}
					window[leftWord]--
				}
			}
		}
	}
	return res
}

// @lc code=end
