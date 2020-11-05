/*
 * @lc app=leetcode.cn id=214 lang=golang
 *
 * [214] 最短回文串
 *
 * https://leetcode-cn.com/problems/shortest-palindrome/description/
 *
 * algorithms
 * Hard (32.34%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    14.6K
 * Total Submissions: 42.8K
 * Testcase Example:  '"aacecaaa"'
 *
 * 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
 *
 * 示例 1:
 *
 * 输入: "aacecaaa"
 * 输出: "aaacecaaa"
 *
 *
 * 示例 2:
 *
 * 输入: "abcd"
 * 输出: "dcbabcd"
 *
 */

// @lc code=start

func shortestPalindrome(s string) string {
	if s == ""{
		return s
	}
	ps := s + "#" + reverseStr(s)

	match := kmp(ps)

	return reverseStr(s[match:]) + s
}

func kmp(s string) int {
	n := len(s)
	next := make([]int, n+1)
	next[0] = -1
	next[1] = 0

	left, right := 0, 2

	for right <= n{
		if s[left] == s[right-1]{
			left ++
			next[right] = left
			right ++
		}else if next[left] == -1{
			left = 0
			right ++
		}else{
			left = next[left]
		}
	}
	fmt.Println(next)
	return next[n]
}



// 反转字符串
func reverseStr(s string) string {
	var bs = []byte(s)
	for i, j:=0,len(s)-1; i<j; i, j=i+1,j-1{
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
// @lc code=end
