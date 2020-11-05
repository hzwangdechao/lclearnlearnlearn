/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (48.11%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    82.6K
 * Total Submissions: 171.5K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2,2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [4,9]
 *
 * 说明：
 *
 *
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 * 进阶:
 *
 *
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 *
 *
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	// set 用于计数
	set := make(map[int]int)
	for _, num := range nums1{
		if _, ok := set[num]; !ok{
			// 第一次出现
			set[num] = 1
		}else{
			set[num] ++
		}
	}

	for _, num  := range nums2{
		if k, ok := set[num]; ok && k > 0{
			// 如果num的值在map中存在， 并且数量大于0， 添加结果， 并将数量-1
			res = append(res, num)
			set[num] --
		}

	}
	return res
}

// 排序双指针法
// 核心： 首先对数组进行排序， 然后双指针进行迭代检索
// funcintersect(nums1 []int, nums2 []int) []int {
//		// 双指针的前提是数组是排好序的数组
// 	    sort.Ints(nums1)
// 		sort.Ints(nums2)
// 		res := []int{}
// 		lOne, lTwo, rOne, rTwo := 0, 0, len(nums1), len(nums2)
// 		for lOne < rOne && lTwo < rTwo {
// 			if nums1[lOne] == nums2[lTwo] {
// 				res = append(res, nums1[lOne])
// 				lOne++
// 				lTwo++
// 			} else if nums1[lOne] < nums2[lTwo] {
//			// 如果nums1 偏小的话， 将nums1的位置向后移一个
// 				lOne++
// 			} else {
// 				lTwo++
// 			}
// 		}
// 		return res
// 	}

// @lc code=end
