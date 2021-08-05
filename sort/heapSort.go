package main

import "fmt"

/*
	堆的两个条件：
		1、必须是完全二叉树
		2、必须父节点的值大于左右子节点的值

	算法步骤：
		1、将待排序序列 H[0……n-1]构建成一个堆，根据（升序降序需求）选择大顶堆或小顶堆。构建堆的过程可能需要递归操作，以致满足堆的两个性质；
		2、取出每次构建完堆的堆顶元素。然后把堆尾元素放到堆首；
		3、重复步骤 2，直到堆的尺寸为 1
*/

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//进行排序操作
func heapSort(arr []int) []int {
	arrlen := len(arr)
	buildMaxHeap(arr, arrlen) // 执行完毕，会构建出一个大根堆
	for i := arrlen - 1; i >= 0; i-- {
		//fmt.Println(i)
		swap(arr, 0, i)         // 每次交换堆顶和堆尾元素
		arrlen -= 1             // 堆规模缩小1
		heapify(arr, 0, arrlen) // 此时从上往下进行heapify
	}
	return arr
}

//对整个数组构建大根堆，从最后一个非叶子节点开始构建
func buildMaxHeap(arr []int, arrlen int) {
	for i := arrlen / 2; i >= 0; i-- { // i从数组长度的一半开始递减
		fmt.Println(i, "=====", arrlen)
		heapify(arr, i, arrlen)
	}
}

//从下标为数组节点的一半开始进行heapify
func heapify(arr []int, i, arrlen int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	//fmt.Println(i, "++++++++")
	if left < arrlen && arr[left] > arr[max] { // 最下边一层的节点的left或者right不一定存在，不能超过数组长度
		max = left
	}
	if right < arrlen && arr[right] > arr[max] {
		max = right
	}
	if max != i { // 如果下标为i的数本身不是所在子树的最大元素，就交换 
		swap(arr, i, max)
		heapify(arr, max, arrlen)
	}
}

func main() {
	list := []int{4, 2, 5, 19, 7, 3, 4}
	fmt.Println(heapSort(list))
}
