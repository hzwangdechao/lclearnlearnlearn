/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode-cn.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (54.18%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    18K
 * Total Submissions: 33.1K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x
 * 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 注意: 输入数组的长度不会超过 10000。
 *
 */
//  2 -1   2   2   -1   -1
// [1, 2, 1][1, 2, 1]

// @lc code=start

// 将nums的长度翻倍， 然后利用496栈的相同方法
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)
	var (
		res   = make([]int, len(nums))
		stack []int
	)
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]

		// 删除掉栈中小于等于num的元素
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}

		// 如果栈的长度为0， 则向res中添加 -1
		if len(stack) == 0 {
			res[i] = -1
		} else {
			// 否则栈顶的元素就是第一个比num大的元素
			res[i] = stack[len(stack)-1]
		}
		// 最后将元素添加到栈中
		stack = append(stack, num)
	}

	return res[:n]
}

// 构造多次循环数组的方法
func nextGreaterElements1(nums []int) []int {
	var res []int

	for i, _ := range nums {
		maxNum := RecycleArray(nums, i)
		res = append(res, maxNum)
	}

	return res
}

// 构造循环数组， 找出第一个比nums[k]大的元素
func RecycleArray(nums []int, k int) int {
	var (
		tmp = nums[k]
		s   = k - 1
	)
	// 循环遍历len(nums)次即可
	for i := 0; i < len(nums); i++ {
		s = (s + 1 + len(nums)) % len(nums)
		num := nums[s]
		if num > tmp { // 找到了第一个比tmp的大元素直接返回
			return num
		}
	}
	return -1
}

// @lc code=end
