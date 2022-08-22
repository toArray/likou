package luoqiangLikou

import (
	"fmt"
	"testing"
)

//19. 删除链表的倒数第 N 个结点
func Test19(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	//head.Next.Next = &ListNode{Val: 3, Next: nil}
	//head.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	data := removeNthFromEnd(head, 2)
	fmt.Println(data)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	diff := 0
	slow := head
	fast := head
	var pre *ListNode

	for fast != nil {
		if diff < n {
			fast = fast.Next
			diff++
		} else {
			pre = slow
			slow = slow.Next
			fast = fast.Next
		}
	}

	//删除倒数第一个
	if slow == head {
		return head.Next
	}

	if pre != nil {
		pre.Next = slow.Next
		return head
	}

	return nil
}
