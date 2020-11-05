/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 *
 * https://leetcode-cn.com/problems/video-stitching/description/
 *
 * algorithms
 * Medium (48.03%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    25.9K
 * Total Submissions: 45.3K
 * Testcase Example:  '[[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]\n10'
 *
 * 你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
 *
 * 视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1]
 * 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。
 *
 * 我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
 * 输出：3
 * 解释：
 * 我们选中 [0,2], [8,10], [1,9] 这三个片段。
 * 然后，按下面的方案重制比赛片段：
 * 将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
 * 现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。
 *
 *
 * 示例 2：
 *
 *
 * 输入：clips = [[0,1],[1,2]], T = 5
 * 输出：-1
 * 解释：
 * 我们无法只用 [0,1] 和 [1,2] 覆盖 [0,5] 的整个过程。
 *
 *
 * 示例 3：
 *
 *
 * 输入：clips =
 * [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]],
 * T = 9
 * 输出：3
 * 解释：
 * 我们选取片段 [0,4], [4,7] 和 [6,9] 。
 *
 *
 * 示例 4：
 *
 *
 * 输入：clips = [[0,4],[2,8]], T = 5
 * 输出：2
 * 解释：
 * 注意，你可能录制超过比赛结束时间的视频。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * 0
 *
 *
 */

// @lc code=start

func videoStitching(clips [][]int, T int) int {
	return methodWithEnu(clips, T)
	return methodWithDP(clips, T)
}

// 枚举的方法
func methodWithEnu(clips [][]int, T int) int {
	list := make([]int, T)
	// list[i] 代表从起点l开始，能到达的最远位置
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < T {
			list[l] = max(list[l], r)
		}
	}

	pre, last, res := 0, 0, 0
	for i, v := range list {
		if v > last {
			last = v
		}

		if i == last { // 无法完成覆盖
			return -1
		}

		if i == pre { // 需要进行更新的操作
			res++
			pre = last
		}
	}

	return res
}

// 动态规划的方法
func methodWithDP(clips [][]int, T int) int {
	inf := math.MaxInt32 - 1
	// 构建动态规划数组, dp[i]代表覆盖0~i所需要的最少的数组的数量
	dp := make([]int, T+1)
	dp[0] = 0
	for i := 1; i <= T; i++ {
		dp[i] = inf
	}

	for i := 1; i <= T; i++ { // 寻找从1~T各个元素能够覆盖所需要的最少数组数量
		for _, interval := range clips {
			l, r := interval[0], interval[1]

			// 如果 i 在 l~r的范围内的话
			if i > l && i <= r {
				dp[i] = min(dp[i], dp[l]+1)
			}
		}
	}

	if dp[T] == inf { // 说明不能覆盖到T
		return -1
	}
	return dp[T]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
