package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	node := new(ListNode)
	node.val = 0

	node1 := new(ListNode)
	node1.val = 1
	node.next = node1

	node2 := new(ListNode)
	node2.val = 2
	node1.next = node2

	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode)

			nowNode = nowNode.next
			continue
		}
		break
	}
}
