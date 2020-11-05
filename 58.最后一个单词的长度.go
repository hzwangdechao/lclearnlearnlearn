/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (33.04%)
 * Likes:    194
 * Dislikes: 0
 * Total Accepted:    88.1K
 * Total Submissions: 266.3K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
 *
 *
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */

// @lc code=start
// func lengthOfLastWord(s string) int {
// 	strList := strings.Split(s,' ')

// 	for i:=len(strList)-1;i >= 0;i --{
// 		if len(strList[i]) != 0{
// 			return len(strList[i])
// 		}
// 	}
// 	return 0
// }

func lengthOfLastWord(s string) int {
	count := 0
	found := false
	for i:= len(s)-1; i >= 0;i--{
		if !found && s[i] ==' '{
			// 防止最后一个是' ' 的情况
			continue
		}
		if found && s[i] ==' '{
			// 找到' ' 就停止
			break
		}
		count ++
		found = true
	}
	return count
}
// @lc code=end
