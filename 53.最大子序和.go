/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (51.00%)
 * Likes:    1966
 * Dislikes: 0
 * Total Accepted:    234K
 * Total Submissions: 458.6K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

// @lc code=start
// func maxSubArray(nums []int) int {
// 	// 动态规划解法， 前一个元素大于0， 则将其加到当前元素中

// 	if len(nums) == 0{
// 		return 0
// 	}

// 	// 假设第一个元素是最大和
// 	max := nums[0]

// 	for i:=1; i<len(nums); i++{
// 		// 如果当前元素和前一个元素的和比当前元素大， 则将前一个元素加到当前元素中
// 		if nums[i] + nums[i-1] > nums[i]{
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max{
// 			max = nums[i]
// 		}
// 	}

// 	return max

// }

// func maxSubArray(nums []int) int {
// 	// 贪心算法解法， 若当前的指针所指元素之前的和小于0， 则丢弃当前元素之前的数列
// 	res := math.MinInt32
// 	sum := 0

// 	for i:=0; i<len(nums); i++{
// 		sum += nums[i]

// 		// res = max(sum, res)
// 		if sum > res{
// 			res = sum
// 		}
// 		if sum < 0{
// 			sum = 0
// 		}
// 	}
// 	return res
// }

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0

	for _, v := range nums{
		if sum > 0{
			sum += v
		}else{
			sum = v
		}
		res = max(res, sum)
	}
	return res
}

// sum = 1 res= 1
// res = -2 res = 1


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end
