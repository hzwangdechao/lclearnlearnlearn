/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (41.51%)
 * Likes:    471
 * Dislikes: 0
 * Total Accepted:    36.7K
 * Total Submissions: 88.4K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于
 * nums[i] 的元素的数量。
 *
 *
 *
 * 示例：
 *
 * 输入：nums = [5,2,6,1]
 * 输出：[2,1,1,0]
 * 解释：
 * 5 的右侧有 2 个更小的元素 (2 和 1)
 * 2 的右侧仅有 1 个更小的元素 (1)
 * 6 的右侧有 1 个更小的元素 (1)
 * 1 的右侧有 0 个更小的元素
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start
func countSmaller(nums []int) []int {
	return methodWithMergeSort(nums)
}

// 归并排序的方法
var (
	index    []int
	tmp      []int
	tmpIndex []int
	res      []int
)

func methodWithMergeSort(nums []int) []int {
	index = make([]int, len(nums))
	tmp = make([]int, len(nums))
	tmpIndex = make([]int, len(nums))
	res = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index[i] = i
	}
	left, right := 0, len(nums)-1
	mergeSort(nums, left, right)
	list := make([]int, 0)
	for _, num := range res {
		list = append(list, num)
	}
	return list
}

func mergeSort(nums []int, left, right int) {
	if left >= right { // 一个元素或者没有元素的情况不需要进行归排序
		return
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	l := left
	r := mid + 1
	p := left
	for l <= mid && r <= right {
		if nums[l] <= nums[r] {
			tmp[p] = nums[l]
			tmpIndex[p] = index[l]
			res[index[l]] += (r - mid - 1)
			l++
			p++
		} else {
			tmp[p] = nums[r]
			tmpIndex[p] = index[r]
			r++
			p++
		}
	}
	for l <= mid {
		tmp[p] = nums[l]
		tmpIndex[p] = index[l]
		res[index[l]] += (r - mid - 1)
		l++
		p++
	}

	for r <= right {
		tmp[p] = nums[r]
		tmpIndex[p] = index[r]
		r++
		p++
	}

	for k := left; k <= right; k++ {
		index[k] = tmpIndex[k]
		nums[k] = tmp[k]
	}

}

// 二分查找的方法
func methodWithBinarySearch(nums []int) []int {
	var (
		res          = make([]int, len(nums)) // 存放最终的结果
		sorted       []int                    // 存放排序数组
		binarySearch func(num int) int        // num 在sorted数组中的位置即为 右侧小于当前元素的个数
	)
	binarySearch = func(num int) int {
		// sorted数组为空的话
		if len(sorted) == 0 {
			sorted = append(sorted, num)
			return 0
		}

		left, right := 0, len(sorted)-1
		for left < right {
			mid := (right + left) >> 1
			if sorted[mid] < num {
				left = mid + 1
			} else if sorted[mid] == num {
				right = mid
			} else if sorted[mid] > num {
				right = mid
			}
		}

		// 如果排序数组的最后一个元素比num小的话， 直接将num添加到sorted的尾部
		if sorted[right] < num {
			sorted = append(sorted, num)
			return right + 1
		}

		// 将num添加到sorted的right位置
		tmp := append([]int{}, sorted[:right]...)
		tmp = append(tmp, num)
		tmp = append(tmp, sorted[right:]...)
		sorted = tmp
		// 相当于搜索插入位置的算法， 返回元素在排序数组中第一个插入位置
		return right
	}

	for i := len(nums) - 1; i >= 0; i-- {
		idx := binarySearch(nums[i])
		res[i] = idx
	}
	return res
}

// @lc code=end
