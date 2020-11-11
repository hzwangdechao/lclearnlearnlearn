/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 *
 * https://leetcode-cn.com/problems/reverse-pairs/description/
 *
 * algorithms
 * Hard (28.97%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    7.9K
 * Total Submissions: 27K
 * Testcase Example:  '[1,3,2,3,1]'
 *
 * 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
 *
 * 你需要返回给定数组中的重要翻转对的数量。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,2,3,1]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,4,3,5,1]
 * 输出: 3
 *
 *
 * 注意:
 *
 *
 * 给定数组的长度不会超过50000。
 * 输入数组中的所有数字都在32位整数的表示范围内。
 *
 *
 */

// @lc code=start
func reversePairs(nums []int) int {
	// 利用归并排序的方法求解翻转对的格式

	var mergeCount func(arr []int) int
	mergeCount = func(arr []int) int {
		if len(arr) <= 1 {
			return 0
		}
		n := len(arr)
		n1 := append([]int{}, arr[:n/2]...) // 左半部分
		n2 := append([]int{}, arr[n/2:]...) // 右半部分
		cnt := mergeCount(n1) + mergeCount(n2)

		// 计算个数
		l, r := 0, 0
		for l < len(n1) && r < len(n2) {
			if n1[l] > 2*n2[r] {
				cnt += len(n1) - l // 如果n1[l] > n2[r]*2 的话， 说明l之后的元素也满足这个条件， 所以要加上len(n1)-l
				r++
			} else {
				l++
			}
		}
		// 对数组进行排序
		i, j := 0, 0
		for x := 0; x < n; x++ {
			if i < len(n1) && (j == len(n2) || n1[i] <= n2[j]) {
				arr[x] = n1[i]
				i++
			} else {
				arr[x] = n2[j]
				j++
			}
		}

		return cnt
	}

	return mergeCount(nums)
}

// @lc code=end
