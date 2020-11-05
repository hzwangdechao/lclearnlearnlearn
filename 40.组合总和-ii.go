/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (62.74%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    81.9K
 * Total Submissions: 130.5K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。 
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

// @lc code=start
// func combinationSum2(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	return dfs(candidates, target, 0)
// }

// func dfs(nums []int, target int, start int) [][]int {
// 	res := [][]int{}

// 	for i:=start; i<len(nums); i++{
// 		if i > start && nums[i] == nums[i-1]{
// 			continue
// 		}
// 		if nums[i] > target{
// 			return res
// 		}

// 		if nums[i] == target{
// 			res = append(res, []int{nums[i]})
// 			continue
// 		}

// 		for _, v := range dfs(nums, target-nums[i], i+1){
// 			res = append(res, append([]int{nums[i]}, v...))
// 		}
// 	}

// 	return res
// }


// func combinationSum2(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	res := [][]int{}

// 	var dfs func(int, int, []int)
// 	dfs = func (idx int, target int, path []int)  {
// 		if idx >= len(candidates){
// 			return
// 		}
// 		if target < 0{
// 			return
// 		}

// 		if target == 0{
// 			res = append(res, append([]int{}, path...))
// 			return
// 		}

// 		dfs(idx+1, target, path)
// 		dfs(idx+1, target-candidates[idx], append(path, candidates[idx]))
// 		return
// 	}
// 	dfs(0, target, []int{})
// 	return res
// }


func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	var dfs func([]int, int, int)
	dfs = func (path []int, start int, sum int)  {
		if sum >= target{
			if sum == target{
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}

		for i:=start; i<len(candidates); i++{
			// 如果当前元素和上一个元素重复的话， 则跳过此次循环
			if i > start && candidates[i] == candidates[i-1]{
				continue
			}

			// 做选择
			path = append(path,candidates[i])
			// 继续向下遍历
			dfs(path, i+1, sum+candidates[i])
			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	dfs([]int{}, 0, 0)
	return res
}
// @lc code=end
