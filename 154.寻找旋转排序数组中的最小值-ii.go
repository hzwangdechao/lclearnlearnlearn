/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (48.74%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    30K
 * Total Submissions: 60.4K
 * Testcase Example:  '[1,3,5]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 注意数组中可能存在重复的元素。
 *
 * 示例 1：
 *
 * 输入: [1,3,5]
 * 输出: 1
 *
 * 示例 2：
 *
 * 输入: [2,2,2,0,1]
 * 输出: 0
 *
 * 说明：
 *
 *
 * 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
 * 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	return iteration(nums)

}

func recursive(nums []int)int  {
	if len(nums) == 1{
		return nums[0]
	}

	left := 0
	right := len(nums)-1


	for left < right{
		mid := (right -left)/2 + left

		if nums[mid] < nums[right]{
			// 中间元素比最后一个元素小的时候， 再次比较中间元素的左侧元素
			right = mid
		}else if nums[mid] > nums[right]{
			// 中间元素比最后一个元素大的时候， 再次比较中间元素的右侧元素
			left = mid + 1
		}else{
			right --
		}

	}
	return nums[left]

}

func iteration(nums []int)int  {
	if len(nums) == 1{
		return nums[0]
	}
	for i:=1; i<len(nums); i++{
		if nums[i] < nums[i-1]{
			return nums[i]
		}
	}
	return nums[0]

}
// @lc code=end
