/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 *
 * https://leetcode-cn.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (38.26%)
 * Likes:    247
 * Dislikes: 0
 * Total Accepted:    18.5K
 * Total Submissions: 48K
 * Testcase Example:  '"1 + 1"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。
 *
 * 示例 1:
 *
 * 输入: "1 + 1"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: " 2-1 + 2 "
 * 输出: 3
 *
 * 示例 3:
 *
 * 输入: "(1+(4+5+2)-3)+(6+8)"
 * 输出: 23
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
	var num, res int

	sign := 1  // 控制正负号
	stack := make([]int, 0)

	for i:=0; i<len(s); i++{
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num =0
			for ; i<len(s) && s[i]>='0' && s[i] <='9'; i++{
				num = 10*num + int(s[i]-'0')
			}
			res += sign*num
			fmt.Println(i)
			i--
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			// 入栈
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		case ')':
			// 出栈
			sign = stack[len(stack)-1]
			temp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res = sign*res + temp
		}

		fmt.Println("ads", i)
	}
	return res
}
// @lc code=end
