/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 *
 * https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/description/
 *
 * algorithms
 * Easy (61.37%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    8.2K
 * Total Submissions: 13.3K
 * Testcase Example:  '[1,2,2,6,6,6,6,7,10]'
 *
 * 给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。
 *
 * 请你找到并返回这个整数
 *
 *
 *
 * 示例：
 *
 *
 * 输入：arr = [1,2,2,6,6,6,6,7,10]
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 10^4
 * 0 <= arr[i] <= 10^5
 *
 *
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	return methodWithBinarySearch(arr)
	// return recursive(arr)
}


func methodWithBinarySearch(arr []int)int  {
    n := len(arr)
    maybe :=  []int{arr[n/4], arr[n/2], arr[3*n/4]}
    for i := range maybe {
        if findLeft(arr, maybe[i]+1) - findLeft(arr, maybe[i]) > n/4 {
            return maybe[i]
        }
    }
    return 0
}

// 找到第一个大于等于目标元素的索引位置
func findLeft(arr []int, t int ) int {
	start, end := 0, len(arr)
	for start < end{
		mid := (start+end) >> 1
		if t <= arr[mid]{
			end = mid
		}else{
			start = mid + 1
		}
	}
	return start

}


// 通过遍历进行查找
func recursive(arr []int)int  {
	n := len(arr)

	cur := arr[0]
	cnt := 0
	for i:=0; i<n; i++{
		if arr[i] == cur{
			cnt += 1
			if cnt*4 > n{
				return cur
			}
		}else{
			cur = arr[i]
			cnt = 1
		}
	}
	return -1
}
// @lc code=end
