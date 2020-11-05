/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (76.64%)
 * Likes:    845
 * Dislikes: 0
 * Total Accepted:    173.4K
 * Total Submissions: 226.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
func permute1(nums []int) [][]int {
	res := make([][]int, 0)
	traceBack(&res, nums, 0)

	return res
}

func traceBack(res *[][]int, nums []int, first int)  {
	// 做添加
	if first == len(nums){
		temp := make([]int, len(nums))
		fmt.Println(nums)
		copy(temp, nums)
		*res = append(*res, temp)
	}

	// 进行遍历
	for i:=first; i<len(nums); i++{
		nums[first], nums[i] = nums[i], nums[first]
		traceBack(res, nums, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	if n == 0{
		return res
	}
	visited := make([]bool, n) // 记录元素是否被访问过

	var dfs func([]int, int)
	dfs = func (path []int, idx int)  {
		if idx == n{
			fmt.Println(path)
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
		}

		for i:=0; i<n; i++{
			if visited[i]{
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
