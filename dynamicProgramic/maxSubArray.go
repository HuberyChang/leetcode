package main

import "fmt"

/*
	dp[i]：表示以 nums[i] 结尾的连续子数组的最大和。
	如果要得到 dp[i]，那么 nums[i] 一定会被选取。并且 dp[i] 所表示的连续子序列与 dp[i-1] 所表示的连续子序列很可能就差一个 nums[i]

	dp[i] = dp[i-1]+nums[i] , if (dp[i-1] >= 0)
	dp[i] = nums[i] , if (dp[i-1] < 0)
	综上：
		dp[i]=max(nums[i], dp[i−1]+nums[i])
*/

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	slic := []int{-1, 2, -3, -4, 5, -9, 2, 3}
	fmt.Println(maxSubArray(slic))
}
