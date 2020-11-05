/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 *
 * https://leetcode-cn.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (55.13%)
 * Likes:    221
 * Dislikes: 0
 * Total Accepted:    58K
 * Total Submissions: 105.1K
 * Testcase Example:  '"abccccdd"'
 *
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *
 * 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 *
 * 注意:
 * 假设字符串的长度不会超过 1010。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "abccccdd"
 *
 * 输出:
 * 7
 *
 * 解释:
 * 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i:=0; i<len(s); i++{
		m[s[i]] += 1
	}

	res := 0
	for _, v := range m{
		res += v / 2 * 2
	}
	fmt.Println(res)
	if res < len(s){
		return res + 1
	}
	return res
}
// @lc code=end
