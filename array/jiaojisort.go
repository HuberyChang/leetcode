package main

import (
	"fmt"
	"sort"
)

/*
	如果两个数组是有序的，则可以便捷地计算两个数组的交集。

	首先对两个数组进行排序，然后使用两个指针遍历两个数组。

	初始时，两个指针分别指向两个数组的头部。每次比较两个指针指向的两个数组中的数字，如果两个数字不相等，则将指向较小数字的指针右移一位，
	如果两个数字相等，将该数字添加到答案，并将两个指针都右移一位。当至少有一个指针超出数组范围时，遍历结束。

*/
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	var result []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result
}

func main() {
	list1 := []int{1, 2, 3, 4, 9}
	list2 := []int{5, 9, 7, 9, 10, 4}
	fmt.Println(intersect(list1, list2))
}
