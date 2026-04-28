package queue

import "fmt"

type CircularQueue struct {
	items []int
	size  int
	front int
	rear  int
	count int
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, size),
		size:  size,
		front: 0,
		rear:  -1,
		count: 0,
	}
}
func (q *CircularQueue) isEmpty() bool {
	return q.count == 0
}

func (q *CircularQueue) isFull() bool {
	return q.count == q.size
}

func (q *CircularQueue) Enqueue(value int) {
	if q.isFull() {
		fmt.Println("Queue is Full")
		return
	}
	q.rear = (q.rear + 1) % q.size
	q.items[q.rear] = value
	q.count++
}

func (q *CircularQueue) Dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue is Empty")
		return -1
	}
	val := q.items[q.rear]
	q.front = (q.front + 1) % q.size
	q.count--
	return val
}

func (q *CircularQueue) Peek() int {
	if q.isEmpty() {
		fmt.Println("Queue is Empty")
	}
	return q.items[q.front]
}
