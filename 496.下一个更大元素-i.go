/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 *
 * https://leetcode-cn.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (65.79%)
 * Likes:    288
 * Dislikes: 0
 * Total Accepted:    48.5K
 * Total Submissions: 73.7K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * 给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2
 * 中的下一个比其大的值。
 *
 * nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
 *
 *
 *
 * 示例 1:
 *                                3  4  -1  -1
 * 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
 * 输出: [-1,3,-1]
 * 解释:
 * ⁠   对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
 * ⁠   对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
 * ⁠   对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
 *
 * 示例 2:
 *
 * 输入: nums1 = [2,4], nums2 = [1,2,3,4].
 * 输出: [3,-1]
 * 解释:
 * 对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
 * ⁠   对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1和nums2中所有元素是唯一的。
 * nums1和nums2 的数组大小都不超过1000。
 *
 *
 */

// @lc code=start   [4 3]
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 利用栈 从后向前遍历
	var (
		stack []int
		res   = make([]int, 0)
		m     = make(map[int]int) // 存放nums2中每个元素的下一个更大元素
	)

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]

		// 如果stack的长度大于0 并且 num大于 stack的最后一个元素， 将最后一个元素出栈
		for len(stack) > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
		}

		// 如果stack为空的话，则当前num的下一个更大元素为-1
		if len(stack) == 0 {
			m[num] = -1
		} else { // 否则当前num的下一个更大元素为stack的最后一个元素
			m[num] = stack[len(stack)-1]

		}
		// 最终将nums[i]添加到stack中
		stack = append(stack, num)
	}

	for _, num := range nums1 {
		x := m[num]
		res = append(res, x)
	}

	return res
}

// @lc code=end
