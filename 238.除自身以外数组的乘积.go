/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 *
 * https://leetcode-cn.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (70.66%)
 * Likes:    564
 * Dislikes: 0
 * Total Accepted:    75.8K
 * Total Submissions: 107.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i]
 * 之外其余各元素的乘积。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,2,3,4]
 * 输出: [24,12,8,6]
 *
 *
 *
 * 提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
 *
 * 说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
 *
 * 进阶：
 * 你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	return method2(nums)
}

func method2(nums []int) []int {
	n := len(nums)
	if n == 0{
		return []int{}
	}
	ans := make([]int, n)

	// 首先计算当前元素的左侧元素的乘积
	ans[0] = 1
	for i:=1; i<n; i++{
		ans[i] = nums[i-1] * ans[i-1]
	}

	// 然后从后向前计算当前元素的右侧元素的乘积
	r := 1
	for i:=n-1; i>=0; i--{
		ans[i] = ans[i] * r

		r *= nums[i]
	}

	return ans

}

// 初始化两个数组用来存放元素左侧元素的乘积和右侧元素的乘积
func method1(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}
	if len(nums) == 1{
		return nums
	}
	n := len(nums)
	L, R , ans := make([]int, n), make([]int, n), make([]int, n)
	L[0] = 1
	for i:=1; i<n; i++{
		L[i] = L[i-1] * nums[i-1]
	}
	fmt.Println(L)

	R[n-1] = 1
	for i:=n-2; i>=0; i--{
		R[i] = R[i+1] * nums[i+1]
	}
	fmt.Println(R)
	for i:=0; i<n; i++{
		ans[i] = L[i] * R[i]
	}
	return ans

}
// @lc code=end
