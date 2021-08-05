package main

import "fmt"

func main() {
	var n int
	fmt.Print("请输入一个正整数：")
	fmt.Scanf("%d", &n)
	sum := 1
	for i := 2; i*i < n; i++ { // 一个整数的因子不可能超过它的一半
		//if i*i != n { // i*i<n就不可能i*i=n，可以去掉判断
		if n%i == 0 {
			sum += i + n/i
		} else {
			sum += i
		}
		//}
	}
	if sum == n {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	} 
}
