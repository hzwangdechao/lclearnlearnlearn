/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (57.43%)
 * Likes:    1017
 * Dislikes: 0
 * Total Accepted:    311.8K
 * Total Submissions: 542.9K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */

// @lc code=start
//func isPalindrome(x int) bool {
//	original := x
//	// 整数反转的方法， 如果反转后的方法与原先的整数相等，证明是回文数
//	if x < 0{
//		// 负数不是回文数
//		return false
//	}
//	y := 0
//	for x != 0{
//		y = y*10 + x%10
//		x /=10
//	}
//
//	if y == original {
//		return true
//	}
//	return false
//}
import "strconv"
func isPalindrome(x int) bool {
	// 利用字符串的方法
	if x < 0{
		return false
	}
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < len(str); i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true


}
// @lc code=end
