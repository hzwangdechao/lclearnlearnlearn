/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 *
 * https://leetcode-cn.com/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (41.99%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    9.5K
 * Total Submissions: 22.5K
 * Testcase Example:  '2736'
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 *
 * 示例 1 :
 *
 *
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 *
 *
 * 注意:
 *
 *
 * 给定数字的范围是 [0, 10^8]
 *
 *
 */

// @lc code=start
func maximumSwap(num int) int {
	// 首先将数字转成字符串
	str := []byte(strconv.Itoa(num))
	// 存储字符串中每个元素最后出现的位置
	numSlice := [10]int{}
	for i := 0; i < len(str); i++ {
		numSlice[str[i]-'0'] = i
	}
	// 从左向右开始遍历str
	for i := 0; i < len(str); i++ {
		cur := int(str[i] - '0')
		for d := 9; d > cur; d-- {
			if numSlice[d] > i {
				str[i], str[numSlice[d]] = str[numSlice[d]], str[i]
				res, _ := strconv.Atoi(string(str))
				return res
			}
		}
	}
	return num

}

// @lc code=end
