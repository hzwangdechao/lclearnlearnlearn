/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (45.82%)
 * Likes:    351
 * Dislikes: 0
 * Total Accepted:    36.4K
 * Total Submissions: 79K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 *
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 *
 * 说明：
 *
 *
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * s: "cbaebabacd" p: "abc"
 *
 * 输出:
 * [0, 6]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s: "abab" p: "ab"
 *
 * 输出:
 * [0, 1, 2]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 *
 *
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	var res []int
	sMap := make(map[byte]int)
	pMap := make(map[byte]int) // 记录p中每个单词出现的次数

	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}
	pLen := len(p)

	check := func() bool {
		for k, v := range pMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]]++

		for check() {
			if r-l+1 == pLen {
				res = append(res, l)
			}
			sMap[s[l]]--
			l++
		}
	}

	return res
}

// func findAnagrams(s string, p string) []int {
// 	var res []int
// 	origin := changShape([]byte(p))
// 	step := len(p)
// 	for i:=0; i+step <= len(s); i++{
// 		cur := changShape([]byte(s[i:i+step]))
// 		if cur == origin{
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

// func changShape(s []byte) [26]int {
// 	cur := [26]int{}
// 	for _, b := range s{
// 		cur[b-'a'] += 1
// 	}
// 	return cur
// }

// @lc code=end
