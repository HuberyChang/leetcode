package main

import "fmt"

/*
	状态：
		dp[i][j] : 表示包含第i行j列元素的最小路径和
	初始值：
		dp[0][0] = [0][0]位置所在的元素值
	状态转移方程：
		dp[i][j] = min(dp[i-1][j-1],dp[i-1][j]) + triangle[i][j]
	找到最后一行元素中，路径和最小的一个，就是我们的答案：
		l：dp数组长度
		result = min(dp[l-1,0]，dp[l-1,1]，dp[l-1,2]....)

	除了最顶上的元素之外，最左边的元素只能从自己头顶而来，最右边的元素只能从自己左上角而来。
	位于第2行的元素，都是特殊元素（因为都只能从[0,0]的元素走过来）
*/

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l < 1 {
		return 0
	}
	if l == 1 {
		return triangle[0][0]
	}
	ret := 1<<31 - 1
	triangle[0][0] = triangle[0][0]
	triangle[1][1] = triangle[1][1] + triangle[0][0]
	triangle[1][0] = triangle[1][0] + triangle[0][0]
	for i := 2; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == (len(triangle[i]) - 1) {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}
		}
	}
	for _, k := range triangle[l-1] {
		ret = min(ret, k)
	}
	return ret
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	triang := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triang))
}
