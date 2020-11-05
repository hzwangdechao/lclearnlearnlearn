/*
 * @lc app=leetcode.cn id=753 lang=golang
 *
 * [753] 破解保险箱
 *
 * https://leetcode-cn.com/problems/cracking-the-safe/description/
 *
 * algorithms
 * Hard (52.66%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 2.6K
 * Testcase Example:  '1\n1'
 *
 * 有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位是 k 位序列 0, 1, ..., k-1 中的一个 。
 *
 * 你可以随意输入密码，保险箱会自动记住最后 n 位输入，如果匹配，则能够打开保险箱。
 *
 * 举个例子，假设密码是 "345"，你可以输入 "012345" 来打开它，只是你输入了 6 个字符.
 *
 * 请返回一个能打开保险箱的最短字符串。
 *
 *
 *
 * 示例1:
 *
 * 输入: n = 1, k = 2
 * 输出: "01"
 * 说明: "10"也可以打开保险箱。
 *
 *
 *
 *
 * 示例2:
 *
 * 输入: n = 2, k = 2
 * 输出: "00110"
 * 说明: "01100", "10011", "11001" 也能打开保险箱。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n 的范围是 [1, 4]。
 * k 的范围是 [1, 10]。
 * k^n 最大可能为 4096。
 *
 *
 *
 *
 */

// @lc code=start
func crackSafe(n int, k int) string {
	var (
		m = map[string]bool{}
		res []string
	)

	var dfs func(node string)
	dfs = func (node string)  {
		for i:=0; i<k; i++{
			nei := node + strconv.Itoa(i)
			if _, ok := m[nei]; !ok{
				m[nei] = true
				dfs(nei[1:])

				res = append(res, strconv.Itoa(i))
			}
		}
	}

	y := ""
	for i:=0; i<n-1; i++{
		y += "0"
	}

	dfs(y)

	return strings.Join(res, "") + y

}
// @lc code=end
