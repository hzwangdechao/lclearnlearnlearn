/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.39%)
 * Likes:    656
 * Dislikes: 0
 * Total Accepted:    184.8K
 * Total Submissions: 286.9K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */

// @lc code=start

func findKthLargest(nums []int, K int) int{
	idx := quickSort(0, len(nums)-1, len(nums)-K, nums)
	return nums[idx]
	// fmt.Println(nums[idx])
	// fmt.Println(nums)

}

func quickSort(left, right, pos int, nums []int) int {
	partitionIdx := partition(left, right, nums)
	if partitionIdx == pos {
		return partitionIdx
	} else if partitionIdx > pos {
		return quickSort(left, partitionIdx-1, pos, nums)
	} else {
		return quickSort(partitionIdx+1, right, pos, nums)
	}
}

func partition(left, right int, nums []int) int {
	l := left
	temp := nums[left]
	for left < right {
		// 从右向左找到第一个比temp小的元素
		for left < right && nums[right] >= temp {
			right--
		}
		// 从左向右找到第一个比temp大的元素
		for left < right && nums[left] <= temp {
			left++
		}
		if left < right {
			// 交换left和right的位置
			nums[left], nums[right] = nums[right], nums[left]

			// 另一种写法
			//nums[left] = nums[right] ^ nums[left]
			//nums[right] = nums[right] ^ nums[left]
			//nums[left] = nums[right] ^ nums[left]
		}
	}

	nums[l] = nums[left]
	nums[left] = temp

	return left
}

// @lc code=end
