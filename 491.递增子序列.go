/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (49.00%)
 * Likes:    118
 * Dislikes: 0
 * Total Accepted:    10.8K
 * Total Submissions: 21.1K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
 *
 * 示例:
 *
 *
 * 输入: [4, 6, 7, 7]
 * 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7],
 * [4,7,7]]
 *
 * 说明:
 *
 *
 * 给定数组的长度不会超过15。
 * 数组中的整数范围是 [-100,100]。
 * 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
 *
 *
 */

// @lc code=start


//  递归枚举+剪枝
var (
	temp []int
	ans [][]int
)

func findSubsequences(nums []int) [][]int {
	ans = [][]int{}
    dfs(0, math.MinInt32, nums)

	return ans
}

func dfs(cur, last int, nums []int)  {
	if cur == len(nums){
		// 最少需要两个元素
		if len(temp) >=2{
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}

	if nums[cur] >= last{
		// 如果当前元素比假定的last元素大的话， 则将其插入到temp数组中
		temp = append(temp, nums[cur])
		// 继续向后遍历
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[cur] != last{
		dfs(cur+1, last, nums)
	}

}


// @lc code=end
