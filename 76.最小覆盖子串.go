/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (38.63%)
 * Likes:    705
 * Dislikes: 0
 * Total Accepted:    72.8K
 * Total Submissions: 187.6K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。
 *
 *
 *
 * 示例：
 *
 * 输入：S = "ADOBECODEBANC", T = "ABC"
 * 输出："BANC"
 *
 *
 *
 * 提示：
 *
 *
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 */

// @lc code=start
func minWindow(s string, t string) string {
	// sliding window
	sMap, tMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	// fmt.Println(tMap)

	var check func() bool
	check = func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	ansL, ansR := -1, -1
	length := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {

		// 可以省略判断
		// if r < len(s) && tMap[s[r]] > 0 {
		sMap[s[r]]++
		// }

		for check() && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				ansL, ansR = l, l+length
			}

			// 可以省略判断
			// if _, ok := tMap[s[l]]; ok {
			sMap[s[l]]--
			// }

			l++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

// @lc code=end
