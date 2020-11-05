/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 *
 * https://leetcode-cn.com/problems/split-array-largest-sum/description/
 *
 * algorithms
 * Hard (44.02%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    13.9K
 * Total Submissions: 28.4K
 * Testcase Example:  '[7,2,5,10,8]\n2'
 *
 * 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
 *
 * 注意:
 * 数组长度 n 满足以下条件:
 *
 *
 * 1 ≤ n ≤ 1000
 * 1 ≤ m ≤ min(50, n)
 *
 *
 * 示例:
 *
 *
 * 输入:
 * nums = [7,2,5,10,8]
 * m = 2
 *
 * 输出:
 * 18
 *
 * 解释:
 * 一共有四种方法将nums分割为2个子数组。
 * 其中最好的方式是将其分为[7,2,5] 和 [10,8]，
 * 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
 *
 *
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	// left 当前数组中的最大值， right 数组中所有元素的和
	left, right :=0, 0
	for i:=0; i<len(nums); i++{
		right += nums[i]
		if left < nums[i]{
			left = nums[i]
		}
	}

	for left < right{
		mid := (right-left)/2 + left
		if check(nums, mid, m){
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}

// 验证已经分隔的次数是否不超过m
func check(nums []int, x, m int)bool  {
	sum, cnt := 0, 1
	for i:=0; i<len(nums); i++{
		if sum + nums[i] > x{
			cnt ++
			sum = nums[i]
		}else{
			sum += nums[i]
		}
	}
	fmt.Println(x, cnt, m)
	return cnt <= m
}

// @lc code=end
