package main

import "fmt"

func quicksSort(ss []int, low, high int) []int {
	if low >= high {
		return nil
	}
	base := ss[low]
	m, n := low, high
	for {
		if m < n {
			for {
				if m < n && ss[n] >= base {
					n--
				} else {
					ss[m] = ss[n]
					break
				}
			}
			for {
				if m < n && ss[m] < base {
					m++
				} else {
					ss[n] = ss[m]
					break
				}
			}
		} else {
			break
		}
		ss[m] = base
		quicksSort(ss, low, m-1)
		quicksSort(ss, m+1, high)
	}
	return ss
}

func main() {
	//输入数据
	//var n int
	//fmt.Scanf("%d\n", &n)
	//s := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scanf("%d", &s[i])
	//}
	//quicksSort(s, 0, n-1)
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%d", s[i])
	//}

	//指定数据
	st := []int{4, 1, 3, 5, 7, 8, 8, 89, 9}
	ret := quicksSort(st, 0, len(st)-1)
	fmt.Println(ret)
}
