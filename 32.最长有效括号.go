/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (33.50%)
 * Likes:    948
 * Dislikes: 0
 * Total Accepted:    99K
 * Total Submissions: 293.7K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 *
 * 示例 1:
 *
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 *
 *
 * 示例 2:
 *
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	return method3(s)
	stack := make([]int, 0) // 存储左括号的索引
	stack = append(stack, -1)
	res := 0

	for i:=0; i<len(s); i++{
		// fmt.Println(stack)
		if s[i] == '('{
			// 如果是左括号的话， 就将括号的索引入栈
			stack = append(stack, i)
		}else{
			stack = stack[:len(stack)-1]
			// 从栈中取出元素

			if len(stack) == 0{
				stack = append(stack, i)
			}else{
				res = max(res, i-stack[len(stack)-1])
			}

		}
	}
	return res
}

func method3(s string) int {
	// 从左向右计算左括号和右括号的长度
	res := 0
	left, right := 0, 0
	for i:=0; i<len(s); i++{
		if s[i] == '('{
			left ++
		}else{
			right ++
		}

		if left == right{
			res = max(res, 2*right)
		}else if right > left{
			left, right =0, 0
		}
	}

	left, right = 0, 0

	for i:=len(s)-1; i>=0; i--{
		if s[i] == ')'{
			right ++
		}else{
			left ++
		}
		if left == right{
			res = max(res, 2*left)
		}else if left > right{
			left, right = 0, 0
		}
	}

	return res

}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}
// @lc code=end
