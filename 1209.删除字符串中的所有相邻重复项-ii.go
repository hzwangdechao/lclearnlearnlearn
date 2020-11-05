/*
 * @lc app=leetcode.cn id=1209 lang=golang
 *
 * [1209] 删除字符串中的所有相邻重复项 II
 *
 * https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/description/
 *
 * algorithms
 * Medium (52.99%)
 * Likes:    50
 * Dislikes: 0
 * Total Accepted:    6.2K
 * Total Submissions: 11.6K
 * Testcase Example:  '"abcd"\n2'
 *
 * 给你一个字符串 s，「k 倍重复项删除操作」将会从 s 中选择 k 个相邻且相等的字母，并删除它们，使被删去的字符串的左侧和右侧连在一起。
 *
 * 你需要对 s 重复进行无限次这样的删除操作，直到无法继续为止。
 *
 * 在执行完所有删除操作后，返回最终得到的字符串。
 *
 * 本题答案保证唯一。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abcd", k = 2
 * 输出："abcd"
 * 解释：没有要删除的内容。
 *
 * 示例 2：
 *
 * 输入：s = "deeedbbcccbdaa", k = 3
 * 输出："aa"
 * 解释：
 * 先删除 "eee" 和 "ccc"，得到 "ddbbbdaa"
 * 再删除 "bbb"，得到 "dddaa"
 * 最后删除 "ddd"，得到 "aa"
 *
 * 示例 3：
 *
 * 输入：s = "pbbcggttciiippooaais", k = 2
 * 输出："ps"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * 2 <= k <= 10^4
 * s 中只含有小写英文字母。
 *
 *
 */

// @lc code=start

func removeDuplicates(s string, k int) string {
	stack := []byte{s[0]}

	for i:=1; i<len(s); i++{
		// fmt.Println(string(stack))
		// 如果链表的数量小于K-1的话直接将当前字符添加到栈中
		if len(stack) < k-1{
			stack = append(stack, s[i])
		}else{
			// 遍历栈中的倒数K个元素
			del := true
			for j:=len(stack)-1; j>len(stack)-k; j--{
				if stack[j] != s[i]{
					// 如果栈中的倒数k个元素与s[i]不相等
					del = false
					break
				}
			}
			if del{
				stack = stack[:len(stack)-k+1]
			}else{
				stack = append(stack, s[i])
			}

		}
	}
	return string(stack)
}
// @lc code=end
