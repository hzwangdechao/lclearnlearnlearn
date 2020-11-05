/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode-cn.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (48.91%)
 * Likes:    363
 * Dislikes: 0
 * Total Accepted:    48.6K
 * Total Submissions: 99.2K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 * 注意:
 *
 *
 * 每个数组中的元素不会超过 100
 * 数组的大小不会超过 200
 *
 *
 * 示例 1:
 *
 * 输入: [1, 5, 11, 5]
 *
 * 输出: true
 *
 * 解释: 数组可以分割成 [1, 5, 5] 和 [11].
 *
 *
 *
 *
 * 示例 2:
 *
 * 输入: [1, 2, 3, 5]
 *
 * 输出: false
 *
 * 解释: 数组不能分割成两个元素和相等的子集.
 *
 *
 *
 *
 */

// @lc code=start

// [1 5 11 5]
// sum = 11
// n = 4
// // dp[4][9]=true  >>>  容量为9的背包， 若只用前4个物品, 可以有一种方法可以加好把背包装满
// [
// 	   0 1 2 3 4 5 6 7 8 9 10 11
//  	0 [t f f f f f f f f f f f]
// 	1 [t t f f f f f f f f f f]
// 	2 [t f f f f f f f f f f f]
// 	3 [t f f f f f f f f f f f]
// 	4 [t f f f f f f f f f f f]
// ]

func canPartition(nums []int) bool {
}

func method1(nums []int) bool {
	var (
		sum = 0         // 数组中元素的总和
		n   = len(nums) // 数组的长度
		dp  = make([][]bool, n+1)
	)

	for _, num := range nums {
		sum += num
	}
	// 如果和为奇数的话不能平分成两个等和子集
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	for i := 0; i < n+1; i++ {
		x := make([]bool, sum+1)
		x[0] = true
		dp[i] = x
	}

	// dp[4][9] = true  >>> 容量为9的背包， 如果只用前四个物品， 可以有一种方法把背包装满
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ { // j 背包的总容量  nums[i-1] 背包中的第i-1个物品   j-nums[i-1] 背包的剩余容量
			left := j - nums[i-1]
			if left < 0 { // 背包的剩余容量不足以装入第i-1个物品
				dp[i][j] = dp[i-1][j]
			} else { // 背包的剩余容量可以装下第i-1个物品
				// 选择不装入或者装入当前物品
				dp[i][j] = dp[i-1][j] || dp[i-1][left]
			}
		}
	}
	return dp[n][sum]
}

func method2(nums []int) bool {
	var (
		sum = 0
		n   = len(nums)
	)

	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			left := j - nums[i-1] //背包的剩余容量
			if left >= 0 {
				dp[j] = dp[j] || dp[left]
			}
		}
	}
	return dp[sum]
}

// @lc code=end
