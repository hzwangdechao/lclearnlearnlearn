/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 *
 * https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/
 *
 * algorithms
 * Easy (81.92%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    40.4K
 * Total Submissions: 48.8K
 * Testcase Example:  '[8,1,2,2,3]'
 *
 * 给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
 *
 * 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
 *
 * 以数组形式返回答案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [8,1,2,2,3]
 * 输出：[4,0,1,1,3]
 * 解释：
 * 对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
 * 对于 nums[1]=1 不存在比它小的数字。
 * 对于 nums[2]=2 存在一个比它小的数字：（1）。
 * 对于 nums[3]=2 存在一个比它小的数字：（1）。
 * 对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [6,5,4,8]
 * 输出：[2,1,0,3]
 *
 *
 * 示例 3：
 *
 * 输入：nums = [7,7,7,7]
 * 输出：[0,0,0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 500
 * 0 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	numCount := make([]int, 101)
	for _, num := range nums {
		numCount[num] ++
	}

	pre := numCount[0]
	numCount[0] = 0 // numCount[i]=10  比numCount[i]小的数字有10个
	for i := range numCount[1:] {
		numCount[i+1], pre = pre, pre+numCount[i+1]
	}

	// 只录用在nums中出现的数字
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = numCount[num]
	}

	return res
}


// @lc code=end
