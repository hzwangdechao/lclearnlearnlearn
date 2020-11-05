/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (40.86%)
 * Likes:    757
 * Dislikes: 0
 * Total Accepted:    120.2K
 * Total Submissions: 293.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 *
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 *
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i:=0; i<=n; i++{
		dp[i] = make([]int, amount+1)
	}
	for i:=0; i<=amount; i++{
		dp[0][i] = amount+1
	}
	dp[0][0] = 0
	fmt.Println(dp)
	for i:=1; i<=n; i++{
		for j:=1; j<=amount; j++{
			left := j - coins[i-1]
			if left < 0{
				dp[i][j] = dp[i-1][j]
			}else{
				dp[i][j] = min(dp[i-1][j], 1+dp[i][left])
			}
		}
	}
	fmt.Println(dp)
	if dp[n][amount] > amount{
		return -1
	}else{
		return dp[n][amount]
	}
}

func min(x, y int) int {
	if x < y{
		return x
	}
	return y
}

// 一维数组的方法
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i:=0; i<=amount; i++{
		dp[i] = amount+1
	}
	dp[0] = 0

	for i:=1; i<=len(coins); i++{
		for j:=1; j<=amount; j++{
			left := j - coins[i-1]
			if left >= 0{
				dp[j] = min(dp[j], dp[left]+1)
			}
		}
	}
	if dp[amount] > amount{
		return -1
	}
	return dp[amount]
}
// @lc code=end
