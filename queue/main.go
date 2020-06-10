package main

import "fmt"

// 循环队列
func main() {
	s := Constructor(3)
	fmt.Println(s.EnQueue(1), s.data)
	fmt.Println(s.EnQueue(2), s.data)
	fmt.Println(s.EnQueue(3), s.data)
	fmt.Println(s.EnQueue(4), s.data)
	fmt.Println(s.DeQueue(), s.data)
	fmt.Println(s.EnQueue(4), s.data)
}

type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	mq := MyCircularQueue{}
	mq.data = make([]int, k)
	mq.head = -1
	mq.tail = -1
	mq.size = k
	return mq
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.size
	this.data[this.tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % this.size
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.size == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
