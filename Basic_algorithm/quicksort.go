package main

import "fmt"

func quick_sort(q []int, start, end int) {
	if start >= end {
		return
	}
	x := q[start]
	i, j := start, end
	for i < j {
		for q[j] >= x && i < j {
			j--
		}
		for q[i] < x && i < j {
			i++
		}
		q[i], q[j] = q[j], q[i]
	}
	quick_sort(q, start, i-1)
	quick_sort(q, i+1, end)
}

func main() {
	//var n int
	//fmt.Scanf("%d\n", &n)
	//s := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d\n", &s[i])
	//}
	//quick_sort(s, 0, n-1)
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%d", s[i])
	//}
	s := []int{3, 1, 2, 3, 5}
	quick_sort(s, 0, len(s)-1)
	fmt.Println(s)
}
