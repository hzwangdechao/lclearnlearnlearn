/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (49.75%)
 * Likes:    142
 * Dislikes: 0
 * Total Accepted:    21.7K
 * Total Submissions: 43.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,3]
 * 输出: 6
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,3,4]
 * 输出: 24
 *
 *
 * 注意:
 *
 *
 * 给定的整型数组长度范围是[3,10^4]，数组中所有的元素范围是[-1000, 1000]。
 * 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 *
 *
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	// 可能出现的情况： 三个最大数   一个最大数*两个最小数
	x := nums[0]*nums[1]*nums[len(nums)-1]
	y := nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3]
	if x > y {
		return x
	}
	return y

}
// @lc code=end
