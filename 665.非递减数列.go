/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 *
 * https://leetcode-cn.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (22.18%)
 * Likes:    298
 * Dislikes: 0
 * Total Accepted:    21.9K
 * Total Submissions: 97.9K
 * Testcase Example:  '[4,2,3]'
 *
 * 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
 *
 * 我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [4,2,3]
 * 输出: true
 * 解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,2,1]
 * 输出: false
 * 解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
 *
 *
 *
 *
 * 说明：
 *
 *
 * 1 <= n <= 10 ^ 4
 * - 10 ^ 5 <= nums[i] <= 10 ^ 5
 *
 *
 */

// @lc code=start
// 前一个元素总是比后一个元素小
// 最多改变一个元素
func checkPossibility(nums []int) bool {
	count := 0
	for i:=1; i<len(nums); i++{
		if nums[i] < nums[i-1]{
			count+=1
		}
		if i+1<len(nums) && i-2>=0{  // 防止越界
			if nums[i+1] < nums[i-1] && nums[i-2] > nums[i]{  // 当前元素的后一个元素比前一个元素小  && 当前元素的 小于前两个元素
				return false
			}
		}

		if count > 1{
			return false
		}
	}
	return true

}
// @lc code=end
