/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (63.95%)
 * Likes:    664
 * Dislikes: 0
 * Total Accepted:    72.6K
 * Total Submissions: 110.8K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * 示例 1:
 *
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 *
 * 说明：
 *
 *
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
 *
 *
 */

// @lc code=start
func findDuplicate(nums []int) int {

		return binarySearch(nums)
		slow, fast := 0, 0
		for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {

		}
		slow = 0
		for slow != fast {
			slow = nums[slow]
			fast = nums[fast]
		}

		return slow


}


func binarySearch(nums []int) int {
	l, r, ans := 1, len(nums)-1, 0

	for l <= r{
		mid := (l+r) >> 1
		cnt := 0
		for i:=0; i<len(nums); i++{
			if nums[i] <= mid{
				cnt += 1
			}
		}

		if cnt <= mid{
			l = mid + 1
		}else{
			r = mid - 1
			ans = mid
		}
	}

	return ans
}
// @lc code=end
