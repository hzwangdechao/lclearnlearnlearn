/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 *
 * https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Easy (34.55%)
 * Likes:    345
 * Dislikes: 0
 * Total Accepted:    31.6K
 * Total Submissions: 90.8K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * 给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
 *
 * 你找到的子数组应是最短的，请输出它的长度。
 *
 * 示例 1:
 *
 *
 * 输入: [2, 6, 4, 8, 10, 9, 15]
 * 输出: 5
 * 解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
 *
 *
 * 说明 :
 *
 *
 * 输入的数组长度范围在 [1, 10,000]。
 * 输入的数组可能包含重复元素 ，所以升序的意思是<=。
 *
 *
 */

// @lc code=start

// 从左向右找最大， 从右向左找最小
func findUnsortedSubarray(nums []int) int {
	length := len(nums)

	if length < 2{
		return 0
	}

	high, low, minVal, maxVal :=  0, length-1, nums[length-1], nums[0]

	for i:=1; i<length; i++{
		maxVal = max(maxVal, nums[i])
		minVal = min(minVal, nums[length-i-1])

		// 如果当前值比最大值还小的话， 将high变化成当前位置
		if nums[i] < maxVal{
			high = i
		}

		// 如果当前值比最小值还大的话， 将low变化成当前位置
		if nums[length-i-1] > minVal{
			low = length-i-1
		}
	}
	fmt.Println(low, high)
	if low < high{
		return high - low + 1
	}

	return 0

}


func max(x, y int)int  {
	if x > y{
		return x
	}
	return y
}

func min(x, y int)int  {
	if x < y{
		return x
	}
	return y
}

// @lc code=end
