package main

import "fmt"

/*
	计数排序算法思想：
		针对的数据类型：
			量大但是范围小：比如某大型企业数万名员工年龄排序，如何快速得知高考名次（腾讯面试）


	变形后的计数数组实际上就是从下标1开始每一位都加上前面的值。加到最后就知道一共有多少人，同时还知道每个元素最终排的位置。然后再倒序输出就是排序
*/

func countingSort(arr []int) []int {

	//1.得到数列的最大值和最小值，并算出差值d
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	d := max - min

	//2.创建统计数组并统计对应元素个数
	bucketLen := d + 1
	countArray := make([]int, bucketLen)
	for i := 0; i < len(arr); i++ {
		countArray[arr[i]-min]++
	}

	//3.统计数组做变形，后面的元素等于前面的元素之和
	sum := 0
	for i := 0; i < len(countArray); i++ {
		sum += countArray[i]
		countArray[i] = sum
	}

	//4.倒序遍历原始数列，从统计数组找到正确位置，输出到结果数组
	//变形后的计数数组实际上就是从下标1开始每一位都加上前面的值。加到最后就知道一共有多少人，同时还知道每个元素最终排的位置。然后再倒序输出就是排序
	//countArray下标是5的元素，值是4，代表小绿的成绩排名位置在第4位。放在sortArray中，下标就要减1，因为下标从0开始 
	sortArray := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sortArray[countArray[arr[i]-min]-1] = arr[i]
		countArray[arr[i]-min]--
	}

	return sortArray
}

func main() {
	list := []int{2, 5, 1, 9, 4}

	fmt.Println(countingSort(list))
}
