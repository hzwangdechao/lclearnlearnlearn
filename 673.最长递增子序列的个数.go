/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 *
 * https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (35.69%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    11.3K
 * Total Submissions: 31.6K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未排序的整数数组，找到最长递增子序列的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,5,4,7]
 * 输出: 2
 * 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,2,2,2,2]
 * 输出: 5
 * 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
 *
 *
 * 注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。
 *
 */

// @lc code=start

type s struct{
	Length int
	Count int
}

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	ans := 0
	maxIndex := 0 // 记录最长子序列的位置
	dp := []s{}

	for i:=0; i<len(nums); i++{
		cur := s{1, 1}
		dp = append(dp, cur)

		for j:=0; j<i; j++{
			if nums[i] > nums[j]{
				if dp[j].Length + 1 > dp[i].Length{
					dp[i].Length = dp[j].Length + 1
					dp[i].Count = dp[j].Count
				}else if dp[j].Length + 1 == dp[i].Length{
					dp[i].Count += dp[j].Count
				}
			}
		}
		fmt.Println(dp)
		if dp[i].Length > dp[maxIndex].Length{
			maxIndex = i
			ans = dp[i].Count
		}else if dp[i].Length == dp[maxIndex].Length{
			ans += dp[i].Count
		}
	}
	return ans
}



// @lc code=end
