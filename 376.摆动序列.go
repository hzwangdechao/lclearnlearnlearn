/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 *
 * https://leetcode-cn.com/problems/wiggle-subsequence/description/
 *
 * algorithms
 * Medium (42.67%)
 * Likes:    259
 * Dislikes: 0
 * Total Accepted:    27.9K
 * Total Submissions: 65.4K
 * Testcase Example:  '[1,7,4,9,2,5]'
 *
 * 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。
 *
 * 例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和
 * [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
 *
 * 给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
 *
 * 示例 1:
 *
 * 输入: [1,7,4,9,2,5]
 * 输出: 6
 * 解释: 整个序列均为摆动序列。
 *
 *
 * 示例 2:
 *
 * 输入: [1,17,5,10,13,15,10,5,16,8]
 * 输出: 7
 * 解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
 *
 * 示例 3:
 *
 * 输入: [1,2,3,4,5,6,7,8,9]
 * 输出: 2
 *
 * 进阶:
 * 你能否用 O(n) 时间复杂度完成此题?
 *
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	return greedy(nums)
}

// 线性动态规划
func lineDp(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	n := len(nums)
	up := make([]int, n)   // up[i] 以nums[i]结尾的最长上升摆动的长度
	down := make([]int, n) // down[i] 以nums[i]结尾的最长下降摆动的长度
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] { // 上升摆动
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] { // 下降摆动
			down[i] = up[i-1] + 1
			up[i] = up[i-1]
		} else if nums[i] == nums[i-1] { // 没有摆动
			down[i] = down[i-1]
			up[i] = up[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}

// 优化后的动态规划， up和down只存放一个数字
func optimizeDp(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	n := len(nums)
	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] { // 上升摆动， 增加up的大小
			up = down + 1
		} else if nums[i] < nums[i-1] { // 下降摆动， 增加down的大小
			down = up + 1
		}
	}
	return max(up, down)
}

// 贪心算法， 只在出现波峰的时候增加数量
func greedy(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	pre := 0
	res := 1
	for i := 1; i < len(nums); i++ {
		cur := nums[i] - nums[i-1]

		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0){
			res ++
			pre = cur
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
