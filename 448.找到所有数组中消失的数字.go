/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (57.85%)
 * Likes:    393
 * Dislikes: 0
 * Total Accepted:    43.8K
 * Total Submissions: 74.1K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 *
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 *
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 *
 * 示例:
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [5,6]
 *
 *
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}

	// 将数组每个值放到map中
	m := make(map[int]bool, 0)
	for _, num := range nums{
		m[num] = true
	}

	res := make([]int, 0)
	for i:=1; i<=len(nums); i++{
		if _, ok := m[i]; !ok{
			res = append(res, i)
		}
	}
	return res

}
// @lc code=end
