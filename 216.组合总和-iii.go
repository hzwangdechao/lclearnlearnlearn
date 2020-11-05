/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (71.75%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    30.3K
 * Total Submissions: 42.1K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 *
 * 说明：
 *
 *
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。 
 *
 *
 * 示例 1:
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 *
 *
 * 示例 2:
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */

// @lc code=start

func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := n
	res := [][]int{}
	var dfs func(path []int, start int, sum int)
	dfs = func (path []int, start int, sum int)  {
		if sum >= target{
			if sum == target && len(path) == k{
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
			}
		}

		for i:=start; i<len(nums); i++{
			if i > start && nums[i] == nums[i-1]{
				continue
			}

			// 做选择
			path = append(path, nums[i])
			dfs(path, i+1, sum+nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0)
	return res
}
// @lc code=end
