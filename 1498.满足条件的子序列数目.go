/*
 * @lc app=leetcode.cn id=1498 lang=golang
 *
 * [1498] 满足条件的子序列数目
 *
 * https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/description/
 *
 * algorithms
 * Medium (31.31%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 8K
 * Testcase Example:  '[3,5,6,7]\n9'
 *
 * 给你一个整数数组 nums 和一个整数 target 。
 *
 * 请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的 非空 子序列的数目。
 *
 * 由于答案可能很大，请将结果对 10^9 + 7 取余后返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,5,6,7], target = 9
 * 输出：4
 * 解释：有 4 个子序列满足该条件。
 * [3] -> 最小元素 + 最大元素 <= target (3 + 3 <= 9)
 * [3,5] -> (3 + 5 <= 9)
 * [3,5,6] -> (3 + 6 <= 9)
 * [3,6] -> (3 + 6 <= 9)
 *
 *
 * 示例 2：
 *
 * 输入：nums = [3,3,6,8], target = 10
 * 输出：6
 * 解释：有 6 个子序列满足该条件。（nums 中可以有重复数字）
 * [3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]
 *
 * 示例 3：
 *
 * 输入：nums = [2,3,3,4,6,7], target = 12
 * 输出：61
 * 解释：共有 63 个非空子序列，其中 2 个不满足条件（[6,7], [7]）
 * 有效序列总数为（63 - 2 = 61）
 *
 *
 * 示例 4：
 *
 * 输入：nums = [5,2,4,1,7,6,8], target = 16
 * 输出：127
 * 解释：所有非空子序列都满足条件 (2^7 - 1) = 127
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 * 1 <= target <= 10^6
 *
 *
 */

// @lc code=start
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)
	// p := 1e9+7
	ans := 0
	pow := map[int]int{0:1}

	for i:=1; i<n; i++{
		pow[i] = (pow[i-1]*2) % (1e9+7)
	}

	i, j := 0, len(nums)-1
	for i<=j{
		for nums[i] + nums[j] >target{
			if j == 0{
				return ans
			}
			j --
		}
		ans += pow[j-i]
		ans %=  1e9+7
		i++
	}
	return ans

}
// @lc code=end

// [2,3,3,4,6,7]
// 2^4 = 16
// 2^5 = 32

// 32 + 16 + 6 + 8 + 4 +


// func numSubseq(nums []int, target int) int {
// 	sort.Ints(nums)
// 	i, j, ans, pow := 0, len(nums)-1, 0, map[int]int{0: 1}

// 	for n := 1; n < len(nums); n++ {
// 		pow[n] = (pow[n-1] * 2) % (1e9 + 7)
// 	}

// 	for j >= i {
// 		for nums[i]+nums[j] > target {
// 			if j == 0 {
// 				// 从右边开始找到第一个元素， 还是比target大， 说明直接返回ans
// 				return ans
// 			}
// 			j--
// 		}
// 		ans += pow[j-i]
// 		ans %= 1e9 + 7
// 		i++
// 	}
// 	return ans
// }
