package main

import "fmt"

/*
	dp[i] ：表示以nums[i]结尾的最长上升子序列的长度

	如果nums[i]比前面的所有元素都小，那么dp[i]等于1（即它本身）（该结论正确）
	如果nums[i]前面存在比他小的元素nums[j]，那么dp[i]就等于dp[j]+1（该结论错误，比如nums[3]>nums[0]，即9>1,但是dp[3]并不等于dp[0]+1）

	因为dp[i]前面比他小的元素，不一定只有一个！

	可能除了 nums[j]，还包括 nums[k]，nums[p] 等等等等。所以 dp[i] 除了可能等于 dp[j]+1，还有可能等于 dp[k]+1，dp[p]+1 等等等等。
	所以我们求 dp[i]，需要找到 dp[j]+1，dp[k]+1，dp[p]+1 等等等等 中的最大值。

			dp[i] = max(dp[j]+1，dp[k]+1，dp[p]+1，.....)
			只要满足：
				nums[i] > nums[j]
				nums[i] > nums[k]
				nums[i] > nums[p]
				....
*/

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxlen(dp[j]+1, dp[i])
				fmt.Printf("dp[%d]:%d\n", i, dp[i])
			}
		}
		result = maxlen(result, dp[i])
	}
	return result
}

func maxlen(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	lenslic := []int{1, 9, 5, 9, 3}
	fmt.Println(lengthOfLIS(lenslic))
}
