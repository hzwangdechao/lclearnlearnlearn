/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (43.47%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    111.5K
 * Total Submissions: 253.1K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */



// @lc code=start
func isPalindrome(s string) bool {
	left := 0
	right := len(s) -1
	s = strings.ToLower(s)  // 将字符串全部转换成小写
	for left < right{
		// 如果遇见不合法的字符， 则将指针向前或向后移动一位， 并跳过本次循环
		if !isWord(s[left]){
			left ++
			continue
		}
		if !isWord(s[right]){
			right --
			continue
		}

		// 如果前后字符不相等的话， 直接返回false
		if s[left] != s[right]{
			return false
		}

		// 相等的情况下， 将指针向前和向后移动一位, 进行下一位的判断
		left ++
		right --

	}
	// 全部遍历完成的情况没有出现不相等的情况， 证明是一个回文字符串
	return true
}

func isWord(s byte) bool {
	if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}


// @lc code=end
