/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 *
 * https://leetcode-cn.com/problems/freedom-trail/description/
 *
 * algorithms
 * Hard (41.17%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 7.8K
 * Testcase Example:  '"godding"\n"gd"'
 *
 * 视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。
 *
 * 给定一个字符串 ring，表示刻在外环上的编码；给定另一个字符串 key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
 *
 * 最初，ring 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00
 * 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
 *
 * 旋转 ring 拼出 key 字符 key[i] 的阶段中：
 *
 *
 * 您可以将 ring 顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00
 * 方向对齐，并且这个字符必须等于字符 key[i] 。
 * 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key
 * 的下一个字符（下一阶段）, 直至完成所有拼写。
 *
 *
 * 示例：
 *
 *
 *
 *
 *
 *
 * 输入: ring = "godding", key = "gd"
 * 输出: 4
 * 解释:
 * ⁠对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 * ⁠对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 * ⁠当然, 我们还需要1步进行拼写。
 * ⁠因此最终的输出是 4。
 *
 *
 * 提示：
 *
 *
 * ring 和 key 的字符串长度取值范围均为 1 至 100；
 * 两个字符串中都只有小写字符，并且均可能存在重复字符；
 * 字符串 key 一定可以由字符串 ring 旋转拼出。
 *
 */

// @lc code=start
func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	pos := [26][]int{} // 记录ring中每个元素出现的位置
	for i := 0; i < n; i++ {
		pos[ring[i]-'a'] = append(pos[ring[i]-'a'], i)
	}
	inf := math.MaxInt32 / 2
	dp := make([][]int, m) // dp[i][j]记录key[i]的字符转动到 ring[j]的位置最少需要转动的次数
	for i := 0; i < m; i++ {
		l := make([]int, n)
		for j := 0; j < n; j++ {
			l[j] = inf
		}
		dp[i] = l
	}

	// key的第一个字符可以再原始基础上进行转动, 因为不需要查看上一个字符的转动位置
	for _, j := range pos[key[0]-'a'] {
		dp[0][j] = min(j, n-j) + 1
	}

	// key从第二个位置开始需要根据上一个转动位置来计算
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] { // 本次需要转动到的位置
			for _, k := range pos[key[i-1]-'a'] { // 上一次转动停留的位置
				// 上一次转动的次数 + 上次停留的位置到本次需要转动到达位置的距离差 + 1
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1) // 转动到dp[i][j]所需要的做少的转动次数
			}
		}
	}

	// 从dp[m-1]... 中选出最小值作为结果返回
	return min(dp[m-1]...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
