/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 *
 * https://leetcode-cn.com/problems/kth-missing-positive-number/description/
 *
 * algorithms
 * Easy (53.86%)
 * Likes:    7
 * Dislikes: 0
 * Total Accepted:    5K
 * Total Submissions: 9.2K
 * Testcase Example:  '[2,3,4,7,11]\n5'
 *
 * 给你一个 严格升序排列 的正整数数组 arr 和一个整数 k 。
 *
 * 请你找到这个数组里第 k 个缺失的正整数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [2,3,4,7,11], k = 5
 * 输出：9
 * 解释：缺失的正整数包括 [1,5,6,8,9,10,12,13,...] 。第 5 个缺失的正整数为 9 。
 *
 *
 * 示例 2：
 *
 * 输入：arr = [1,2,3,4], k = 2
 * 输出：6
 * 解释：缺失的正整数包括 [5,6,7,...] 。第 2 个缺失的正整数为 6 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 1000
 * 1 <= arr[i] <= 1000
 * 1 <= k <= 1000
 * 对于所有 1 <= i < j <= arr.length 的 i 和 j 满足 arr[i] < arr[j] 
 *
 *
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	return methodWithEnumerate(arr, k)
}

// 枚举法
func methodWithEnumerate(arr []int, k int) int {
	missCount := 0
	lastMiss := -1
	current := 1
	ptr := 0

	for missCount < k{
		if current == arr[ptr]{
			if ptr+1 < len(arr){
				ptr += 1
			}
		}else{
			missCount += 1
			lastMiss = current
		}
		current += 1
	}

	return lastMiss
}
// @lc code=end
class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        if arr[0] > k:
            return k

        l, r = 0, len(arr)
        while l < r:
            mid = (l + r) >> 1
            x = arr[mid] if mid < len(arr) else 10**9
            if x - mid - 1 >= k:
                r = mid
            else:
                l = mid + 1

        return k - (arr[l - 1] - (l - 1) - 1) + arr[l - 1]
