/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (40.12%)
 * Likes:    715
 * Dislikes: 0
 * Total Accepted:    87.4K
 * Total Submissions: 217.7K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	preMax, preMin, res := nums[0], nums[0], nums[0]

	for i:=1; i<len(nums); i++{
		curMax := max(nums[i]*preMax, max(nums[i]*preMin, nums[i]))
		curMin := min(nums[i]*preMin, min(nums[i]*preMax, nums[i]))
		res = max(curMax, res)
		preMax, preMin = curMax, curMin
	}

	return res

}
func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y{
		return x
	}
	return y
}
// @lc code=end
