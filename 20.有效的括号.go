/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (41.48%)
 * Likes:    1548
 * Dislikes: 0
 * Total Accepted:    267.8K
 * Total Submissions: 645.5K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

 // 栈 ： 先进后出
// @lc code=start
func isValid(s string) bool {
	// 注意题目中有三种括号组合，先设置一个对应关系 Map
	// 左括号类型入栈，右括号类型匹配栈顶元素，如果不是对应关系，则肯定 false，匹配上之后移除栈顶元素
	// 最后看栈是否为空，空则是有效括号，非空则是无效的
	// 同时注意题目中空字符串是返回 true

	if s == "" {
		return true
	}

	var stack []uint8

	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i:=0; i < len(s); i++{
		// stack 只保留左括号
		if string(s[i]) == "[" || string(s[i]) == "{" || string(s[i]) == "("{
			stack = append(stack, s[i])
		}else{
			if len(stack) == 0{
				// 说明出现了非法字符
				return false
			}
			if m[s[i]] != stack[len(stack)-1]{
				// 这个右括号对应的左括号不是 stack 里面的最后一个元素
				return false
			}
			// 将stack最后一个元素出栈
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}


// @lc code=end
