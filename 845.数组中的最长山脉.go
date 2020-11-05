/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 *
 * https://leetcode-cn.com/problems/longest-mountain-in-array/description/
 *
 * algorithms
 * Medium (36.63%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    18.7K
 * Total Submissions: 46.4K
 * Testcase Example:  '[2,1,4,7,3,2,5]'
 *
 * 我们把数组 A 中符合下列属性的任意连续子数组 B 称为 “山脉”：
 *
 *
 * B.length >= 3
 * 存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... >
 * B[B.length - 1]
 *
 *
 * （注意：B 可以是 A 的任意子数组，包括整个数组 A。）
 *
 * 给出一个整数数组 A，返回最长 “山脉” 的长度。
 *
 * 如果不含有 “山脉” 则返回 0。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[2,1,4,7,3,2,5]
 * 输出：5
 * 解释：最长的 “山脉” 是 [1,4,7,3,2]，长度为 5。
 *
 *
 * 示例 2：
 *
 * 输入：[2,2,2]
 * 输出：0
 * 解释：不含 “山脉”。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 *
 *
 */

// @lc code=start
func longestMountain(A []int) int {
	n := len(A)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		// 从左向右遍历数组， 记录左侧山脉的长度 (左侧山脉的元素需要是递增的)
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		// 从右向左遍历数组， 记录右侧山脉的长度（右侧山脉的元素需要是递减的）
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	var res int
	// 遍历两个数组, 如果left[i] 和 right[i] 同时大于0， 说明出现的山脉
	for i, l := range left {
		r := right[i]

		if l > 0 && r > 0 && l+r+1 > res {
			res = l + r + 1
		}
	}
	return res
}

// @lc code=end
