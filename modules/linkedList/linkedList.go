package linkedList

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// type LinkedList struct {
// 	head *ListNode
// }

//	func New() *LinkedList {
//	    return &LinkedList{Head: nil, Size: 0}
//	}
func printLLChild(l *ListNode, val string) string {
	if l.Next != nil {
		val += strconv.FormatInt(int64(l.Next.Val), 10) + ","
		return printLLChild(l.Next, val)
	}
	return val + "]"
}

func PrintLL(l *ListNode) string {
	llString := "["+strconv.FormatInt(int64(l.Val), 10)+","
	return printLLChild(l, llString)
}

func lengthChild(l *ListNode, i int) int {
	if l == nil {
		return i
	}
	return lengthChild(l.Next, i+1)
}

func Length(l *ListNode) int {
	return lengthChild(l.Next, 0)
}
