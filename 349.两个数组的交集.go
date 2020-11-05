/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (69.14%)
 * Likes:    178
 * Dislikes: 0
 * Total Accepted:    65.5K
 * Total Submissions: 94.7K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	// return hashMap(nums1, nums2)

	// 下面是二分查找的方法

	res := []int{}
	if len(nums1) == 0 || len(nums2) == 0{
		return res
	}
	var short,long []int
	// 找出长数组和短数组
	if len(nums1) >= len(nums2){
		short, long = nums2, nums1
	}else{
		short, long = nums1, nums2
	}
	// 对短数组进行排序
	sort.Ints(short)

	for _, num := range long{
		if binarySearch(short, num) && !contains(res, num){
			res = append(res, num)
		}
	}
	return res

}

func binarySearch(nums []int, target int)bool  {
	left, right := 0, len(nums)-1
	for left <= right{
		mid := (right-left)/2+left
		if nums[mid] == target{
			return true
		}else if nums[mid] > target{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return false
}

func contains(nums []int, target int) bool {
	for _, num := range nums{
		if num == target{
			return true
		}
	}
	return false
}



func hashMap(nums1 []int, nums2[]int)[]int  {
	res := []int{}
	set := make(map[int]bool)
	for _, num := range nums1{
		set[num] = true
	}

	for _, num := range nums2{
		if k, ok := set[num]; ok && k{
			res = append(res, num)
			// 将值改为false， 防止重复添加
			set[num] = false
		}
	}
	return res
}

// @lc code=end
