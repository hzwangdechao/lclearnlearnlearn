/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (37.34%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    18.5K
 * Total Submissions: 49.5K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 说明:
 * 不允许旋转信封。
 *
 * 示例:
 *
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 *
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	// 先对数组的第一个数字进行排序，如果第一个数字相同的话， 则降序。
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var nums []int
	for _, envelop := range envelopes {
		nums = append(nums, envelop[1])
	}

	// 对nums求最长上升子序列的长度
	var dp []int
	for _, num := range nums {

		if len(dp) == 0 || dp[len(dp)-1] < num { // 如果dp长度为0或者dp的最后一个元素比当前元素小的话， 直接将当前元素添加到dp中
			dp = append(dp, num)
		} else {
			// 利用二分查找的方法， 找到dp中第一个num应该插入的位置， 并将指定位置的值替换成num
			left, right := 0, len(dp)-1
			loc := right

			for left <= right {
				mid := (right-left)/2 + left

				if dp[mid] == num {
					loc = mid
					right = mid - 1
				} else if dp[mid] > num {
					loc = mid
					right = mid - 1
				}else if dp[mid] < num{
					left = mid + 1
				}
			}
			dp[loc] = num
		}
	}

	return len(dp)

}
// @lc code=end
