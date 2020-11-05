/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 *
 * https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/description/
 *
 * algorithms
 * Medium (65.12%)
 * Likes:    231
 * Dislikes: 0
 * Total Accepted:    19.4K
 * Total Submissions: 29.4K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 *
 * 找到所有出现两次的元素。
 *
 * 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 *
 * 示例：
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [2,3]
 *
 *
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	// return withMap(nums)
	return withIndex(nums)

}

// 利用index找出出现两次的数据
func withIndex(nums []int)[]int  {

	res := make([]int, 0)
	for i:=0; i<len(nums); i++{
		abs := nums[i]
		if abs < 0{
			abs = abs * -1
		}
		loc := abs - 1
		fmt.Println(loc, nums[loc])
		if nums[loc] < 0{
			res = append(res, loc+1)
		}
		nums[loc] = -nums[loc]

	}

	return res

}


// 利用map找到出现两次的数据
func withMap(nums []int)[]int {
	res := make([]int, 0)

	m := make(map[int]int)
	for _, num := range nums{
		if _, ok := m[num]; !ok{
			m[num] = 1
		}else{
			m[num] += 1
		}
	}

	for k, v := range m{
		if v == 2{
			res = append(res, k)
		}
	}
	return res

}
// @lc code=end
