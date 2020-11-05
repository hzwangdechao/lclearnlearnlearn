/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (53.54%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    11.9K
 * Total Submissions: 22K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动将会使 n - 1 个元素增加 1。
 *
 *
 *
 * 示例:
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 3
 *
 * 解释:
 * 只需要3次移动（注意每次移动会增加两个元素的值）：
 *
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 *
 *
 */

// @lc code=start
func minMoves(nums []int) int {
	res := 0

	// 先对数组进行排序
	sort.Ints(nums)

	min := nums[0]

	for i:=len(nums)-1; i>0; i--{
		diff := nums[i] - min
		res += diff
	}

	return res
}
// @lc code=end
