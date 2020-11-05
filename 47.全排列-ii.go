/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (59.60%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    80.9K
 * Total Submissions: 135.6K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */

// @lc code=start
func permuteUnique2(nums []int) [][]int {
	res := make([][]int, 0)
	traceBack(&res, nums, 0)
	return res
}

func traceBack(res *[][]int, nums []int, first int)  {
	if first == len(nums){
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	// seen := map[int]bool{}
	for i:=first; i < len(nums); i++{
		if !check(nums, first, i){
			continue
		}
		// if !seen[nums[i]]{
			nums[i], nums[first] = nums[first], nums[i]
			traceBack(res, nums, first+1)
			nums[i], nums[first] = nums[first], nums[i]
			// seen[nums[i]] = true
		// }
	}
}

func check(nums []int, start int, end int) bool {

	for i:=start; i<end; i++{
		if nums[i] == nums[end]{
			return false
		}
	}
	return true
}


func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	if n == 0{
		return res
	}
	visited := make([]bool, n)

	var dfs func([]int, int)
	dfs = func (path []int, idx int)  {
		if idx == n{
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i:=0; i<n; i++{
			if (visited[i]) || (i>0 && !visited[i-1] && nums[i-1] == nums[i]){
				continue
			}

			visited[i] = true
			dfs(append(path, nums[i]), idx+1)
			visited[i] = false
		}
	}

	dfs([]int{}, 0)
	return res
}

// @lc code=end
