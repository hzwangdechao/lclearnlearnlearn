/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (48.71%)
 * Likes:    527
 * Dislikes: 0
 * Total Accepted:    71.8K
 * Total Submissions: 146.7K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 *
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 * Sd
 *
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n * k == 0{
		// 空数组或k为0的情况
		return []int{}
	}
	if k == 1{
		return nums
	}
	left := make([]int, n)
	left[0] = nums[0]
	right := make([]int, n)
	right[n-1] = nums[n-1]

	for i:=1; i<n; i++{
		// 从左向右填充left
		if i % k == 0{
			left[i] = nums[i]
		}else{
			left[i] = max(left[i-1], nums[i])
		}

		// 从右向左填充right
		j := n - i - 1

		fmt.Println(j)
		if (j+1) % k == 0{
			right[j] = nums[j]
		}else{
			right[j] = max(right[j+1], nums[j])
		}
	}

	var output []int

	for i:=0; i<n-k+1; i++{
		output = append(output, max(left[i+k-1], right[i]))
	}
	return output
}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y

}
// @lc code=end
