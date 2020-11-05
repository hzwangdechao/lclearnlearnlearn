/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (36.82%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    22.9K
 * Total Submissions: 61.6K
 * Testcase Example:  '"3+2*2"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
 *
 * 示例 1:
 *
 * 输入: "3+2*2"
 * 输出: 7
 *
 *
 * 示例 2:
 *
 * 输入: " 3/2 "
 * 输出: 1
 *
 * 示例 3:
 *
 * 输入: " 3+5 / 2 "
 * 输出: 5
 *
 *
 * 说明：
 *
 *
 * 你可以假设所给定的表达式都是有效的。
 * 请不要使用内置的库函数 eval。
 *
 *
 */

// @lc code=start
func calculate(s string) int {
	var num, ans int

	sign := '+'
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num = 10*num + int(s[i]-'0')
		}

		if (!isDigit(s[i]) && s[i] != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num = pre * num
				stack = append(stack, num)
			case '/':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num = pre / num
				stack = append(stack, num)
			}

			sign = int32(s[i])
			num = 0
		}
	}

	fmt.Println(stack)
	for _, v := range stack {
		ans += v
	}

	return ans

}

func isDigit(b byte) bool {

	s := []byte{'0','1','2','3','4','5','6','7','8','9'}

	for _, x := range s{
		if b == x{
			return true
		}
	}
	return false
}
// @lc code=end
