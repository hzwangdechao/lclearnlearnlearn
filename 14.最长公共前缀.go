/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (37.09%)
 * Likes:    981
 * Dislikes: 0
 * Total Accepted:    239.3K
 * Total Submissions: 645.3K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start

func longestCommonPrefix(strs []string) string {
	// 先找到strs中长度最短的字符串
	if len(strs) == 0{
		return ""
	}
	short := strs[0]
	for _, str := range strs{
		if len(str) < len(short){
			short = str
		}
	}

	for i, s := range short{
		for _, str := range strs{
			if byte(s) != str[i]{
				return short[:i]
			}
		}
	}
	return short

}

// @lc code=end
