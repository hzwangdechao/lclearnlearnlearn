/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.55%)
 * Likes:    675
 * Dislikes: 0
 * Total Accepted:    111.6K
 * Total Submissions: 143.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})  // 将空集存放到结果数组中

	for i:=0; i<len(nums); i++{
		tempArr := make([][]int, 0)
		for _, c := range res{
			temp := make([]int, 0)
			// 取出原始元素
			temp = append(temp, c...)
			// 将新元素添加到temp中
			temp = append(temp, nums[i])
			// 将temp放到tempArr中去
			tempArr = append(tempArr, temp)
		}

		// 将tempArr中的元素添加到res中
		for _, arr := range tempArr{
			res = append(res, arr)
		}
	}

	return res

}
// @lc code=end
