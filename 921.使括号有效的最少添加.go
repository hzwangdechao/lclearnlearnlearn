/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 *
 * https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/description/
 *
 * algorithms
 * Medium (72.31%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 16.9K
 * Testcase Example:  '"())"'
 *
 * 给定一个由 '(' 和 ')' 括号组成的字符串 S，我们需要添加最少的括号（ '(' 或是 ')'，可以在任何位置），以使得到的括号字符串有效。
 *
 * 从形式上讲，只有满足下面几点之一，括号字符串才是有效的：
 *
 *
 * 它是一个空字符串，或者
 * 它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
 * 它可以被写作 (A)，其中 A 是有效字符串。
 *
 *
 * 给定一个括号字符串，返回为使结果字符串有效而必须添加的最少括号数。
 *
 *
 *
 * 示例 1：
 *
 * 输入："())"
 * 输出：1
 *
 *
 * 示例 2：
 *
 * 输入："((("
 * 输出：3
 *
 *
 * 示例 3：
 *
 * 输入："()"
 * 输出：0
 *
 *
 * 示例 4：
 *
 * 输入："()))(("
 * 输出：4
 *
 *
 *
 * 提示：
 *
 *
 * S.length <= 1000
 * S 只包含 '(' 和 ')' 字符。
 *
 *
 *
 *
 */
// "())"
// 0
// 1
// @lc code=start
func minAddToMakeValid(S string) int {
	left, right := 0, 0
	for i:=0; i<len(S); i++{
		if S[i] == '('{
			// 如果是左括号的话, 将左括号的需求量增加1
			right += 1
		}else{
			// 如果是右括号的话， 将右括号的需求量减少1
			right -= 1

			if right == -1{
				// 右括号的数量太多了, 将左括号的数量加1。并重置右括号的需求量
				left += 1
				right = 0
			}
		}
	}

	return left + right
}
// @lc code=end
