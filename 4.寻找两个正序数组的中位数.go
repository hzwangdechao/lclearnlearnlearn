/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (37.35%)
 * Likes:    2512
 * Dislikes: 0
 * Total Accepted:    181.5K
 * Total Submissions: 485.9K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 两个数组的总长度
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1{
		// 长度为奇数的情况， 取中间数
		mid := totalLength/2
		return float64(getKthNum(nums1, nums2, mid+1))

	}else{
		// 长度为偶数的情况， 取两个数的平均值
		mid1, mid2 := totalLength/2 - 1, totalLength/2
		return (float64(getKthNum(nums1, nums2, mid1+1)) + float64(getKthNum(nums1, nums2, mid2+1))) / 2.0
	}

	return 0
}

// 寻找两个数组中第k个元素
func getKthNum(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1){
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2){
			return nums1[index1 + k - 1]
		}
		if k == 1{
			return min(nums1[index1], nums2[index2])
		}

		half := k/2
		newIndex1 := min(index1+half, len(nums1)) -1 // 防止出现index 比 数组长度大的情况
		newIndex2 := min(index2+half, len(nums2)) -1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 <= pivot2{
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		}else{
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}


func min(x, y int) int {
	if x < y{
		return x
	}
	return y
}


// @lc code=end
