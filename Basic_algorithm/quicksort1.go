package main

import "fmt"

func quicks_sort(ss []int, low, high int) []int {
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
		quicks_sort(ss, low, m-1)
		quicks_sort(ss, m+1, high)
	}
	return ss
}

func main() {
	st := []int{4, 1, 3, 5, 7, 8, 8, 89, 9}
	ret := quicks_sort(st, 0, len(st)-1)
	fmt.Println(ret)
}
