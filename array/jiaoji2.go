package main

import "fmt"

/*
	示例 :
		输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
		输出：[4,9]

	说明：
		1、输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
		2、我们可以不考虑输出结果的顺序。

	根据题目说明，需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>
*/

func jiaoji(nums1, nums2 []int) []int {
	// 对小数组进行哈希映射，会减少内存的使用
	if len(nums1) > len(nums2) {
		return jiaoji(nums2, nums1)
	}

	res := []int{}
	num := make(map[int]int, 0)
	//对num1进行遍历，将元素的值和出现的次数存入map 
	for _, v1 := range nums1 {
		num[v1]++
	}
	//对num2进行遍历，判断num2的元素在map中对应值的次数是否为正
	for _, v2 := range nums2 {
		if num[v2] > 0 {
			res = append(res, v2)
			num[v2]--
		}
	}
	return res
}

func main() {
	lis1 := []int{1, 2, 2, 4}
	lis2 := []int{2, 2}
	fmt.Println(jiaoji(lis1, lis2))
}
