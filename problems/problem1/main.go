package main

import (
	"dsa-go/modules/linkedList"
	"fmt"
)

func ll2Num(l *linkedList.ListNode) int {
	if l.Next == nil {
		return l.Val
	}
	sum := ll2Num(l.Next)*10 + l.Val
	return sum
}

func addTwoNumbersChild(l1 *linkedList.ListNode, l2 *linkedList.ListNode, carry int) *linkedList.ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 == nil && l2 == nil && carry == 1 {
        return &linkedList.ListNode{Val: carry, Next: nil}
    }
	val1, val2 := 0, 0
	if l1 != nil {
		val1 = l1.Val
	}
	if l2 != nil {
		val2 = l2.Val
	}

	val := val1 + val2 + carry
	l3 := &linkedList.ListNode{Val: val % 10}
	carryNext := val / 10

	var next1, next2 *linkedList.ListNode
	if l1 != nil {
		next1 = l1.Next
	}
	if l2 != nil {
		next2 = l2.Next
	}

	l3.Next = addTwoNumbersChild(next1, next2, carryNext)

	return l3
}

func addTwoNumbers(l1 *linkedList.ListNode, l2 *linkedList.ListNode) *linkedList.ListNode {
	return addTwoNumbersChild(l1, l2, 0)
}

func main() {
	l1a := linkedList.ListNode{Val: 2, Next: nil}
	l1b := linkedList.ListNode{Val: 4, Next: nil}
	l1c := linkedList.ListNode{Val: 3, Next: nil}
	l1d := linkedList.ListNode{Val: 9, Next: nil}
	l1a.Next = &l1b
	l1b.Next = &l1c
	l1c.Next = &l1d

	l2a := linkedList.ListNode{Val: 5, Next: nil}
	l2b := linkedList.ListNode{Val: 6, Next: nil}
	l2c := linkedList.ListNode{Val: 9, Next: nil}
	l2a.Next = &l2b
	l2b.Next = &l2c

	// l9a := linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9}}}}}}}
	// l9b := linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: &linkedList.ListNode{Val: 9, Next: nil}}}}
	// edge case [0] [0]
	// l0a := linkedList.ListNode{Val: 0, Next: nil}
	// l0b := linkedList.ListNode{Val: 0, Next: nil}

	// fmt.Println("Hello this is the code", linkedList.PrintLL(&l1a))
	fmt.Println("Hello this is the code", linkedList.PrintLL(addTwoNumbers(&l1a, &l2a)))
}
