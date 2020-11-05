/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 *
 * https://leetcode-cn.com/problems/create-maximum-number/description/
 *
 * algorithms
 * Hard (31.74%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 17.4K
 * Testcase Example:  '[3,4,6,5]\n[9,1,2,5,8,3]\n5'
 *
 * 给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n)
 * 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
 *
 * 求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
 *
 * 说明: 请尽可能地优化你算法的时间和空间复杂度。
 *
 * 示例 1:
 *
 * 输入:
 * nums1 = [3, 4, 6, 5]
 * nums2 = [9, 1, 2, 5, 8, 3]
 * k = 5
 * 输出:
 * [9, 8, 6, 5, 3]
 *
 * 示例 2:
 *
 * 输入:
 * nums1 = [6, 7]
 * nums2 = [6, 0, 4]
 * k = 5
 * 输出:
 * [6, 7, 6, 0, 4]
 *
 * 示例 3:
 *
 * 输入:
 * nums1 = [3, 9]
 * nums2 = [8, 9]
 * k = 3
 * 输出:
 * [9, 8, 9]
 *
 */

// @lc code=start
/*
	 题目拆解：
		 从nums1中取出i个数， 使其组成最大数， 其中 0<=i<=len(nums1)  记为l1
		 从nums2中取出k-i个数， 使其组成最大数  其中  0<=k-i<=len(nums2)  记为l2
		 将l1和l2利用归并排序的方式进行合并， 合并后的数组为 list
		 将list与res进行比较， 判断哪个组成的元素更大
*/

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// 从数组中选取n个数， 组成最大数

	// 合并两个数组

	res := []int{}
	for i := 0; i < k+1; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			l1 := pickMax(nums1, i)      // 从nums1中抽取i个元素
			l2 := pickMax(nums2, k-i)    // 从nums2中抽取k-i个元素
			list := mergeTwoList(l1, l2) // 两个数组进行合并
			if isBigger(list, res) {     // 如果当前合并后的数组比res大的话， 替换掉res
				res = list
			}
		}

	}
	return res
}

// 合并两个有序数组
func mergeTwoList(nums1, nums2 []int) []int {
	res := []int{}
	for len(nums1) > 0 && len(nums2) > 0 {
		n1 := nums1[0]
		n2 := nums2[0]
		if isBigger(nums1, nums2) { // 与基础的比较方式不同
			res = append(res, n1)
			nums1 = nums1[1:]
		} else {
			res = append(res, n2)
			nums2 = nums2[1:]
		}
	}
	if len(nums1) > 0 {
		res = append(res, nums1...)
	}
	if len(nums2) > 0 {
		res = append(res, nums2...)
	}

	return res
}

// 从nums中选出k个数， 使其组成的数字最大， nums中元素的相对位置保持不变
func pickMax(nums []int, k int) []int {
	drop := len(nums) - k // 需要从nums中删除元素的个数
	stack := []int{}
	for i := 0; i < len(nums); i++ {

		for len(stack) > 0 && stack[len(stack)-1] < nums[i] && drop > 0 {
			stack = stack[:len(stack)-1]
			drop--
		}

		stack = append(stack, nums[i])
	}
	return stack[:k]
}

// 判断哪个数组组成的数字更大
func isBigger(nums1, nums2 []int) bool {
	l1, l2 := len(nums1), len(nums2)
	n := min(l1, l2)
	for i := 0; i < n; i++ {
		if nums1[i] == nums2[i] {
			continue
		} else if nums1[i] > nums2[i] {
			return true
		} else {
			return false
		}
	}
	return l1 > l2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
