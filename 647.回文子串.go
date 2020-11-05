/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (62.45%)
 * Likes:    367
 * Dislikes: 0
 * Total Accepted:    59.4K
 * Total Submissions: 91.9K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 * 输入："abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 * 输入："aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 输入的字符串长度不会超过 1000 。
 *
 *
 */

// @lc code=start

// 与lc第五题计算最长回文子串相似，只是不用比较是否是最长回文子串了， 只要是回文子串就将结果+1
func countSubstrings(s string) int {
	length := len(s)
	if length == 0{
		return 0
	}
	dp := make([][]int, length)

	for i:=0; i<length; i++{
		dp[i] = make([]int, length)
	}

	res := 0
	for step:=0; step < length; step++{
		for start:=0; start+step<length; start++{
			end := start + step
			if step == 0{
				dp[start][end] = 1
			}else if step == 1{
				if s[start] == s[end]{
					dp[start][end] = 1
				}
			}else{
				if s[start] == s[end] && dp[start+1][end-1] == 1{
					dp[start][end] = 1
				}
			}

			if dp[start][end] == 1{
				res += 1
			}

		}
	}

	return res
}
// @lc code=end
