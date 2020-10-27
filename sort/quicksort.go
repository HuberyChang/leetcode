package main

import "fmt"

func quick_sort(s []int, low, high int) {
	if low >= high {
		return
	}
	// 设置基准值
	x := s[low]
	// 设置哨兵
	start := low
	end := high
	for low < high {
		// 从右向左找，找到第一个比基准值小的数
		for low < high && s[high] >= x {
			high--
		}
		// 从左向右找，找到第一个比基准值大的数
		for low < high && s[low] <= x {
			low++
		}
		//if low > high {
		// 交换左右两侧的值
		s[low], s[high] = s[high], s[low]
		//}
	}
	// 将基准值移到哨兵相遇点
	s[high], s[start] = x, s[high]
	// 递归，左右两侧分别排序 
	quick_sort(s, start, low)
	quick_sort(s, high+1, end)
}

func main() {
	//var n int
	//fmt.Scanf("%d\n", &n)
	//s := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d", &s[i])
	//}
	//quick_sort(s, 0, n-1)
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%d", s[i])
	//}
	s := []int{4, 1, 7, 3, 5, 2}
	fmt.Println(s)
	quick_sort(s, 0, len(s)-1)
	fmt.Println(s)
}
