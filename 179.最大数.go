/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (37.09%)
 * Likes:    392
 * Dislikes: 0
 * Total Accepted:    42.9K
 * Total Submissions: 115.6K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
 *
 * 示例 1:
 *
 * 输入: [10,2]
 * 输出: 210
 *
 * 示例 2:
 *
 * 输入: [3,30,34,5,9]
 * 输出: 9534330
 *
 * 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 */

// @lc code=start
func largestNumber(nums []int) string {
	var strs []string
	var res string
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	res = strings.Join(strs, "")

	if res[0] == '0'{
		return "0"
	}
	return res
}

// @lc code=end
