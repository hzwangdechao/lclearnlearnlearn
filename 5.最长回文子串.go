/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (31.12%)
 * Likes:    2474
 * Dislikes: 0
 * Total Accepted:    324.5K
 * Total Submissions: 1M
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	length := len(s)
	dp := make([][]int, length)

	for i:=0; i<length; i++{
		dp[i] = make([]int, length)
	}
	res := ""
	for step:=0; step<length; step++{
		// step： 每次移动的步伐
		for start:=0; start+step<length; start++{
			end := start + step
			if step == 0{
				// 步伐为0的时候不需要进行判断
				dp[start][end] = 1
			}else if step == 1{
				// 步伐为1的时候只需要判断开始位置和结束位置是否一样
				if s[start] == s[end]{
					dp[start][end] = 1
				}
			}else{
				// 其他时候不仅需要判断开始位置和结束位置是否一样， 还要判断中间位置是否一样、
				if s[start] == s[end] && dp[start+1][end-1] == 1{
					dp[start][end] = 1
				}
			}
			// 判断此次回文子串是不是比之前的结果还长
			if step + 1 > len(res) && dp[start][end] == 1{
				res = s[start:start+step+1]
			}
		}
	}

	return res

}




// @lc code=end
