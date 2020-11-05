/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (52.12%)
 * Likes:    569
 * Dislikes: 0
 * Total Accepted:    81K
 * Total Submissions: 155.4K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	var (
		curNum       int // 当前的数字
		curNumLength int // 当前数字的长度
		maxNumLength int // 最长的连续序列的长度
	)
	for _, num := range nums {
		if !numSet[num-1] {  // 如果当前数字不是连续数字的话
			curNum = num
			curNumLength = 1  // 当前数字开头的连续数字的长度

			for numSet[curNum+1] {
				curNumLength += 1
				curNum += 1
			}

			if curNumLength > maxNumLength {
				maxNumLength = curNumLength
			}
		}
	}

	return maxNumLength
}

// @lc code=end
