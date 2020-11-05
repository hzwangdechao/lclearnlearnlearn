/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.33%)
 * Likes:    867
 * Dislikes: 0
 * Total Accepted:    154.7K
 * Total Submissions: 402.5K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end{
		mid := (end-start)/2+start

		// 如果中间元素和target相等则直接返回
		if nums[mid] == target{
			return mid
		}

		if nums[0] <= nums[mid]{
			// 如果开始位置比中间元素小的话
			if nums[0] <= target && nums[mid] > target{
				// 在左边的元素中查找
				end = mid - 1
			}else{
				// 在右边的元素中查找
				start = mid + 1
			}
		}else{
			// 开始位置比中间位置大的情况
			if nums[mid] < target && nums[len(nums)-1] >= target{
				start = mid + 1
			}else{
				end = mid - 1
			}
		}

	}
	return -1
}
// @lc code=end
