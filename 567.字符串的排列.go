/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (37.54%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    45.5K
 * Total Submissions: 121K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 *
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 *
 * 示例1:
 *
 *
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 *
 *
 *
 *
 * 示例2:
 *
 *
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 *
 *
 *
 *
 * 注意：
 *
 *
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 *
 *
 */

// @lc code=start
func checkInclusion(t string, s string) bool {
	tMap, sMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]]++

		for check() && l <= r {
			if len(t) == r-l+1 {
				return true
			}
			sMap[s[l]]--
			l++
		}
	}
	return false

}

// @lc code=end
