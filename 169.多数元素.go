/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (64.24%)
 * Likes:    698
 * Dislikes: 0
 * Total Accepted:    200.2K
 * Total Submissions: 311.6K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */

// @lc code=start
func majorityElement(nums []int) int {
	return method2(nums)
}

func method2(nums []int)int  {
	sort.Ints(nums)
	step := len(nums)/2
	for i:=0; i<=step; i++{
		if nums[i] == nums[i+step]{
			return nums[i]
		}
	}
	return -1

}
func method1(nums []int)int  {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
// @lc code=end
