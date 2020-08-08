package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/*
		一趟扫描的实现方法
		思路：定义2个索引p和q，p，q距离相差n+1个节点距离，因为链表最后一个节点是nil，
		当q==nil时，p的下一个节点p.Next为要删除的节点
	*/

	// 虚拟头结点
	//dummyHead := &ListNode{Next: head}
	//
	//p, q := dummyHead, dummyHead
	//
	//for i := 0; i < n+1; i++ {
	//	q = q.Next
	//}
	//
	//for q != nil {
	//	p = p.Next
	//	q = q.Next
	//}
	//p.Next = p.Next.Next
	//
	//return dummyHead.Next

	/*
			本题已经有head节点，新建一个别的哨兵节点
		思路：
			首先我们定义好哨兵节点result，指向哨兵节点的目标元素指针cur，以及目标指针cur的前一个指针pre，此时pre指向nil。
			接下来我们开始遍历整个链表。
			当head移动到距离目标元素cur的距离为N-1时，同时开始移动cur。
			当链表遍历完之后，此时head指向nil，这时的cur就是我们要找的待删除的目标元素。我们只要同时定位到要删除的元素的前1个元素。
			最后我们通过pre.Next = pre.Next.Next完成删除操作，就完成了整个解题过程。
	*/
	result := &ListNode{Next: head}
	var pre *ListNode // 只定义，未分配变量，pre的值默认为nil
	cur := result
	i := 1 // 从第一个节点开始
	for head != nil {
		if i >= n { // 从i=n开始，指针开始移动，因为i从1开始，所以移动后，head和cur之间隔着n-1个节点
			pre = cur
			cur = cur.Next
		}
		head = head.Next // head不为nil就一直往后移动指针
		i++
	}
	pre.Next = pre.Next.Next

	return result.Next

}

func main() {

}
