/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (47.47%)
 * Likes:    496
 * Dislikes: 0
 * Total Accepted:    145.2K
 * Total Submissions: 305.4K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 *
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	n1 := m -1
	n2 := n -1
	l := m+n -1
	for n2 >=0{
		if n1 >=0 && nums1[n1] > nums2[n2] {
			// 将比较大的放大数放到nums1的后面
			nums1[l] = nums1[n1]
			n1 --
		}else{
			nums1[l] = nums2[n2]
			n2 --
		}
		l --
	}
}
// @lc code=end
