package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	一趟扫描的实现方法
	思路：定义2个索引p和q，p，q距离相差n+1个节点距离，因为链表最后一个节点是nil，
	当q==nil时，p的下一个节点p.Next为要删除的节点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{}
	sentry.Next = head

	var start, end *ListNode
	start = sentry
	end = sentry
	for i := 0; i < n+1; i++ {
		end = end.Next
	}
	if start != nil {
		start = start.Next
		end = end.Next
	}
	end.Next = end.Next.Next
	return sentry.Next
}

func main() {

}
