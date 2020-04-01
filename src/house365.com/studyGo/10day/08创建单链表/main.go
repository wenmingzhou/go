package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func createNote(head *ListNode, max int) {
	cur := head
	for i := 1; i < max; i++ {
		cur.val = i
		cur.next = &ListNode{}
		cur = cur.next
	}
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
	fmt.Println()
}

func main() {
	var head = new(ListNode)
	createNote(head, 6)
	PrintNode("前：", head)
}
