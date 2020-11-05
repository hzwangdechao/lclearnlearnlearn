/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 *
 * https://leetcode-cn.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Hard (40.27%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    18.5K
 * Total Submissions: 45.5K
 * Testcase Example:  '"bcabc"'
 *
 *
 * 给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
 *
 *
 *
 * 示例 1:
 *
 * 输入: "bcabc"
 * 输出: "abc"
 *
 *
 * 示例 2:
 *
 * 输入: "cbacdcbc"
 * 输出: "acdb"
 *
 *
 *
 * 注意：该题与 1081
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters
 * 相同
 *
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	m := make(map[byte]int)  // 存放每个字符的数量
	for i:=0; i<len(s); i++{
		m[s[i]] += 1
	}

	stack := make([]byte, 0)  // 构建一个栈来存放元素
	inStack := [26]bool{} // 判断字符是否出现在栈中
	// 遍历字符串s
	for i:=0; i<len(s); i++{
		m[s[i]] -= 1  // 剩余数量减1
		if inStack[s[i]-'a']{
			// 出现在栈中的情况
			continue
		}
		// 没有在栈中出现的情况

		// 如果栈不为空的话， 与栈中的元素进行比较， 如果比之前的元素小， 就将之前的元素去除. 如果栈顶元素在后面的字符串中没有出现， 就不进行去除
		for len(stack) > 0 && stack[len(stack)-1] > s[i] && m[stack[len(stack)-1]]>0{
			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		inStack[s[i]-'a'] = true
	}

	var res string

	for _, b := range stack{
		res += string(b)
	}
	return res
}
// @lc code=end
