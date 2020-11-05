/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (44.45%)
 * Likes:    896
 * Dislikes: 0
 * Total Accepted:    142.5K
 * Total Submissions: 312K
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

// @lc code=start
func rob(nums []int) int {

	return rob2(nums)
	// return rob1(nums)

}

func rob2(nums []int)int  {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])  // 存放最大的钱数

	for i:=2; i<len(nums); i++{
		first, second = second, max(second, nums[i]+first)
	}
	return second

}

func rob1(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	if len(nums) ==2{
		return max(nums[0], nums[1])
	}

	// 存放每次偷取时的最大利润
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i:=2; i < len(nums); i++{
		// 前一家可偷取到的最大钱数    偷取这一家可得到的钱数
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}


func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}

// @lc code=end
