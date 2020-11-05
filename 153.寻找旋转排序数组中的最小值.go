/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (51.14%)
 * Likes:    222
 * Dislikes: 0
 * Total Accepted:    63.6K
 * Total Submissions: 123.9K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */

// @lc code=start
func findMin(nums []int) int {

	return binary_search(nums)
	if len(nums) == 1{
		return nums[0]
	}
	for i:=1; i<len(nums); i++{
		if nums[i] < nums[i-1]{
			return nums[i]
		}
	}
	return  nums[0]

}

// 利用二分查找的方式
func binary_search(nums []int) int   {
	if len(nums) == 1{
		return nums[0]
	}
	left := 0
	right := len(nums)-1
	if nums[right] > nums[0]{
		return nums[0]
	}
	for left <= right{
		mid := (right-left)/2+left
		if nums[mid-1] > nums[mid]{
			return nums[mid]
		}
		if nums[mid+1] < nums[mid]{
			return nums[mid+1]
		}


		// 如果中间元素比第一个元素大， 则继续查找中间元素的右侧
		if nums[mid] > nums[left]{
			left = mid + 1
		}else if nums[mid] < nums[left]{
			// 如果中间元素比第一个元素小， 则继续查找中间元素的左侧
			right = mid -1
		}
	}
	return nums[0]

}
// @lc code=end
