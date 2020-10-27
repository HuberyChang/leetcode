package main

import "fmt"

func onezero(nums [][]int, total int) int {
	dp := make([][]int, len(nums))
	//fmt.Printf("%T\n", nums)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}
	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}
	for i := 1; i < len(nums); i++ {
		for j := nums[i][0]; j <= total; j++ {
			dp[i][j] = maxonezero(dp[i-1][j], dp[i-1][j-nums[i][0]]+nums[i][1])
		}
	}
	return dp[len(nums)-1][total]
}

func maxonezero(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	ret := onezero([][]int{[]int{5, 12}, []int{4, 3}, []int{7, 10}, []int{2, 3}, []int{6, 6}}, 15)
	fmt.Println(ret)
}
