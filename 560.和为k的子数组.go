/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.81%)
 * Likes:    581
 * Dislikes: 0
 * Total Accepted:    69.2K
 * Total Submissions: 153.6K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 *
 * 示例 1 :
 *
 *
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 *
 *
 * 说明 :
 *
 *
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 *
 *
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	return prefixSum(nums, k)
}


// 前缀和的方法
func prefixSum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)  // 键为和， 值为出现的次数
	pre := 0
	m[0] = 1
	for _, num := range nums{
		pre += num // 计算前缀和
		if v, ok := m[pre-k]; ok{
			res += v
		}
		m[pre] += 1
	}
	return res
}

// 枚举法
func enumerate(nums []int, k int) int {
	res := 0

	for i:=0; i<len(nums); i++{
		sum := 0
		for j:=i; j>=0; j--{
			sum += nums[j]
			if sum == k{
				res += 1
			}
		}
	}
	return res
}
// @lc code=end
