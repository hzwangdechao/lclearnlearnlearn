/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 *
 * https://leetcode-cn.com/problems/last-stone-weight-ii/description/
 *
 * algorithms
 * Medium (46.84%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    4.9K
 * Total Submissions: 10.4K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * 有一堆石头，每块石头的重量都是正整数。
 *
 * 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
 *
 *
 * 如果 x == y，那么两块石头都会被完全粉碎；
 * 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
 *
 *
 * 最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。
 *
 *
 *
 * 示例：
 *
 * 输入：[2,7,4,1,8,1]
 * 输出：1
 * 解释：
 * 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
 * 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
 * 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
 * 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 *
 *
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones{
		sum += num
	}
	avg := sum / 2

	n := len(stones)
	// dp := make([][]int, n+1)
	// for i:=0; i<=n; i++{
	// 	dp[i] = make([]int, avg+1)
	// }

	dp := make([]int, avg+1)

	for i:=1; i<=n; i++{
		for j:=1; j<=avg; j++{
			left := j-stones[i-1]  // 剩余容量
			if left < 0{
				dp[j] = dp[j]
			}else{
				r1 := dp[j]
				r2 := dp[left] + stones[i-1]
				if r1 > r2{
					dp[j] = r1
				}else{
					dp[j] = r2
				}
			}

		}
	}

	return sum - dp[avg]*2
}


// @lc code=end
