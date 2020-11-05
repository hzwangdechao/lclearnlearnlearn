/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (44.74%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    81.2K
 * Total Submissions: 179K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 *
 *
 * 示例：
 *
 * s = "leetcode"
 * 返回 0
 *
 * s = "loveleetcode"
 * 返回 2
 *
 *
 *
 *
 * 提示：你可以假定该字符串只包含小写字母。
 *
 */

// @lc code=start
func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, b := range s{
		if _, ok := m[b]; !ok{
			m[b] = 1
		}else{
			m[b] ++
		}
	}
	for i, b := range s{
		if v, ok := m[b]; ok{
			if v ==1 {
				return i
			}
		}
	}

	return -1
}
// @lc code=end
