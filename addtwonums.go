/* 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807 */

package main

import (
	"fmt"
)

func test() {
	// l1 := &ListNode{9, &ListNode{1, &ListNode{6, nil}}}
	// l2 := &ListNode{rand.Intn(10), &ListNode{rand.Intn(10), &ListNode{rand.Intn(10), nil}}}
	l2 := &ListNode{5, nil}
	l1 := &ListNode{5, nil}
	fmt.Println("add ", l1, l2, "result", addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// fmt.Println("addTwoNumbers", l1, l2)
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{0, nil}
	t := 0
	var next1, next2 *ListNode
	if l1 != nil {
		t += l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		t += l2.Val
		next2 = l2.Next
	}
	if t >= 10 {
		switch {
		case next1 != nil:
			next1.Val += t / 10
		case next2 != nil:
			next2.Val += t / 10
		default:
			next1 = &ListNode{Val: t / 10}
		}

	}
	result.Val = t % 10
	if next1 != nil || next2 != nil {
		result.Next = addTwoNumbers(next1, next2)
	}

	return result
}

//ListNode 1
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return fmt.Sprintf("{%v, ->%v}", node.Val, node.Next)
}
