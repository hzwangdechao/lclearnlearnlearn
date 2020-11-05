/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (50.10%)
 * Likes:    211
 * Dislikes: 0
 * Total Accepted:    53.5K
 * Total Submissions: 104.1K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	add := 0  // 进位
	ans := ""  // 最后返回的结果

	for i, j:=len(num1)-1, len(num2)-1; i>=0 || j>=0 || add != 0; i, j=i-1, j-1{
		var x, y int
		if i>=0{
			x = int(num1[i]-'0')
		}
		if j>=0{
			y = int(num2[j]-'0')
		}

		result := x+y+add

		ans = strconv.Itoa(result%10) + ans

		add = result/10
	}

	return ans
}
// @lc code=end
