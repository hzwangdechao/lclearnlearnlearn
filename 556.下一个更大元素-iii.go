/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 *
 * https://leetcode-cn.com/problems/next-greater-element-iii/description/
 *
 * algorithms
 * Medium (31.62%)
 * Likes:    99
 * Dislikes: 0
 * Total Accepted:    6.9K
 * Total Submissions: 21.9K
 * Testcase Example:  '12'
 *
 * 给定一个32位正整数 n，你需要找到最小的32位整数，其与 n 中存在的位数完全相同，并且其值大于n。如果不存在这样的32位整数，则返回-1。
 *
 * 示例 1:
 *
 *
 * 输入: 12
 * 输出: 21
 *
 *
 * 示例 2:
 *
 *
 * 输入: 21
 * 输出: -1
 *
 *
 */

/*														   高位   低位
从低位到高位，如果每位数字是依次递增的，那么就不存在符合题意的数，比如44332211。
需要从低位到高位遍历，判断是否能找到一个比高位某数更大的最小值。以12443111为例，可以发现低位的3是比高位的2大的最小值(4也比2大，但不是最小值)，那么可以将3和2进行交换，得到13442111，然后再把4以后的数按从小到大的顺序进行排列即可，答案为13111244.
怎样找到比低位大的最小值呢？可以考虑使用栈，从后往前，维持一个递增栈，当遇到小于栈顶的元素时(索引为i)，依次出栈，并将栈顶元素的索引保存下来，从而找到低位比当前元素大的最小值所在的位置idx，将两个元素交换。然后从i+ 1开始，进行从小到大排序。即可得到答案。
*/
/*

19443151
19543131

19543311


19443511




*/

// @lc code=start
func nextGreaterElement(n int) int {
	var (
		res      int
		stack    []int             // 单调递增的栈， 存放元素的索引
		str      = strconv.Itoa(n) // 将数字转换成字符串
		numSlice = []byte(str)     // 将字符串转换成 字节切片的形式， 方便后面的位置交换
		idx      int               // 需要交换位置的索引
	)

	// 从后向前遍历
	i := len(str) - 1
	for ; i >= 0; i-- {

		// 构造单调递增栈， 如果当前元素比栈顶元素小的话， 则将栈顶元素出栈， 最后一个出栈的元素就是第一个比当前元素大的元素
		for len(stack) > 0 && numSlice[i] < numSlice[stack[len(stack)-1]] {
			idx = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		// 如果找到了idx， 则交换位置并且推出遍历
		if idx != 0 {
			numSlice[i], numSlice[idx] = numSlice[idx], numSlice[i]
			break
		}
		// 最后将元素的索引添加到栈中
		stack = append(stack, i)
	}

	// 从后向前遍历完毕之后， 如果idx=0的话， 说明n是332211   21   4433  这样的情况， 直接返回-1
	if idx == 0 {
		return -1
	}

	// 对numSlice中i+1之后的元素进行升序排序
	tmp := numSlice[i+1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	// 将字节列表转换成数字
	nums := append(numSlice[:i+1], tmp...)
	for _, num := range nums {
		res = res*10 + int(num-'0')
	}
	// 越界判断
	if res > math.MaxInt32{
		return -1
	}
	return res
}

// @lc code=end
