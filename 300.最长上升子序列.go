/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (45.08%)
 * Likes:    927
 * Dislikes: 0
 * Total Accepted:    135K
 * Total Submissions: 298.9K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	return recursive(nums)
}

func binarySearch(nums []int) int {
	dp := []int{}
	for _, num := range nums{
		if len(dp) == 0 || dp[len(dp)-1] < num{
			// 如果dp为空， 或者dp的最后一个元素比当前元素小的话， 则将当前元素添加到dp的后面
			dp = append(dp, num)
		}else{
			// 利用二分查找， 找到dp中第一个比dp小的元素进行替换
			left, right := 0, len(dp)-1
			loc := right
			for left <= right{
				mid := (right-left)/2+left
				if dp[mid] > num{
					loc = mid
					right = mid -1
				}else if dp[mid] == num{
					loc = mid
					right = mid -1
				}else if dp[mid] < num{
					left = mid + 1
				}
			}

			dp[loc] = num
		}
	}
	return len(dp)
}



func recursive(nums []int) int {
	n := len(nums)
	dp := []int{}
	ans := 0
	for i:=0; i<n; i++{
		dp = append(dp, 1)

		for j:=0; j<i; j++{
			if nums[i] > nums[j]{
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	fmt.Println(dp)
	return ans
}

func max(x, y int ) int {
	if x > y{
		return x
	}
	return y
}

// @lc code=end
