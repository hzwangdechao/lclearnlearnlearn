/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (52.63%)
 * Likes:    1677
 * Dislikes: 0
 * Total Accepted:    150.1K
 * Total Submissions: 285.1K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */

// @lc code=start
func trap(height []int) int {
	// return methodWithTwoList(height)
	// return methodWithTwoPoint(height)
	return methodWithStack(height)
}

// 利用栈来实现
func methodWithStack(height []int) int {
	if len(height) == 0{
		return 0
	}
	res := 0
	cur := 0
	stack := make([]int, 0)
	for cur < len(height){
		for len(stack) > 0 && height[cur] > height[stack[len(stack)-1]]{
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0{
				break
			}

			distance := cur - stack[len(stack)-1] - 1
			high := min(height[cur], height[stack[len(stack)-1]]) - height[last]

			res += distance * high
		}
		stack = append(stack, cur)
		cur ++
	}
	return res
}

func methodWithTwoPoint(height []int) int {
	if len(height) == 0{
		return 0
	}
	size := len(height)
	leftMax := height[0]
	rightMax := height[size-1]
	res := 0
	left := 0
	right := size - 1
	for left < right{
		if height[left] < height[right]{
			if height[left] > leftMax{
				leftMax = height[left]
			}else{
				res += leftMax - height[left]
			}
			left ++
		}else{
			if height[right] > rightMax{
				rightMax = height[right]
			}else{
				res += rightMax - height[right]
			}
			right --
		}
	}
	return res
}

// 使用两个数组来记录左侧元素最大值和右侧元素最大值
func methodWithTwoList(height []int) int {
	if len(height) == 0{
		return 0
	}
	left, right := make([]int, len(height)), make([]int, len(height))
	left[0] = height[0]
	right[len(height)-1] = height[len(height)-1]
	for i:=1; i<len(height); i++{
		left[i] = max(left[i-1], height[i])
	}
	for i:=len(height)-2; i>=0; i--{
		right[i] = max(right[i+1], height[i])
	}

	fmt.Println(left)
	fmt.Println(right)

	res := 0
	for i:=0; i<len(height); i++{
		res += min(left[i], right[i]) - height[i]
	}

	return res
}

func min(x , y int) int {
	if x < y{
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}

// @lc code=end
