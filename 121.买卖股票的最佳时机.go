/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (53.96%)
 * Likes:    1005
 * Dislikes: 0
 * Total Accepted:    217.7K
 * Total Submissions: 401.9K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
 *
 * 注意：你不能在买入股票前卖出股票。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 */

// @lc code=start

// // 暴力解法--会超时
// func maxProfit(prices []int) int {
// 	profit := 0
// 	for i:=0; i<len(prices)-1; i++{
// 		for j:=i+1; j<len(prices); j++{
// 			c := prices[j]-prices[i]
// 			fmt.Println(c)
// 			if c > profit{
// 				profit = c
// 			}
// 		}
// 	}
// 	fmt.Println(profit)
// 	return profit
// }
func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	n := len(prices)
	dp := make([][]int, n)
	for i:=0; i<n; i++{
		dp[i] = make([]int, 2)
	}

	for i:=0; i<n; i++{
		if i == 0{
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}



















func maxProfitCopy(prices []int) int {
	if len(prices) == 0{
		return 0
	}

	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++{
		dp[i] = make([]int, 2)
	}

	for i:=0; i<len(prices); i++{
		if i == 0{
			// 第一天的情况
			dp[i][0] = 0  // 第一天卖出股票的收益
			dp[i][1] = 0 - prices[i]  // 第一天持有股票的收益
			continue
		}
		//         前一天就没有股票   昨天持有股票，今天卖掉了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]) // 没有持有股票
		//         前一天就持有股票  昨天没有持有股票， 今天刚买的股票
		dp[i][1] = max(dp[i-1][1], -prices[i]) // 持有股票
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}


// 一次遍历
func maxProfit1(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices{
		profit := price - minPrice
		if profit > maxProfit{
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}

// @lc code=end
