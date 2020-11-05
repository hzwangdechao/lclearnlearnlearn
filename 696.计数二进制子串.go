/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 *
 * https://leetcode-cn.com/problems/count-binary-substrings/description/
 *
 * algorithms
 * Easy (53.26%)
 * Likes:    224
 * Dislikes: 0
 * Total Accepted:    27K
 * Total Submissions: 44.4K
 * Testcase Example:  '"00110"'
 *
 * 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
 *
 * 重复出现的子串要计算它们出现的次数。
 *
 * 示例 1 :
 *
 *
 * 输入: "00110011"
 * 输出: 6
 * 解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
 *
 * 请注意，一些重复出现的子串要计算它们出现的次数。
 *
 * 另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: "10101"
 * 输出: 4
 * 解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
 *
 *
 * 注意：
 *
 *
 * s.length 在1到50,000之间。
 * s 只包含“0”或“1”字符。
 *
 *
 */
// 00111011
// 0011
// 01
// 10
// 01

// [2, 3, 1, 2]
// @lc code=start
func countBinarySubstrings(s string) int {
	// 首先计算相同元素连续出现的次数
	counts := []int{}
	l, r := 0, len(s)

	for l < r {
		cur := s[l]  // 当前字符串
		count := 0  // 相同字符出现的次数
		for l<r && s[l] == cur{
			l ++
			count ++
		}

		counts = append(counts, count)
	}
	res := 0
	for i:=1; i<len(counts); i++{
		res += min(counts[i], counts[i-1])
	}

	return res
}
func min(x, y int)int  {
	if x < y{
		return x
	}
	return y

}
// @lc code=end
