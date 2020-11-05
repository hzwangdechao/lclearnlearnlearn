/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 *
 * https://leetcode-cn.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (30.08%)
 * Likes:    331
 * Dislikes: 0
 * Total Accepted:    29.1K
 * Total Submissions: 96.5K
 * Testcase Example:  '"1432219"\n3'
 *
 * 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
 *
 * 注意:
 *
 *
 * num 的长度小于 10002 且 ≥ k。
 * num 不会包含任何前导零。
 *
 *
 * 示例 1 :
 *
 *
 * 输入: num = "1432219", k = 3
 * 输出: "1219"
 * 解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: num = "10200", k = 1
 * 输出: "200"
 * 解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
 *
 *
 * 示例 3 :
 *
 *
 * 输入: num = "10", k = 2
 * 输出: "0"
 * 解释: 从原数字移除所有的数字，剩余为空就是0。
 *
 *
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	remain := len(num) - k
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			// 如果栈中最后一个元素比当前元素大的话， 将其从栈中移除掉
			stack = stack[:len(stack)-1]
			k -= 1
		}
		// 将当前元素添加到栈中
		stack = append(stack, num[i])
	}
	stack = stack[:remain]

	i := 0
	for i < len(stack) {  // 去掉字符串开头的0
		if stack[i] == '0' && i < len(stack)-1 {
			i++
		} else {
			break
		}
	}
	if string(stack[i:]) == "" {
		return "0"
	}
	return string(stack[i:])
}

// @lc code=end
