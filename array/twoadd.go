package main

import "fmt"

/*
	给定 nums = [2, 7, 11, 15], target = 9
	第一次循环：
		i=0,k=2,m[target-k]=m[7],此时map中不存在key=7的key和value，跳出if，将m[2]=0塞入字典
	第二次循环：
		i=1,k=7,m[target-k]=m[2],此时发现map在有key=2,value=0的记录，if判断为true，进入if内执行
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, k := range nums {
		if value, ok := m[target-k]; ok {
			return []int{value, i}
		}
		m[k] = i
	}
	return []int{0, 0}
}

func main() {
	numbers := []int{2, 3, 4, 5, 6, 7}
	tag := 12
	fmt.Println(twoSum(numbers, tag))
}
