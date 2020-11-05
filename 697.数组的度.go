/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 *
 * https://leetcode-cn.com/problems/degree-of-an-array/description/
 *
 * algorithms
 * Easy (52.76%)
 * Likes:    148
 * Dislikes: 0
 * Total Accepted:    18.4K
 * Total Submissions: 34.2K
 * Testcase Example:  '[1,2,2,3,1]'
 *
 * 给定一个非空且只包含非负数的整数数组 nums, 数组的度的定义是指数组里任一元素出现频数的最大值。
 *
 * 你的任务是找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
 *
 * 示例 1:
 *
 *
 * 输入: [1, 2, 2, 3, 1]
 * 输出: 2
 * 解释:
 * 输入数组的度是2，因为元素1和2的出现频数最大，均为2.
 * 连续子数组里面拥有相同度的有如下所示:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * 最短连续子数组[2, 2]的长度为2，所以返回2.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,2,3,1,4,2]
 * 输出: 6
 *
 *
 * 注意:
 *
 *
 * nums.length 在1到50,000区间范围内。
 * nums[i] 是一个在0到49,999范围内的整数。
 *
 *
 */

// @lc code=start

// 找到数组中出现次数最大的值（可以是一个， 也可以是多个）
// 找到这些值的开始位置和结束位置距离的最小值

func findShortestSubArray(nums []int) int {
	// 统计每个元素出现的次数
	maxMap := make(map[int]int, 0)
	for _, num := range nums{
		if _, ok := maxMap[num]; !ok{
			maxMap[num] = 1
		}else{
			maxMap[num] += 1
		}
	}
	// 找出最大的出现次数
	maxCount := 0
	for _, v := range maxMap{
		if v > maxCount{
			maxCount = v
		}
	}
	// 找出出现次数最大的值
	maxKey := make([]int, 0)
	for k, v := range maxMap{
		if v == maxCount {
			maxKey = append(maxKey, k)
		}
	}
	if maxCount == 1 {
		// 所有元素只出现过一次
		return 1
	}

	// 假设最小的度为 math.MaxInt64
	minDegree := math.MaxInt64

	for _, key := range maxKey{
		start := 0  // 最先出现的位置
		end := 0  // 最后一次出现的位置

		for i:=0; i<len(nums);i++{
			if nums[i] == key{
				start = i
				// 找到开始位置后开始寻找结束位置
				for j:=len(nums)-1;j>i;j--{
					if nums[j] == key{
						end = j
						break
						// 找到结束位置后退出循环
					}
				}
				break
			}
		}
		curDegree := end - start + 1  // 当前元素的度

		if curDegree < minDegree{
			minDegree = curDegree
		}
	}

	return minDegree
}
// @lc code=end
