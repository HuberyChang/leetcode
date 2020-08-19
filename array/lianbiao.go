// 单链表：
package main

import "fmt"

// 创建节点结构/类型
type Node struct {
	Data interface{}
	Next *Node
}

// 创建链表结构
type LList struct {
	Head   *Node
	Length int // 这里的链表长度不计入头节点
}

// a.设计接口：
type Method interface {
	Insert(i int, v interface{}) // 增
	Delete(i int)                // 删
	GetLength() int              // 获取长度
	Search(v interface{}) int    // 查
	isNull() bool                // 判断是否为空
}

// b.初始化函数：
// 创建节点
func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}

// 创建空链表
func CreateList() *LList {
	return &LList{CreateNode(nil), 0}
}

// c.基于链表结构体实现接口 Method 中的方法：
// 在 i 处插入节点（前插？？——即插入到原来的第 i 个节点之前，成为现在的第 i 个节点）
func (list *LList) Insert(i int, v interface{}) {
	s := CreateNode(v)
	pre := list.Head
	for count := 0; count < i; count++ { // 从head节点开始，head节点算第0个节点
		if count == i-1 {
			s.Next = pre.Next // 这两句不可颠倒
			pre.Next = s
			list.Length++
		}
		pre = pre.Next
	}
}

// 删除第 i 处节点
func (list *LList) Delete(i int) {
	pre := list.Head
	for count := 0; count < i; count++ {
		s := pre.Next
		if count == i-1 {
			pre.Next = s.Next
			list.Length--
		}
		pre = pre.Next
	}
}

// 返回链表长度
func (list *LList) GetLength() int {
	return list.Length
}

// 查询值 v 所在的位置
func (list *LList) Search(v interface{}) int {
	pre := list.Head.Next
	for i := 1; i <= list.Length; i++ {
		if pre.Data == v {
			return i
		}
		pre = pre.Next
	}
	return 0
}

// 判空
func (list *LList) isNull() bool {
	pre := list.Head.Next
	if pre == nil {
		return true
	}
	return false
}

// d.设计链表打印输出函数:
func PrintList(list *LList) {
	pre := list.Head.Next
	fmt.Println("LList shows as follows: ...")
	for i := 1; i <= list.Length; i++ {
		fmt.Printf("%v\n", pre.Data)
		pre = pre.Next
	}
}

// main 函数：
func main() {
	lList := CreateList()
	fmt.Println("List is null: ", lList.isNull())
	var M Method
	M = lList // 接口类型的变量可以存储所有实现该接口的类型变量
	M.Insert(1, 3)
	M.Insert(2, 6)
	M.Insert(1, 5)
	M.Insert(1, 7)

	PrintList(lList)
	fmt.Println("List length is: ", lList.Length)
	fmt.Println("元素6在位置：", M.Search(6))
	fmt.Println("元素100在位置：", M.Search(100))
	fmt.Println("List is null: ", lList.isNull())

	M.Delete(2)
	PrintList(lList)
	fmt.Println("List length is: ", lList.Length)
}
