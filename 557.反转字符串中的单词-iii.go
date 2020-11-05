/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (71.15%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    66.7K
 * Total Submissions: 93.6K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 * 示例 1:
 *
 *
 * 输入: "Let's take LeetCode contest"
 * 输出: "s'teL ekat edoCteeL tsetnoc" 
 *
 *
 * 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 *
 */

// @lc code=start
func reverseWords(s string) string {
	array := []byte(s)
	start :=0

	for i:=0; i<len(array); i++{
		if array[i] == ' '{
			reverse(array, start, i-1)
			start = i+1
		}
	}
	reverse(array, start, len(array)-1)
	return string(array)
}

func reverse(s []byte, start, end int)  {
	for start <= end{
		s[start], s[end] = s[end], s[start]
		start ++
		end --
	}
}
// @lc code=end
