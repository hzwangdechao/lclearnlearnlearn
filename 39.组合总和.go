/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
  Medium (68.88%)
 * Likes:    647
 * Dislikes: 0
 * Total Accepted:    89.9K
 * Total Submissions: 130.6K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */

// @lc code=start
// func combinationSum2(candidates []int, target int) [][]int {
// 	// 先对数组进行排序， 返回进行递归判断
// 	sort.Ints(candidates)
// 	return dfs(candidates, target)

// }

// func dfs(nums []int, target int) [][]int  {
// 	res := [][]int{}
// 	for i, num := range nums{
// 		if target - num < 0{
// 			break
// 		}
// 		if target - num == 0{
// 			res = append(res, []int{num})
// 			continue
// 		}

// 		for _, v := range dfs(nums[i:], target-num){
// 			res = append(res, append([]int{num}, v...))
// 		}
// 	}
// 	return res

// }


// func combinationSum(candidates []int, target int) (ans [][]int) {
// 	comb := []int{}
// 	var dfs func(target, idx int)
// 	dfs = func(target, idx int) {
// 		if idx == len(candidates) {
// 			return
// 		}
// 		if target == 0 {
// 			ans = append(ans, append([]int(nil), comb...))
// 			return
// 		}
// 		// 直接跳过
// 		dfs(target, idx+1)
// 		// 选择当前数
// 		if target-candidates[idx] >= 0 {
// 			comb = append(comb, candidates[idx])
// 			dfs(target-candidates[idx], idx)
// 			comb = comb[:len(comb)-1]
// 		}
// 	}
// 	dfs(target, 0)
// 	return
// }
// func combinationSum(candidates []int, target int) [][]int {
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
// 		dfs(idx, target-candidates[idx], append(path, candidates[idx]))
// 		return
// 	}
// 	dfs(0, target, []int{})
// 	return res
// }
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{} // 存放最终的结果
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

		for i:= start; i<len(candidates); i++{
			// 做选择
			path = append(path, candidates[i])
			// 继续遍历
			dfs(path, i, sum+candidates[i])
			// 撤销选择
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0)
	return res
}

// @lc code=end
