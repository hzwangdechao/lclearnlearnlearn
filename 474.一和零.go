/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 *
 * https://leetcode-cn.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (54.98%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    12.9K
 * Total Submissions: 23.5K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * 在计算机界中，我们总是追求用有限的资源获取最大的收益。
 *
 * 现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。
 *
 * 你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。
 *
 * 注意:
 *
 *
 * 给定 0 和 1 的数量都不会超过 100。
 * 给定字符串数组的长度不会超过 600。
 *
 *
 * 示例 1:
 *
 *
 * 输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
 * 输出: 4
 *
 * 解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
 *
 *
 * 示例 2:
 *
 *
 * 输入: Array = {"10", "0", "1"}, m = 1, n = 1
 * 输出: 2
 *
 * 解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。
 *
 *
 */

// @lc code=start

// 背包空间大小m
// 最大载重n
// len(strs)个物品

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i:=0; i<=m; i++{
		rows := make([]int, n+1)
		dp[i] = rows
	}

	for _, str := range strs{
		// 用来统计0和1的数量
		mp :=make(map[int32]int)
		for _, s := range str{
			mp[s] += 1
		}
		count0 := mp['0']
		count1 := mp['1']



		for i:=m; i>=mp['0']; i--{
			for j:=n; j>=mp['1']; j--{
				dp[i][j] = max(dp[i][j], 1+dp[i-count0][j-count1])
			}
		}

	}

	fmt.Println(dp)
	return dp[m][n]

}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y

}
// @lc code=end
