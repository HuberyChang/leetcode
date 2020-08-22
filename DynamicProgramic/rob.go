package main

import "fmt"

/*	[1 2 3 1]
	到第i间房，可以分两种情况：偷、不偷
	第一间房：S0=H0=1
	前两间房：S1=max(S0,H1)=2
	前三间房：S2=max(S1,S0+H2)=4
	前四间房：S3=max(S1,S2+H3)=4
	递推公式：
		Sn=max(Sn-1,Sn-2+Hn)
	偷前n-1间房屋的最高金额或者
	偷前n-2间房屋的最高金额加第n间房屋的金额
*/

func rob(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return maxrob(nums[0], nums[1])
	}
	nums[1] = maxrob(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = maxrob(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func maxrob(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	num := []int{1, 2, 4, 3, 6, 9, 4, 7}
	fmt.Println(rob(num))
}
