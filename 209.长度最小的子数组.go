/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (44.39%)
 * Likes:    389
 * Dislikes: 0
 * Total Accepted:    77.6K
 * Total Submissions: 174.8K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续
 * 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 *
 *
 *
 * 示例：
 *
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 *
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	if s == 0 {
		return 0
	}
	res := math.MaxInt32
	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= s { // 缩小左边界
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
