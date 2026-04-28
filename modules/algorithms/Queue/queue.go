package queue

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}
func (q *Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is Empty")
		return -1
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

func (q *Queue) Peek() int {
	if q.isEmpty() {
		fmt.Println("Queue is Empty")
		return -1
	}
	return q.items[0]
}
