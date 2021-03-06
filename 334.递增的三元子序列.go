/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 *
 * https://leetcode-cn.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (38.89%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    22.5K
 * Total Submissions: 57.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
 *
 * 数学表达式如下:
 *
 * 如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
 * 使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
 *
 * 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5]
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: [5,4,3,2,1]
 * 输出: false
 *
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3{
		return false
	}
	// 分别记录第一小的数和第二小的数
	f, s := math.MaxInt32, math.MaxInt32
	for _, num := range nums{
		if num <= f{
			f = num
		}else if num <= s{
			s = num
		}else{
			return true
		}
	}
	return false

}
// @lc code=end
