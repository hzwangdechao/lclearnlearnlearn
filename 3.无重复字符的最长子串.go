/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (34.33%)
 * Likes:    3601
 * Dislikes: 0
 * Total Accepted:    470.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	res := 0
	sMap := make(map[byte]int) // 统计每个字符串出现的次数

	for l, r := 0, 0; r < len(s); {
		// 遍历右侧窗口， 直到出现重复元素位置
		for r < len(s) && sMap[s[r]] == 0 {
			sMap[s[r]] ++
			r++
		}

		curLen := r - l // 当前不重复字符串的长度

		res = max(curLen, res)

		// 左侧窗口向右移动
		sMap[s[l]] --
		l++

	}

	return res
}

// func lengthOfLongestSubstring(s string) int {
// 	ans := 0  // 最终返回的最长字符串的长度
// 	hash := make(map[byte]int, 0)  // 保存每个字符数量的hash字典
// 	rightPoint := -1  // 滑动窗口的右指针

// 	for i:=0; i<len(s); i++{
// 		if i!=0{
// 			// 删除左指针的前一个字符
// 			delete(hash, s[i-1])
// 		}

// 		for rightPoint+1<len(s) && hash[s[rightPoint+1]] == 0{
// 			// 知道出现重复字符为止  停止滑动
// 			hash[s[rightPoint+1]] ++
// 			rightPoint ++
// 		}
// 		// 当前无重复字符串的长度
// 		cur := rightPoint - i + 1

// 		// 比较大小
// 		ans = max(ans, cur)
// 	}
// 	return ans

// }

func max(x, y int)int  {
	if x > y{
		return x
	}
	return y

}

// @lc code=end
